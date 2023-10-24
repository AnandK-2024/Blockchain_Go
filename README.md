# Blockchain_Go

## Overview
This is a blockchain project developed from scratch using Golang. It includes various components such as blocks, transactions, account state, virtual machine, transaction pool, encoding/decoding, API, and ECDSA signature/verification. Additionally, I have implemented the TCP Ghost protocol for the transfer and receipt of transactions and blocks between peers in a peer-to-peer network.

## Features
- **Blocks**: The blockchain consists of a series of blocks, each containing a set of transactions.
- **Transactions**: Transactions represent the transfer of value between accounts on the blockchain.
- **Account State**: The current state of each account is maintained, reflecting the balance and other relevant information.
- **Virtual Machine**: A virtual machine is implemented to execute smart contracts and process transactions.
- **Transaction Pool**: A transaction pool holds pending transactions that are waiting to be included in a block.
- **Encoding/Decoding**: The project includes functionality to encode and decode data structures for efficient storage and transmission.
- **API**: An API is provided to interact with the blockchain, allowing users to submit transactions, query account balances, and retrieve block information.
- **ECDSA Signature/Verification**: The project includes the implementation of ECDSA (Elliptic Curve Digital Signature Algorithm) for secure transaction signing and verification.
- **TCP Ghost Protocol**: The TCP Ghost protocol is utilized for the efficient transfer and receipt of transactions and blocks between peers in a peer-to-peer network.

## Future Plans
In the future, I plan to make the blockchain compatible with native ERC-721 (NFT) and ERC-20 tokens. This will enable the transfer and management of these tokens on the blockchain.

## Getting Started
To get started with this project, follow the steps below:

1. Clone the repository: `git clone https://github.com/AnandK-2024/Blockchain_Go.git`
2. Install the required dependencies:
3. Build the project: `go build`
4. Run the blockchain node: `go run main.go`
5. Access the API endpoints to interact with the blockchain.

## API Endpoints
The following API endpoints are available for interacting with the blockchain:


- `GET /blocks/{blockHash or height}`: Retrieve a specific block by its hash or by its block number.
- `POST /tx`: Submit a new transaction to the blockchain.
- `GET /tx/{transactionHash}`: Retrieve a specific transaction by its hash.
- `GET /accounts/{accountAddress}`: Retrieve the account information for a specific address.

## Contributions
Contributions to this project are welcome. If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.



## Contact
If you have any questions or inquiries regarding this project, please feel free to contact me at [anandkumar@iitbhilai.ac.in](mailto:anandkumar@iitbhilai.ac.in).