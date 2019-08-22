const ECDSA = artifacts.require("ECDSA");

contract("ECDSA", () => {
    before(async() => {
        contract = await ECDSA.deployed()
    })
    describe('Test verify ecdsa generated by golang', () => {
        let msgHash = '0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8'; // 32 bytes
        let address = ['0xcb859a5fC20EEeCc4Cec191d8CCe5e31a2CC1dAF','0x6294A3d1caE192f06dEe33152559531D447A076f','0xA815f542096c8De2ECe5aB18d4cf9D2aBc5202EC','0x4B895A89606aD73d2Fd7b887583858d6f2Cd229c','0xdd0523326fD818a16783D392324003D1df163758','0xf5b0A7D1270642e55a92A99D1AF9bb2B59982C71','0xbEBE7795d8297c4A59542a81541878e2dBA95253','0x894a0bEbb56cE3099A34f26b259D4038AE0E6088','0x6D6abB339E215a92c190f045D88E8aF79A32Dd16','0x81a3B54a6216585C6A262AAF2c4340Ac53F7c10c']
        let s = ['0x4f8daa779f713e539bbe8942e390325fc360a4811b0a568fbcfd9c2547f612d9','0x48c8fb9529c8de1d9097366747c555847f29e73838c36c33eec1fcc266441fc1','0x1e81121869b2d39c39e0d5bce27a233e7488c254c4119b8455241d89b17090dc','0x1c8060da0d3650501b6c7c7d4940d224dbc64eab4d8beae4b6358f4b9b538e2a','0x06a086905048b493cbefc4d64fe1ee468836ffec67a81fcd205a92c42667830b','0x5fbcab3fbe8e8e5a978dcfbae54f68917f8056e802cebc2479846413ec054e97','0x280704f0e0f70f1aa43ee9e0fc21fc09edafa00ec29cebff696c504e7fa7c43b','0x7fdc885e553ae521b0879cb86c3891e5cdaf4bad6cd1fd29f2482fa282f93978','0x6ee340fa29ed5b41285297a7bc0451881a4aad138fa6fa43a7b507d58561106d','0x7c5b3abc8224dc35982a1f9dfd34d9a18862aad8f54ed2d905b6e81c7fbcf8a7'];
        let r = ['0x012305d4c970ee1872d87b9c8de53bfa698782222bb318861690f9863f4c5a2a','0x03e6a2f1af4613c257eea0b11977d482f22969335504ed55ff07e3225e677fcc','0x86e2b5df59171a87b94e38af69474eb97c3a741fc6e3d488e3e27d63da44612f','0xd7deea8a087d4344bcb39f897de886be7020a22518b177b3541e2cb81d10bf80','0x778bf6835438261ba8485774aec7c44b6239c30a9e68f74aa417af673df31d35','0x98f4a9d71f60ea86321f84e4ca6cce1ef5827e2ae8e22ad1944d70125820a626','0xcba676c98e3489e27ee9d37d4ec39239f60523390f30c30ad6b4feb60eacc5e3','0xbf5e08d182998ec31f19a8d3744be78d5025b7fb33f87b4e71ec1d797cf92ddc','0xd6f28bc51138223f0e6620f42e4fba1159130d4dfa136e7ab35ff4081e4cee32','0x62e9fd38a63b0180a41f777a30717b198782dc8e43ebec3281a3f0e5e6aca069'];
        let v = ['27','28','28','27','27','27','27','28','28','28']; // uint8

        it.only('Test verify valid signature gas', async() => {
            // const gasUsed = await c.checkMulSig.estimateGas(pks,xPks,yPks,idxr,idxs,xR,yR,R,sig,mess)
            // txs = await c.scalarMul(xPks[0],xPks[1],xPks[2],xPks[3])
            let muladdr = []
            let mulr = []
            let muls = []
            let mulv = []
            for (i=0;i<30;i++){
                muladdr = muladdr.concat(address)
                mulr = mulr.concat(r)
                muls = muls.concat(s)
                mulv = mulv.concat(v)
            }
            console.log(muladdr.length)
            const result = await contract.verify(muladdr, msgHash, mulv, mulr, muls);
            console.log(result);
            const gasUsed = await contract.verify.estimateGas(muladdr, msgHash, mulv, mulr, muls);
            console.log(gasUsed);
        });

        it.only('Test verify invalid signature gas', async() => {
            // const gasUsed = await c.checkMulSig.estimateGas(pks,xPks,yPks,idxr,idxs,xR,yR,R,sig,mess)
            // txs = await c.scalarMul(xPks[0],xPks[1],xPks[2],xPks[3])
            console.log(v.length);
            const result = await contract.verify(address, msgHash, v, r, s);
            console.log(result);
            const gasUsed = await contract.verify.estimateGas(address, msgHash, v, r, s);
            console.log(gasUsed);
        });
    });
});

//[4 121 190 102 126 249 220 187 172 85 160 98 149 206 135 11 7 2 155 252 219 45 206 40 217 89 242 129 91 22 248 23 152 183 197 37 136 217 92 59 154 162 91 4 3 241 238 247 87 2 232 75 183 89 122 171 230 99 184 47 111 4 239 39 119]