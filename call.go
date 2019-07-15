package bridge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/incognitochain/bridge-eth/vault"
)

func withdraw(v *vault.Vault, auth *bind.TransactOpts, proof *decodedProof) (*types.Transaction, error) {
	auth.GasLimit = 8000000
	tx, err := v.Withdraw(
		auth,
		proof.instruction,

		proof.beaconHeight,
		proof.beaconInstPath,
		proof.beaconInstPathIsLeft,
		proof.beaconInstPathLen,
		proof.beaconInstRoot,
		proof.beaconBlkData,
		proof.beaconBlkHash,
		proof.beaconSignerSig,
		proof.beaconNumR,
		proof.beaconXs,
		proof.beaconYs,
		proof.beaconRIdxs,
		proof.beaconNumSig,
		proof.beaconSigIdxs,
		proof.beaconRx,
		proof.beaconRy,
		proof.beaconR,

		proof.bridgeHeight,
		proof.bridgeInstPath,
		proof.bridgeInstPathIsLeft,
		proof.bridgeInstPathLen,
		proof.bridgeInstRoot,
		proof.bridgeBlkData,
		proof.bridgeBlkHash,
		proof.bridgeSignerSig,
		proof.bridgeNumR,
		proof.bridgeXs,
		proof.bridgeYs,
		proof.bridgeRIdxs,
		proof.bridgeNumSig,
		proof.bridgeSigIdxs,
		proof.bridgeRx,
		proof.bridgeRy,
		proof.bridgeR,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
