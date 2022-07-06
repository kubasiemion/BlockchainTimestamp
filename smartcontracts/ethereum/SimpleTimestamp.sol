// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title CertTimestamp
 * @dev Implements voting process along with vote delegation
 */
contract CertTime {
    struct Notarization {
        uint256 blockTime;
        uint256 blockNumber;
        uint256 CertID;
        bytes32 DocHash;
        address issuer;
        address Holder;
    }

    string public constant url =
        "https://researchbox1.uksouth.cloudapp.azure.com/certtime";

    mapping(uint256 => Notarization) public allCerts;
    mapping(uint8 => Notarization) public certsByIssue;

    /**
     *
     */
    constructor() {
        sequence = 1;
    }

    uint8 private sequence;

    function Sequence() public view returns (uint8) {
        return sequence;
    }

    function IssueCert(
        address holder,
        uint256 serial,
        bytes32 docHash
    ) public {
        allCerts[serial] = Notarization({
            issuer: msg.sender,
            blockTime: block.timestamp,
            blockNumber: block.number,
            Holder: holder,
            CertID: serial,
            DocHash: docHash
        });
        certsByIssue[sequence] = allCerts[serial];

        sequence++;
    }
}
