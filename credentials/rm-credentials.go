package credentials

import (
	"fmt"
	"penguinCredentialManager/vault"
)

// ParseArgsRmCredentials ParseArgsPutCredentials penguin put /credential/name value
func ParseArgsRmCredentials(passphrase string, vaultFilePath string, args []string) {
	if len(args) < 2 {
		printRmCredentialsGuide()
		return
	}
	// Delete credentials
	credentialPath := args[1]
	err := vault.DeleteCredential(passphrase, vaultFilePath, credentialPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func printRmCredentialsGuide() {
	fmt.Print("Usage: ")
	fmt.Println("penguin rm [credential-name]")
}
