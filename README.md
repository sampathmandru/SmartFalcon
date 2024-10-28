# Smart Falcon
## Overview
Smart Falcon is a blockchain-based project leveraging **Hyperledger Fabric** to manage and track assets for a financial institution. The system provides a secure, transparent, and immutable solution for asset management. The project encompasses chaincode development for peer interaction and a REST API for smart contract invocation.

## Table of Contents
- [Overview](#overview)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
- [Channel Creation](#channel-creation)
- [Chaincode Deployment](#chaincode-deployment)
- [API Setup](#api-setup)
- [Technologies Used](#technologies-used)
- [Testing](#testing)
- [License](#license)

## Project Structure
- `chaincode/asset-management/`: Go code for Hyperledger Fabric chaincode
- `asset-management-api/`: REST API code for blockchain interaction
- `docker/`: Docker configuration files
- `test-network/`: Network configuration and channel setup scripts
- Additional files: Configuration and utility scripts from Hyperledger open-source repository

## Setup and Installation
### 1. Environment Setup
The project utilizes **WSL** (Windows Subsystem for Linux) and **Docker** for containerization. Ensure both are properly installed and configured on your system.

### 2. Repository Setup
1. Clone the Hyperledger Fabric samples repository:
```bash
git clone https://github.com/hyperledger/fabric-samples.git
cd fabric-samples
```

2. Create your project directory and initialize your custom chaincode:
```bash
mkdir chaincode/asset-management
cd chaincode/asset-management
```

## Channel Creation
1. Navigate to the test-network directory:
```bash
cd test-network
```

2. Start the network and create a channel:
```bash
./network.sh up
./network.sh createChannel -c mychannel
```

3. Verify channel creation:
- Check Docker Desktop to ensure all containers are running
- You should see containers for:
  - Peers (peer0.org1, peer0.org2)
  - Orderer
  - CLI

## Chaincode Deployment
1. Implement your chaincode in the `chaincode/asset-management` directory:
```bash
# Add your Go chaincode files here:
# - asset_management.go
# - go.mod
# - go.sum
```

2. Package the chaincode:
```bash
peer lifecycle chaincode package asset-management.tar.gz --path ./chaincode/asset-management --lang golang --label asset-management_1.0
```

3. Install the chaincode package:
```bash
peer lifecycle chaincode install asset-management.tar.gz
```

4. Query the package ID:
```bash
peer lifecycle chaincode queryinstalled
```
Note: Save the package ID for later use (format: asset-management_1.0:hash)

5. Approve and commit the chaincode:
```bash
peer lifecycle chaincode approveformyorg -o localhost:7050 --channelID mychannel --name asset-management --version 1.0 --package-id PACKAGE_ID --sequence 1
peer lifecycle chaincode commit -o localhost:7050 --channelID mychannel --name asset-management --version 1.0 --sequence 1
```

## API Setup
1. Navigate to the API directory:
```bash
cd ../asset-management-api
```

2. Run the API server:
```bash
go run main.go
```

The API will now be available to communicate with the chaincode through the peer network.

## Technologies Used
- **Hyperledger Fabric**: Blockchain framework
- **Go**: Programming language for chaincode and REST API
- **Docker**: Containerization platform
- **WSL** (Windows Subsystem for Linux): Development environment
- **Postman**: API testing tool

## Testing
### API Testing with Postman
1. Import the provided Postman collection
2. Test the following endpoints:
   - POST /assets: Create new assets
   - GET /assets/{id}: Query existing assets
   - PUT /assets/{id}: Update asset information
   - DELETE /assets/{id}: Delete assets

### Chaincode Testing
1. Use peer commands to test chaincode functionality:
```bash
peer chaincode invoke -C mychannel -n asset-management -c '{"Args":["CreateAsset","asset1","value1"]}'
peer chaincode query -C mychannel -n asset-management -c '{"Args":["ReadAsset","asset1"]}'
```

## License
This project is based on open-source resources provided by Hyperledger Fabric and adheres to its licensing terms. For detailed licensing information, please refer to the Hyperledger Fabric documentation.
