package bridge

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/checkMulSig"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/incognito_proxy"
	"github.com/incognitochain/bridge-eth/jsonresult"
	"github.com/incognitochain/bridge-eth/privacy"
	"github.com/incognitochain/bridge-eth/vault"
)

const inst_max_path = 8
const comm_size = 8
const pubkey_size = 33

type contracts struct {
	v         *vault.Vault
	vAddr     common.Address
	inc       *incognito_proxy.IncognitoProxy
	incAddr   common.Address
	token     *erc20.Erc20
	tokenAddr common.Address
	ms        *checkMulSig.MulSigP256
	msAddr    common.Address
}

type getProofResult struct {
	Result jsonresult.GetInstructionProof
	Error  string
	Id     int
}

type decodedProof struct {
	instruction  []byte
	beaconHeight *big.Int
	bridgeHeight *big.Int

	beaconInstPath       [inst_max_path][32]byte
	beaconInstPathIsLeft [inst_max_path]bool
	beaconInstPathLen    *big.Int
	beaconInstRoot       [32]byte
	beaconBlkData        [32]byte
	beaconBlkHash        [32]byte
	beaconSignerSig      *big.Int
	beaconNumR           *big.Int
	beaconXs             [comm_size]*big.Int
	beaconYs             [comm_size]*big.Int
	beaconRIdxs          [comm_size]*big.Int
	beaconNumSig         *big.Int
	beaconSigIdxs        [comm_size]*big.Int
	beaconRp             []byte
	beaconRpx            *big.Int
	beaconRpy            *big.Int
	beaconR              []byte

	bridgeInstPath       [inst_max_path][32]byte
	bridgeInstPathIsLeft [inst_max_path]bool
	bridgeInstPathLen    *big.Int
	bridgeInstRoot       [32]byte
	bridgeBlkData        [32]byte
	bridgeBlkHash        [32]byte
	bridgeSignerSig      *big.Int
	bridgeNumR           *big.Int
	bridgeXs             [comm_size]*big.Int
	bridgeYs             [comm_size]*big.Int
	bridgeRIdxs          [comm_size]*big.Int
	bridgeNumSig         *big.Int
	bridgeSigIdxs        [comm_size]*big.Int
	bridgeRp             []byte
	bridgeRpx            *big.Int
	bridgeRpy            *big.Int
	bridgeR              []byte
}

func getAndDecodeBurnProof(txID string) (*decodedProof, error) {
	body := getBurnProof(txID)
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof not found")
	}

	r := getProofResult{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeProof(&r)
}

func getCommittee(url string) ([]byte, []byte, error) {
	payload := strings.NewReader("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getbeaconbeststate\",\n    \"params\": []\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9334")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type beaconBestStateResult struct {
		BeaconCommittee []string
		ShardCommittee  map[string][]string
	}

	type getBeaconBestStateResult struct {
		Result beaconBestStateResult
		Error  string
		Id     int
	}

	r := getBeaconBestStateResult{}
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, nil, err
	}

	// Genesis committee
	beaconOld := []byte{}
	for i, val := range r.Result.BeaconCommittee {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		beaconOld = append(beaconOld, pk...)
	}

	bridgeOld := []byte{}
	for i, val := range r.Result.ShardCommittee["1"] {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		bridgeOld = append(bridgeOld, pk...)
	}

	return beaconOld, bridgeOld, nil
}

func getBurnProof(txID string) string {
	url := "http://127.0.0.1:9554"

	if len(txID) == 0 {
		txID = "fa2f44878614c5aee313ba5c7f81bef9eeb936aadff423c0497ab6e7eef17820"
	}
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getburnproof\",\n    \"params\": [\n    \t\"%s\"\n    ]\n}", txID))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func decodeProof(r *getProofResult) (*decodedProof, error) {
	inst := decode(r.Result.Instruction)
	fmt.Printf("inst: %d %x\n", len(inst), inst)

	// Block heights
	beaconHeight := big.NewInt(0).SetBytes(decode(r.Result.BeaconHeight))
	bridgeHeight := big.NewInt(0).SetBytes(decode(r.Result.BridgeHeight))

	beaconInstRoot := decode32(r.Result.BeaconInstRoot)
	beaconInstPath := [inst_max_path][32]byte{}
	beaconInstPathIsLeft := [inst_max_path]bool{}
	for i, path := range r.Result.BeaconInstPath {
		beaconInstPath[i] = decode32(path)
		beaconInstPathIsLeft[i] = r.Result.BeaconInstPathIsLeft[i]
	}
	beaconInstPathLen := big.NewInt(int64(len(r.Result.BeaconInstPath)))
	fmt.Printf("beaconInstRoot: %x\n", beaconInstRoot)

	beaconBlkData := toByte32(decode(r.Result.BeaconBlkData))
	beaconBlkHash := toByte32(decode(r.Result.BeaconBlkHash))
	fmt.Printf("expected beaconBlkHash: %x\n", keccak256(beaconBlkData[:], beaconInstRoot[:]))
	fmt.Printf("beaconBlkHash: %x\n\n", beaconBlkHash)

	beaconMulSig := &privacy.SchnMultiSig{}
	err := beaconMulSig.SetBytes(decode(r.Result.BeaconSignerSig))
	if err != nil {
		return nil, err
	}
	beaconSignerSig := beaconMulSig.S
	beaconNumR := big.NewInt(int64(len(r.Result.BeaconRIdxs)))
	beaconXs := newBigIntArray()
	beaconYs := newBigIntArray()
	beaconRIdxs := newBigIntArray()
	for i, rIdx := range r.Result.BeaconRIdxs {
		p, err := decompress(r.Result.BeaconPubkeys[rIdx])
		if err != nil {
			return nil, err
		}
		beaconXs[i] = p.X
		beaconYs[i] = p.Y
		beaconRIdxs[i] = big.NewInt(int64(rIdx))
	}
	beaconNumSig := big.NewInt(int64(len(r.Result.BeaconSigIdxs)))
	beaconSigIdxs := newBigIntArray()
	for i, sIdx := range r.Result.BeaconSigIdxs {
		j := findSigIdx(sIdx, r.Result.BeaconRIdxs)
		if j < 0 {
			return nil, fmt.Errorf("failed finding beacon sigIdx %d %v", sIdx, r.Result.BeaconRIdxs)
		}
		beaconSigIdxs[i] = big.NewInt(int64(j))
	}
	beaconRp := beaconMulSig.R.Compress()
	beaconRpx := beaconMulSig.R.X
	beaconRpy := beaconMulSig.R.Y
	beaconR := decode(r.Result.BeaconR)

	// For bridge
	bridgeInstRoot := decode32(r.Result.BridgeInstRoot)
	bridgeInstPath := [inst_max_path][32]byte{}
	bridgeInstPathIsLeft := [inst_max_path]bool{}
	for i, path := range r.Result.BridgeInstPath {
		bridgeInstPath[i] = decode32(path)
		bridgeInstPathIsLeft[i] = r.Result.BridgeInstPathIsLeft[i]
	}
	bridgeInstPathLen := big.NewInt(int64(len(r.Result.BridgeInstPath)))
	// fmt.Printf("bridgeInstRoot: %x\n", bridgeInstRoot)

	bridgeBlkData := toByte32(decode(r.Result.BridgeBlkData))
	bridgeBlkHash := toByte32(decode(r.Result.BridgeBlkHash))
	// fmt.Printf("bridgeBlkHash: %x\n", bridgeBlkHash)

	bridgeMulSig := &privacy.SchnMultiSig{}
	err = bridgeMulSig.SetBytes(decode(r.Result.BridgeSignerSig))
	if err != nil {
		return nil, err
	}
	bridgeSignerSig := bridgeMulSig.S
	bridgeNumR := big.NewInt(int64(len(r.Result.BridgeRIdxs)))
	bridgeXs := newBigIntArray()
	bridgeYs := newBigIntArray()
	bridgeRIdxs := newBigIntArray()
	for i, rIdx := range r.Result.BridgeRIdxs {
		p, err := decompress(r.Result.BridgePubkeys[rIdx])
		if err != nil {
			return nil, err
		}
		bridgeXs[i] = p.X
		bridgeYs[i] = p.Y
		bridgeRIdxs[i] = big.NewInt(int64(rIdx))
	}
	bridgeNumSig := big.NewInt(int64(len(r.Result.BridgeSigIdxs)))
	bridgeSigIdxs := newBigIntArray()
	for i, sIdx := range r.Result.BridgeSigIdxs {
		j := findSigIdx(sIdx, r.Result.BridgeRIdxs)
		if j < 0 {
			return nil, fmt.Errorf("failed finding bridge sigIdx %d %v", sIdx, r.Result.BridgeRIdxs)
		}
		bridgeSigIdxs[i] = big.NewInt(int64(j))
		fmt.Printf("bridgeSigIdxs[%d]: %d\n", i, j)
	}
	bridgeRp := bridgeMulSig.R.Compress()
	bridgeRpx := bridgeMulSig.R.X
	bridgeRpy := bridgeMulSig.R.Y
	bridgeR := decode(r.Result.BridgeR)

	return &decodedProof{
		instruction: inst,

		beaconHeight:         beaconHeight,
		beaconInstPath:       beaconInstPath,
		beaconInstPathIsLeft: beaconInstPathIsLeft,
		beaconInstPathLen:    beaconInstPathLen,
		beaconInstRoot:       beaconInstRoot,
		beaconBlkData:        beaconBlkData,
		beaconBlkHash:        beaconBlkHash,
		beaconSignerSig:      beaconSignerSig,
		beaconNumR:           beaconNumR,
		beaconXs:             beaconXs,
		beaconYs:             beaconYs,
		beaconRIdxs:          beaconRIdxs,
		beaconNumSig:         beaconNumSig,
		beaconSigIdxs:        beaconSigIdxs,
		beaconRp:             beaconRp,
		beaconRpx:            beaconRpx,
		beaconRpy:            beaconRpy,
		beaconR:              beaconR,

		bridgeHeight:         bridgeHeight,
		bridgeInstPath:       bridgeInstPath,
		bridgeInstPathIsLeft: bridgeInstPathIsLeft,
		bridgeInstPathLen:    bridgeInstPathLen,
		bridgeInstRoot:       bridgeInstRoot,
		bridgeBlkData:        bridgeBlkData,
		bridgeBlkHash:        bridgeBlkHash,
		bridgeSignerSig:      bridgeSignerSig,
		bridgeNumR:           bridgeNumR,
		bridgeXs:             bridgeXs,
		bridgeYs:             bridgeYs,
		bridgeRIdxs:          bridgeRIdxs,
		bridgeNumSig:         bridgeNumSig,
		bridgeSigIdxs:        bridgeSigIdxs,
		bridgeRp:             bridgeRp,
		bridgeRpx:            bridgeRpx,
		bridgeRpy:            bridgeRpy,
		bridgeR:              bridgeR,
	}, nil
}

func toByte32(s []byte) [32]byte {
	a := [32]byte{}
	copy(a[:], s)
	return a
}

func decode(s string) []byte {
	d, _ := hex.DecodeString(s)
	return d
}

func decode32(s string) [32]byte {
	return toByte32(decode(s))
}

func keccak256(b ...[]byte) [32]byte {
	h := crypto.Keccak256(b...)
	r := [32]byte{}
	copy(r[:], h)
	return r
}

func decompress(s string) (*privacy.EllipticPoint, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	p := &privacy.EllipticPoint{}
	err = p.Decompress(b)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func findSigIdx(sIdx int, rIdxs []int) int {
	for j, rIdx := range rIdxs {
		if rIdx == sIdx {
			return j
		}
	}
	return -1
}

func newBigIntArray() [comm_size]*big.Int {
	arr := [comm_size]*big.Int{}
	for i := 0; i < comm_size; i++ {
		arr[i] = big.NewInt(0)
	}
	return arr
}
