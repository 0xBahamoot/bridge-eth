package bridge

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/incognito_proxy"
	"github.com/incognitochain/bridge-eth/vault"
)

func TestSimulatedErc20(t *testing.T) {
	proof, err := getAndDecodeBurnProof("")
	if err != nil {
		t.Fatal(err)
	}

	sim, _, _, v, vAddr, token, _, err := setupWithErc20()
	if err != nil {
		t.Fatal(err)
	}

	oldBalance, newBalance, err := depositErc20(sim, token, vAddr, int64(1e9))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("deposit erc20 to vault: %d -> %d\n", oldBalance, newBalance)

	withdrawer := common.HexToAddress("0x0FFBd68F130809BcA7b32D9536c8339E9A844620")
	fmt.Printf("withdrawer init balance: %d\n", getBalanceErc20(token, withdrawer))

	auth.GasLimit = 6000000
	tx, err := v.Withdraw(
		auth,
		proof.instruction,

		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
		proof.beaconInstRoot,
		proof.beaconBlkData,
		proof.beaconBlkHash,
		proof.beaconSignerPubkeys,
		proof.beaconSignerCount,
		proof.beaconSignerSig,
		proof.beaconSignerPaths,
		proof.beaconSignerPathIsLeft,
		proof.beaconSignerPathLen,

		proof.bridgeInstPath,
		proof.bridgeInstPathIsLeft,
		proof.bridgeInstPathLen,
		proof.bridgeInstRoot,
		proof.bridgeBlkData,
		proof.bridgeBlkHash,
		proof.bridgeSignerPubkeys,
		proof.bridgeSignerCount,
		proof.bridgeSignerSig,
		proof.bridgeSignerPaths,
		proof.bridgeSignerPathIsLeft,
		proof.bridgeSignerPathLen,
	)
	if err != nil {
		t.Fatal(err)
	}
	sim.Commit()
	fmt.Printf("withdrawer new balance: %d\n", getBalanceErc20(token, withdrawer))
	printReceipt(sim, tx)
}

func getBalanceErc20(token *erc20.Erc20, addr common.Address) *big.Int {
	bal, _ := token.BalanceOf(nil, addr)
	return bal
}

func depositErc20(
	sim *backends.SimulatedBackend,
	token *erc20.Erc20,
	to common.Address,
	amount int64,
) (*big.Int, *big.Int, error) {
	initBalance := getBalanceErc20(token, to)
	auth := bind.NewKeyedTransactor(genesisAcc.PrivateKey)
	_, err := token.Transfer(auth, to, big.NewInt(amount))
	if err != nil {
		return nil, nil, err
	}
	sim.Commit()
	newBalance := getBalanceErc20(token, to)
	return initBalance, newBalance, nil
}

func setupWithErc20() (
	sim *backends.SimulatedBackend,
	inc *incognito_proxy.IncognitoProxy,
	incAddr common.Address,
	v *vault.Vault,
	vAddr common.Address,
	token *erc20.Erc20,
	tokenAddr common.Address,
	err error,
) {
	p := &Platform{}
	p, err = setupWithCommittee()
	if err != nil {
		return
	}

	token, tokenAddr, err = deployErc20(p.sim)
	if err != nil {
		return
	}

	return p.sim, p.inc, p.incAddr, p.vault, p.vaultAddr, token, tokenAddr, nil
}

func deployErc20(sim *backends.SimulatedBackend) (*erc20.Erc20, common.Address, error) {
	name := "Super duper erc20"
	symbol := "="
	decimals := big.NewInt(0)
	supply := big.NewInt(1000000000000000000)
	addr, _, c, err := erc20.DeployErc20(auth, sim, name, symbol, decimals, supply)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to deploy Erc20 contract: %v", err)
	}
	sim.Commit()
	fmt.Printf("deployed erc20, addr: %x\n", addr)
	fmt.Printf("genesis address erc20 balance: %d\n", getBalanceErc20(c, genesisAcc.Address))
	return c, addr, nil
}
