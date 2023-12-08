package credentials

import (
	"fmt"
	"penguinCredentialManager/vault"
)

// ParseArgsPutCredentials penguin put /credential/name value
func ParseArgsPutCredentials(passphrase string, vaultFilePath string, args []string) {
	if len(args) < 3 {
		printPutCredentialsGuide()
		return
	}
	// Put credentials
	credentialPath := args[1]
	credentialValue := args[2]
	err := vault.PutCredential(passphrase, vaultFilePath, credentialPath, credentialValue)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func printPutCredentialsGuide() {
	fmt.Print("Usage: ")
	fmt.Println("penguin put [credential-name] [value]")
}
