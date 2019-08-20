
# External Contracts
contract Incognito_proxy:
    def extractMetaFromInst(inst: bytes[300], numPk: int128) -> (uint256, uint256, uint256): constant
    def extractCommitteeFromInst(inst: bytes[300], numVals: uint256) -> address[8]: constant
    def findCommitteeFromHeight(blkHeight: uint256, isBeacon: bool) -> uint256: constant
    def instructionInMerkleTree(leaf: bytes32, root: bytes32, path: bytes32[8], left: bool[8], length: int128) -> bool: constant
    def verifySig(signers: address[8], v: uint256[8], r: bytes32[8], s: bytes32[8], blk: bytes32) -> bool: constant
    def instructionApproved(isBeacon: bool, instHash: bytes32, blkHeight: uint256, instPath: bytes32[8], instPathIsLeft: bool[8], instPathLen: int128, instRoot: bytes32, blkData: bytes32, numSig: int128, sigIdxs: uint256[8], v: uint256[8], r: bytes32[8], s: bytes32[8]) -> bool: constant
    def swapBridgeCommittee(inst: bytes[300], numPk: int128, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconNumSig: int128, beaconSigIdxs: uint256[8], beaconSigVs: uint256[8], beaconSigRs: bytes32[8], beaconSigSs: bytes32[8], bridgeInstPath: bytes32[8], bridgeInstPathIsLeft: bool[8], bridgeInstPathLen: int128, bridgeInstRoot: bytes32, bridgeBlkData: bytes32, bridgeNumSig: int128, bridgeSigIdxs: uint256[8], bridgeSigVs: uint256[8], bridgeSigRs: bytes32[8], bridgeSigSs: bytes32[8]) -> bool: modifying
    def swapBeaconCommittee(inst: bytes[300], numPk: int128, beaconInstPath: bytes32[8], beaconInstPathIsLeft: bool[8], beaconInstPathLen: int128, beaconInstRoot: bytes32, beaconBlkData: bytes32, beaconNumSig: int128, beaconSigIdxs: uint256[8], beaconSigVs: uint256[8], beaconSigRs: bytes32[8], beaconSigSs: bytes32[8]) -> bool: modifying
    def beaconComm__Pubkeys(arg0: uint256, arg1: int128) -> address: constant
    def beaconComm__PrevBlk(arg0: uint256) -> uint256: constant
    def beaconComm__NumVals(arg0: uint256) -> uint256: constant
    def bridgeComm__Pubkeys(arg0: uint256, arg1: int128) -> address: constant
    def bridgeComm__PrevBlk(arg0: uint256) -> uint256: constant
    def bridgeComm__NumVals(arg0: uint256) -> uint256: constant
    def latestBeaconBlk() -> uint256: constant
    def latestBridgeBlk() -> uint256: constant
    def mulsig() -> address(MulSigP256): constant


