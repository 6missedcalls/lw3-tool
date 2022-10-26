# Scaffold CLI
Scaffold CLI is a tool designed to help scaffold basic boilerplate code for the Web3 Ecosystem.
## Supported Smart Contracts

- Ethereum ERC721 Smart Contract
- Ethereum ERC20 Smart Contract
## How to Build & Run Scaffold CLI

### Prerequisites
Go 1.17 or 1.18

### Build
```bash
go build -o scaffold
```

### Move to your GOPATH bin
```bash
mv scaffold $GOPATH/bin
```

### Run
```bash
scaffold
```

## CLI Commands
### Ethereum ERC721 Smart Contract
This will generate a basic ERC721 smart contract

Flags:
- `--name` or `-n` - Name of the smart contract
- `--ticker` or `-t` - Ticker of the smart contract

```bash
scaffold ethereum erc-721
```


### Ethereum ERC20 Smart Contract
This will generate a basic ERC20 smart contract

Flags:
- `--name` or `-n` - Name of the smart contract
- `--ticker` or `-t` - Ticker of the smart contract
- `--premint` or `-p` - Premint amount of the smart contract

```bash
scaffold ethereum erc-20
```
## Output of Contracts
All output as of now is hardcoded to the home directory of the user. In the home directory, there will be a folder called `.scaffold` and a subfolder `contracts` which will contain all the generated smart contracts.

## Contributing
All contributions are welcome!