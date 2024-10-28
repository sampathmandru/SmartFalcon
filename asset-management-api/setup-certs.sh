#!/bin/bash

# Get the absolute path of the current directory
API_ROOT=$(pwd)
NETWORK_ROOT="../test-network"

# Create config directory
echo "Creating config directory..."
mkdir -p "${API_ROOT}/config"

echo "Copying certificates and connection profile..."

# Copy Admin certificates
cp "${NETWORK_ROOT}/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/"* "${API_ROOT}/config/org1.example.com-cert.pem"

# Copy Admin private key
cp "${NETWORK_ROOT}/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/"* "${API_ROOT}/config/org1.example.com-key.pem"

# Copy the TLS CA cert
cp "${NETWORK_ROOT}/organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem" "${API_ROOT}/config/ca.crt"

# Create the connection profile
cat > "${API_ROOT}/config/connection-org1.yaml" << EOL
---
name: test-network-org1
version: 1.0.0
client:
  organization: Org1
  connection:
    timeout:
      peer:
        endorser: '300'
organizations:
  Org1:
    mspid: Org1MSP
    peers:
    - peer0.org1.example.com
    certificateAuthorities:
    - ca.org1.example.com
peers:
  peer0.org1.example.com:
    url: grpcs://localhost:7051
    tlsCACerts:
      path: ${API_ROOT}/config/ca.crt
    grpcOptions:
      ssl-target-name-override: peer0.org1.example.com
      hostnameOverride: peer0.org1.example.com
certificateAuthorities:
  ca.org1.example.com:
    url: https://localhost:7054
    caName: ca-org1
    tlsCACerts:
      path: ${API_ROOT}/config/ca.crt
    httpOptions:
      verify: false
EOL

# Set permissions
chmod 644 "${API_ROOT}/config/"*

echo "Verifying files..."
ls -l "${API_ROOT}/config/"

if [ $? -eq 0 ]; then
    echo "Setup completed successfully!"
    echo "Certificate files are in: ${API_ROOT}/config/"
else
    echo "Setup failed!"
    exit 1
fi