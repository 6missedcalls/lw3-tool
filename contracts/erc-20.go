package contracts

import "strings"

// This will write the solidity contract using the erc-20 specification
// https://eips.ethereum.org/EIPS/eip-20

// Create the basic contract
func BaseErc20(Name string, Ticker string, Premint string) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(
		`// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
		
contract ` + Name + ` is ERC20, ERC20Burnable, Pausable, Ownable {
	constructor() ERC20("` + Name + `", "` + Ticker + `") {
		_mint(msg.sender, ` + Premint + `* 10 ** decimals());
	}

	function pause() public onlyOwner {
		_pause();
	}

	function unpause() public onlyOwner {
		_unpause();
	}

	function mint(address to, uint256 amount) public onlyOwner {
		_mint(to, amount);
	}

	function _beforeTokenTransfer(address from, address to, uint256 amount)
		internal
		whenNotPaused
		override
	{
		super._beforeTokenTransfer(from, to, amount);
	}
}
`)
	return sb.String(), nil
}
