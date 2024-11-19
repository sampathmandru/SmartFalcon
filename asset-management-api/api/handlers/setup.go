package handlers

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type AssetManagementAPI struct {
	Contract *gateway.Contract
}

func GetProjectRoot() (string, error) {
    // Get the absolute path to the current working directory
    currentDir, err := os.Getwd()
    if err != nil {
        return "", fmt.Errorf("failed to get current working directory: %v", err)
    }
    
    // Since we know we're running from the project root, use the current directory
    return currentDir, nil
}

func verifySetup() error {
    projectRoot, err := GetProjectRoot()
    if err != nil {
        return err
    }
    
    configPath := filepath.Join(projectRoot, "config")
    log.Printf("Verifying setup in config directory: %s", configPath)
    
    // Check if config directory exists
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        return fmt.Errorf("config directory not found at: %s", configPath)
    }

    // List of required files
    requiredFiles := []string{
        "org1.example.com-cert.pem",
        "org1.example.com-key.pem",
        "ca.crt",
        "connection-org1.yaml",
    }

    // Check each required file
    for _, file := range requiredFiles {
        filePath := filepath.Join(configPath, file)
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            return fmt.Errorf("required file not found: %s", file)
        }
        
        // Read file to ensure it's not empty
        content, err := os.ReadFile(filePath)
        if err != nil {
            return fmt.Errorf("failed to read file %s: %v", file, err)
        }
        if len(content) == 0 {
            return fmt.Errorf("file is empty: %s", file)
        }
    }

    log.Printf("All required files found and verified in: %s", configPath)
    return nil
}

func populateWallet(wallet *gateway.Wallet) error {
    projectRoot, err := GetProjectRoot()
    if err != nil {
        return err
    }
    
    credPath := filepath.Join(projectRoot, "config")
    log.Printf("Looking for credentials in: %s", credPath)

    // Read the certificate
    certPath := filepath.Join(credPath, "org1.example.com-cert.pem")
    cert, err := os.ReadFile(filepath.Clean(certPath))
    if err != nil {
        return fmt.Errorf("failed to read cert file from %s: %v", certPath, err)
    }

    // Read the private key
    keyPath := filepath.Join(credPath, "org1.example.com-key.pem")
    key, err := os.ReadFile(filepath.Clean(keyPath))
    if err != nil {
        return fmt.Errorf("failed to read key file from %s: %v", keyPath, err)
    }

    // Create identity
    log.Printf("Creating identity for Org1MSP...")
    identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

    // Store in wallet
    if err = wallet.Put("org1", identity); err != nil {
        return fmt.Errorf("failed to put identity in wallet: %v", err)
    }

    log.Printf("Successfully populated wallet with identity for Org1MSP")
    return nil
}

func SetupAPI() (*AssetManagementAPI, error) {
    // Verify setup first
    if err := verifySetup(); err != nil {
        return nil, fmt.Errorf("setup verification failed: %v", err)
    }

    projectRoot, err := GetProjectRoot()
    if err != nil {
        return nil, err
    }
    
    // Create wallet in the project root
    wallet, err := gateway.NewFileSystemWallet(filepath.Join(projectRoot, "wallet"))
    if err != nil {
        return nil, fmt.Errorf("failed to create wallet: %v", err)
    }

    if !wallet.Exists("org1") {
        err = populateWallet(wallet)
        if err != nil {
            return nil, fmt.Errorf("failed to populate wallet contents: %v", err)
        }
    }

    // Use the project root for config path
    configPath := filepath.Join(projectRoot, "config", "connection-org1.yaml")
    log.Printf("Using connection profile at: %s", configPath)

    gw, err := gateway.Connect(
        gateway.WithConfig(config.FromFile(filepath.Clean(configPath))),
        gateway.WithIdentity(wallet, "org1"),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to connect to gateway: %v", err)
    }

    network, err := gw.GetNetwork("mychannel")
    if err != nil {
        return nil, fmt.Errorf("failed to get network: %v", err)
    }

    contract := network.GetContract("asset_management")

    api := &AssetManagementAPI{
        Contract: contract,
    }

    return api, nil
}