// SPDX-License-Identifier: MIT

pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;

contract VaultHelper {
    enum Prefix {
        EXECUTE_SIGNATURE,
        REQUEST_WITHDRAW_SIGNATURE
    }
    
    struct PreSignData {
        Prefix prefix;
        address token;
        bytes timestamp;
        uint amount;
    }

    function newPreSignData(Prefix prefix, address token, bytes calldata timestamp, uint amount) pure internal returns (PreSignData memory) {
        PreSignData memory psd = PreSignData(prefix, token, timestamp, amount);
        return psd;
    }

    function _buildPreSignData(Prefix prefix, address token, bytes calldata timestamp, uint amount) pure public {
        // PreSignData memory psd = PreSignData(prefix, token, timestamp, amount);
        // return psd;
    }

    function _buildSignRequestWithdraw(PreSignData memory psd, string calldata incognitoAddress) pure public returns (bytes memory) {
        // do nothing
        bytes memory temp = abi.encode(psd, incognitoAddress);
        return temp;
    }

    function _buildSignExecute(PreSignData memory psd, address recipientToken, address exchangeAddress, bytes calldata callData) pure public returns (bytes memory) {
        // do nothing
        bytes memory temp = abi.encode(psd, recipientToken, exchangeAddress, callData);
        return temp;
    }
}