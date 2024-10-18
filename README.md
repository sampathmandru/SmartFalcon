# Smart Falcon

## Overview

Smart Falcon is a blockchain-based project leveraging **Hyperledger Fabric** to manage and track assets for a financial institution. The system provides a secure, transparent, and immutable solution for asset management. The project encompasses chaincode development for peer interaction and a REST API for smart contract invocation.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Setup and Execution](#setup-and-execution)
- [Technologies Used](#technologies-used)
- [Testing](#testing)
- [License](#license)

## Project Structure

- `chaincode/`: Go code for Hyperledger Fabric chaincode
- `api/`: REST API code for blockchain interaction
- `docker/`: Docker configuration files
- Additional files: Configuration and utility scripts from Hyperledger open-source repository

## Setup and Execution

### 1. Environment Setup

The project utilizes **WSL** (Windows Subsystem for Linux) and **Docker** for containerization. Ensure both are properly installed and configured on your system.

### 2. Cloning the Repository

Initialize the project by cloning the Hyperledger Fabric samples repository:

```bash
git clone https://github.com/hyperledger/fabric-samples.git
```

This provides essential resources for setting up a test network.

### 3. Chaincode Development and Deployment

1. Develop the chaincode in Go, implementing asset management operations.
2. Package and install the chaincode on peer nodes.
3. Set up the chaincode path and verify peer communication.
4. Query the chaincode to retrieve the Package ID for correct path configuration.

### 4. REST API Development

1. Develop a REST API using Go to interact with the deployed chaincode.
2. Implement endpoints for creating, querying, and updating assets on the blockchain.

## Technologies Used

- **Hyperledger Fabric**: Blockchain framework
- **Go**: Programming language for chaincode and REST API
- **Docker**: Containerization platform
- **WSL** (Windows Subsystem for Linux): Development environment
- **Postman**: API testing tool

## Testing

- REST API: Thoroughly tested using Postman to ensure proper communication with the Hyperledger Fabric network.
- Chaincode: Validated using peer commands in the terminal for operations like asset creation and querying.

## License

This project is based on open-source resources provided by Hyperledger Fabric and adheres to its licensing terms. For detailed licensing information, please refer to the Hyperledger Fabric documentation.
