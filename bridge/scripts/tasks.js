task("list-contracts", "Exports & prints the list of deployed contracts")
    .setAction(async (taskArgs, hre) => {
        const filename = `${hre.network.name}-exports.json`;
        await hre.run('export', {export: filename});
        const obj = require(`../${filename}`).contracts;
        let result = {};
        Object.keys(obj).map(k => {
            result[k] = obj[k].address;
        });
        console.log(result);
        const { getImplementation, getInstance, getCodeSlot } = require('./utils');
        let addr = await getImplementation(result.TransparentUpgradeableProxy);
        console.log(`Vault Proxy's current implementation is ${ethers.utils.hexlify(addr)}`);
        const kyb = await getInstance('KBNTrade');
        addr = await kyb.kyberNetworkProxyContract();
        console.log(`KBNTrade is using KyberNetworkRouter ${ethers.utils.hexlify(addr)}`);
        const uni = await getInstance('UniswapV2Trade');
        addr = await uni.uniswapV2();
        console.log(`UniswapV2Trade is using UniswapRouter ${ethers.utils.hexlify(addr)}`);

        // get data from unstructured storage
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0x62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd2')
        console.log(`Using IncognitoProxy instance at ${ethers.utils.hexlify(addr)}`);
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103')
        console.log(`Current Vault Admin : ${ethers.utils.hexlify(addr)}`);
        addr = await getCodeSlot(result.TransparentUpgradeableProxy, '0x7b13fc932b1063ca775d428558b73e20eab6804d4d9b5a148d7cbae4488973f8')
        console.log(`Next Vault Admin : ${ethers.utils.hexlify(addr)}`);
    });

task("unshield-status", "Asks if the contract has processed an unshield")
    .addOptionalParam("id", "The unshield id", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const incTxHash = taskArgs.id;
        const vault = await ethers.getContractAt('Vault', await hre.deployments.get('Vault'));
        const res = await vault.isWithdrawed(incTxHash);
        console.log(res);
    });

task("vault-balance-status", "Balance by token")
    .addOptionalParam("token", "The ERC20 token", "0x0000000000000000000000000000000000000000", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        let token = taskArgs.token;
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const b = await vault.balanceOf(token);
        const res = await vault.totalDepositedToSCAmount(token);
        console.log(`Contract actual balance ${b.toString()}`);
        console.log(`Contract deposited amount ${res.toString()}`);
    });

task("show-committees", "Asks the IncProxy contract for its committee at height")
    .addOptionalParam("height", "The height", 0, types.int)
    .addOptionalParam("address", "The address for IncProxy", "0x0000000000000000000000000000000000000000", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const ip = await getInstance('IncognitoProxy', null, taskArgs.address);
        const res = await ip.findBeaconCommitteeFromHeight(taskArgs.height);
        console.log(res);
    });

subtask("contract-call-setup", "Unshield (withdraw) from Incognito")
    .setAction(async (taskArgs, hre) => {
        const { inc } = require('./external');
        const { getInstance, confirm, toIncDecimals } = require('./utils');
        const { proveEth, formatBurnProof } = require('./prove');
        const { tokenAddresses } = require('./constants');
        const BN = ethers.BigNumber;
        const signers = await ethers.getSigners();
        const defaultSigner = signers[3];
        await inc.init(null, hre.networkCfg().providers[0], taskArgs.shards ? taskArgs.shards : 8);
        const transactor = taskArgs.privkey ? await inc.NewTransactor(taskArgs.privkey) : null;
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const proxy = await getInstance('TransparentUpgradeableProxy');

        return {
            getInstance, confirm, proveEth, formatBurnProof,
            Inc, inc, transactor, tokenAddresses,
            BN, toIncDecimals,
            defaultSigner,
            vault, proxy,
        }
    });

task("shield", "Shield (deposit) to Incognito")
    .addOptionalParam("token", "Token to shield", "0x0000000000000000000000000000000000000000", types.string)
    .addOptionalParam("amount", "Amount to shield (in wei)", "0", types.string)
    .addOptionalParam("to", "Receiver Incognito address", "", types.string)
    .addOptionalParam("privkey", "Your private key (if you're submitting proof from this CLI)", "", types.string)
    .setAction(async taskArgs => {
        let { token, amount, to, privkey } = taskArgs;
        const { Inc, BN, vault, defaultSigner, confirm, transactor, proveEth, tokenAddresses, toIncDecimals, getInstance } = await run('contract-call-setup', { privkey });
        try {
            amount = BN.from(amount);
            const temp = Inc.utils.base58CheckDeserialize(to);
        } catch (e) {
            console.error(e);
            throw 'Invalid Parameter';
        }
        let tx;
        try {
            if (token == tokenAddresses.ETH) {
                tx = await confirm(vault.connect(defaultSigner).deposit(to, {value: amount}), 16);
            } else {
                const tokenInst = await getInstance('contracts/IERC20.sol:IERC20', null, token);
                await confirm(tokenInst.connect(defaultSigner).approve(vault.address, amount));
                tx = await confirm(vault.connect(defaultSigner).depositERC20(token, amount, to), 16);
            }
        } catch(e) {
            console.error('Shielding failed at deposit step');
            throw e;
        }
        if (privkey) {
            if (!tokenAddresses.pTokens[token]) throw `TokenID unavailable for ${token}`;
            console.log('Now submitting deposit proof to issue pETH');
            const proveResult = await proveEth(tx.hash, Inc.utils.base64Encode);
            const issuingResponse = await transactor.shield({ transfer: { fee: 10, tokenID: tokenAddresses.pTokens[token] }, extra: { ethBlockHash: proveResult.ethBlockHash, ethDepositProof: proveResult.encodedProof, txIndex: proveResult.txIndex }});
            let change = await transactor.waitBalanceChange(tokenAddresses.pTokens[token]);
            console.log(`Detected balance change from ${BN.from(change.oldBalance)} to ${BN.from(change.balance)}`);
        }
        console.log('Done !');
        return tx;
    });

task("unshield", "Unshield (withdraw) from Incognito")
    .addOptionalParam("token", "Token to burn", "0x0000000000000000000000000000000000000000", types.string)
    .addOptionalParam("amount", "Amount to burn", "0", types.string)
    .addOptionalParam("to", "Receiver Ethereum address", "", types.string)
    .addOptionalParam("privkey", "Your private key (if you're submitting burn proof from this CLI)", "", types.string)
    .setAction(async taskArgs => {
        try {
            let { token, amount, to, privkey } = taskArgs;
            const { getInstance, toIncDecimals, transactor, Inc, inc, formatBurnProof, defaultSigner, BN, tokenAddresses, confirm, vault } = await run('contract-call-setup', { privkey });
            amount = BN.from(amount);
            if (to.length != 42) to = defaultSigner.address;
            const incAmount = await toIncDecimals(amount, token);
            const burnResult = await transactor.burn({ transfer: { fee: 10, tokenID: tokenAddresses.pTokens[token], info: "BURN TOKEN" }, extra: { remoteAddress: to, burnAmount: incAmount.toString(), burningType: Inc.constants.BurningRequestMeta }});
            await transactor.waitTx(burnResult.Response.txId, 16);
            const burnProof = await inc.rpc.getBurnProof(burnResult.Response.txId, false);
            let params = await formatBurnProof(burnProof);

            let getBalance = (token, addr) => token == tokenAddresses.ETH ? ethers.provider.getBalance(addr) : getInstance('contracts/IERC20.sol:IERC20', null, token).then(t => t.balanceOf(addr));
            const balanceBefore = await getBalance(token, to);
            const tx = await confirm(vault.connect(defaultSigner).withdraw(...params));
            const balance = await getBalance(token, to);
            console.log(`Receiver balance in Ethereum changed from ${balanceBefore.toString()} to ${balance.toString()}`);
            console.log('Done !');
            return tx;
        } catch (e) {
            console.error('Unshield failed')
            throw e;
        }
    });

task("query-var", "Display the vault's Incognito recovery address")
    .addOptionalParam("name", "name of public storage variable to get", "isInitialized")
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const res = await vault[taskArgs.name]();
        console.log(`Vault ${vault.address}: ${taskArgs.name} = ${res}`);
    });

task("testing-balance", "Display the testing accounts' token balances")
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const signers = await ethers.getSigners();
        const { testingTokenNames, deployed } = hre.networkCfg();
        const { tokenFunders, testingTokens } = deployed || {};  
        console.log(tokenFunders, testingTokens)      
        for (let i=0; i < testingTokens.length; i++) {
            console.log(testingTokenNames[i], '\t', testingTokens[i], 'balances');
            const token = await getInstance('contracts/IERC20.sol:IERC20', null, testingTokens[i]);
            for (const s of signers) {
                const b = await token.balanceOf(s.address);
                console.log('...', s.address, ':', b.toString());
            }
        }
    });

task("vault-evid", "Display the a vault's event ID (topic)")
    .addOptionalParam("name", "event name", "Deposit", types.string)
    .setAction(async taskArgs => {
        const { getInstance } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const res = await vault.interface.getEventTopic(taskArgs.name);
        console.log(res);
        // console.log(vault.interface.getEvent('0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d'));
    });

task("decode-event", "Decode event data by name")
    .addOptionalParam("name", "name", "Redeposit", types.string)
    .addOptionalParam("data", "data", "", types.string)
    .setAction(async (taskArgs, hre) => {
        const [signer, vaultAdmin] = await ethers.getSigners();
        const { getInstance, getImplementation, confirm } = require('./utils');
        const vault = await getInstance('Vault', 'TransparentUpgradeableProxy');
        const evdata = taskArgs.data; //'0x000000000000000000000000a6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000065763765a5767657bb9980d18a2f831bbcc90a877b4430119997951ddc7609d110000000000000000000000000000000000000000000000000000000000000065047ebadb13237d189484924e972a7f1b8f8e617773f909c69d4d6b19fc39c82e8455e85136c664f53d6c70a032fec17d6bb17ea344f8dc7f57173b2f42d77e468400000012ac36e69aa6c5403aa610702c5f986edbc5da92fd18c69fbd9ae96ebc6c13c0dc000000000000000000000000000000000000000000000000000000';
        res = vault.interface.decodeEventLog(taskArgs.name, evdata, [vault.interface.getEventTopic(taskArgs.name)]);
        console.log('decoded', taskArgs.name)
        console.dir(res);
        if (res.errorData) {
            // when errorData is available, parse it as string
            const s = '0x' + res.errorData.slice(10);
            console.log('error', ethers.utils.defaultAbiCoder.decode(['string'], s));
        }
    })
