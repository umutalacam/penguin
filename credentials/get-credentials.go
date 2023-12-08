package credentials

import (
	"fmt"
	"penguinCredentialManager/vault"
)

func ParseArgsGetCredentials(passphrase string, vaultFilePath string, args []string) {
	if len(args) < 2 {
		printGetCredentialsGuide()
	}
	credential, err := vault.GetCredential(passphrase, vaultFilePath, args[1])

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(credential)
	}
}

func printGetCredentialsGuide() {
	fmt.Print("Usage: ")
	fmt.Println("penguin get [credential-name]")
}
