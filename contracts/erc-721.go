package contracts

import "strings"

// This will write the solidity contract using the erc-721 specification
// https://eips.ethereum.org/EIPS/eip-721

// Create the basic contract
func BaseErc721(Name string, Ticker string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(
		`
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ` + Name + ` is ERC721, Ownable {
using Counters for Counters.Counter;

Counters.Counter private _tokenIds;

constructor() ERC721("` + Name + `", "` + Ticker + `") {}

function pause() public onlyOwner {
  _pause();
}

function unpause() public onlyOwner {
	_unpause();
}

function safeMint(address to) public onlyOwner {
	uint256 tokenId = _tokenIdCounter.current();
	_tokenIdCounter.increment();
	_safeMint(to, tokenId);
}

function _beforeTokenTransfer(address from, address to, uint256 tokenId)
	internal
	whenNotPaused
	override(ERC721, ERC721Enumerable)
	{
		super._beforeTokenTransfer(from, to, tokenId);
	}

// The following functions are overrides required by Solidity.
function supportsInterface(bytes4 interfaceId)
	public
	view
	override(ERC721, ERC721Enumerable)
	returns (bool)
	{
		return super.supportsInterface(interfaceId);
	}
}
`)
	return sb.String(), nil
}
