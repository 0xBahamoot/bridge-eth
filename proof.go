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
	Error  struct {
		Code       int
		Message    string
		StackTrace string
	}
}

type decodedProof struct {
	Instruction  []byte
	BeaconHeight *big.Int
	BridgeHeight *big.Int

	BeaconInstPath       [inst_max_path][32]byte
	BeaconInstPathIsLeft [inst_max_path]bool
	BeaconInstPathLen    *big.Int
	BeaconInstRoot       [32]byte
	BeaconBlkData        [32]byte
	BeaconNumSig         *big.Int
	BeaconSigVs          [comm_size]*big.Int
	BeaconSigRs          [comm_size][32]byte
	BeaconSigSs          [comm_size][32]byte
	BeaconSigIdxs        [comm_size]*big.Int

	BridgeInstPath       [inst_max_path][32]byte
	BridgeInstPathIsLeft [inst_max_path]bool
	BridgeInstPathLen    *big.Int
	BridgeInstRoot       [32]byte
	BridgeBlkData        [32]byte
	BridgeNumSig         *big.Int
	BridgeSigVs          [comm_size]*big.Int
	BridgeSigRs          [comm_size][32]byte
	BridgeSigSs          [comm_size][32]byte
	BridgeSigIdxs        [comm_size]*big.Int
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

func getCommittee(url string) (*big.Int, []byte, *big.Int, []byte, error) {
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
		return nil, nil, nil, nil, err
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
		return nil, nil, nil, nil, err
	}

	// Genesis committee
	numBeaconVals := big.NewInt(int64(len(r.Result.BeaconCommittee)))
	beaconOld := []byte{}
	for i, val := range r.Result.BeaconCommittee {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		beaconOld = append(beaconOld, pk...)
	}

	numBridgeVals := big.NewInt(int64(len(r.Result.ShardCommittee["1"])))
	bridgeOld := []byte{}
	for i, val := range r.Result.ShardCommittee["1"] {
		pk, _, _ := base58.Base58Check{}.Decode(val)
		fmt.Printf("pk[%d]: %x %d\n", i, pk, len(pk))
		bridgeOld = append(bridgeOld, pk...)
	}

	return numBeaconVals, beaconOld, numBridgeVals, bridgeOld, nil
}

func getBurnProof(txID string) string {
	url := "http://127.0.0.1:9338"
	// url := "https://dev-test-node.incognito.org/"

	if len(txID) == 0 {
		txID = "87c89c1c19cec3061eff9cfefdcc531d9456ac48de568b3974c5b0a88d5f3834"
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
	fmt.Printf("beaconHeight: %d\n", beaconHeight)
	fmt.Printf("bridgeHeight: %d\n", bridgeHeight)

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
	// fmt.Printf("expected beaconBlkHash: %x\n", keccak256(beaconBlkData[:], beaconInstRoot[:]))

	// TODO(@0xbunyip): decode BeaconSigs to R, S and V

	beaconNumSig := big.NewInt(int64(len(r.Result.BeaconSigIdxs)))
	beaconSigIdxs := newBigIntArray()
	for i, sIdx := range r.Result.BeaconSigIdxs {
		beaconSigIdxs[i] = big.NewInt(int64(sIdx))
	}

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

	// TODO(@0xbunyip): decode BeaconSigs to R, S and V

	bridgeNumSig := big.NewInt(int64(len(r.Result.BridgeSigIdxs)))
	bridgeSigIdxs := newBigIntArray()
	for i, sIdx := range r.Result.BridgeSigIdxs {
		bridgeSigIdxs[i] = big.NewInt(int64(sIdx))
		// fmt.Printf("bridgeSigIdxs[%d]: %d\n", i, j)
	}

	return &decodedProof{
		Instruction: inst,

		BeaconHeight:         beaconHeight,
		BeaconInstPath:       beaconInstPath,
		BeaconInstPathIsLeft: beaconInstPathIsLeft,
		BeaconInstPathLen:    beaconInstPathLen,
		BeaconInstRoot:       beaconInstRoot,
		BeaconBlkData:        beaconBlkData,
		BeaconNumSig:         beaconNumSig,
		BeaconSigVs:          [comm_size]*big.Int{},
		BeaconSigRs:          [comm_size][32]byte{},
		BeaconSigSs:          [comm_size][32]byte{},
		BeaconSigIdxs:        beaconSigIdxs,

		BridgeHeight:         bridgeHeight,
		BridgeInstPath:       bridgeInstPath,
		BridgeInstPathIsLeft: bridgeInstPathIsLeft,
		BridgeInstPathLen:    bridgeInstPathLen,
		BridgeInstRoot:       bridgeInstRoot,
		BridgeBlkData:        bridgeBlkData,
		BridgeNumSig:         bridgeNumSig,
		BridgeSigVs:          [comm_size]*big.Int{},
		BridgeSigRs:          [comm_size][32]byte{},
		BridgeSigSs:          [comm_size][32]byte{},
		BridgeSigIdxs:        bridgeSigIdxs,
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

func newBigIntArray() [comm_size]*big.Int {
	arr := [comm_size]*big.Int{}
	for i := 0; i < comm_size; i++ {
		arr[i] = big.NewInt(0)
	}
	return arr
}

func getCommitteeHardcoded() (*big.Int, []byte, *big.Int, []byte) {
	beaconComm := []string{
		"02a96a04ad76a0034efc8819e93308823ce7a3b76fd694f961ee909124096baf00",
		"0242653de0e9af9dd3725008519157314eb5a845dec2cd646ce9e03f780175b700",
		"028c49fc5f3e001c36095335c53b0b7320f6a1c932424e92c9de344b55e80ddf00",
		"0205aae74cb0128a1863c970cbe87e827e28f92a91c2d4768fdb30a279dd081c00",
	}
	numBeaconVals := big.NewInt(int64(len(beaconComm)))
	beacons := []byte{}
	for _, p := range beaconComm {
		pk, _ := hex.DecodeString(p)
		beacons = append(beacons, pk...)
	}

	bridgeComm := []string{
		"0253d262c2b6a55606ff9d32e195231ec57e4d23a6efd1c02143a58fd0c2591d01",
		"02dee56cbbde5ef6d03a9e69bf3784ae4a8460d0058a6082eee4be2ed5c4fd3301",
		"02ec388db662801da0fe3c41f39085369ed4df71d42ec96924012243dc9c67d201",
		"039cc81f72a88a7436eb74bf10c7693af165324ba4d15baeb4e8d2f1c2ce25a101",
	}
	numBridgeVals := big.NewInt(int64(len(bridgeComm)))
	bridges := []byte{}
	for _, p := range bridgeComm {
		pk, _ := hex.DecodeString(p)
		bridges = append(bridges, pk...)
	}
	return numBeaconVals, beacons, numBridgeVals, bridges
}
