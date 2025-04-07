// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable {
    uint256 public tokenCounter;

    constructor(address initialOwner) ERC721("MyNFT", "MNFT") Ownable(initialOwner) {
        tokenCounter = 0;
    }

    function mintNFT(address recipient) public onlyOwner returns (uint256) {
        uint256 newItemId = tokenCounter;
        _safeMint(recipient, newItemId);
        tokenCounter++;
        return newItemId;
    }
}
