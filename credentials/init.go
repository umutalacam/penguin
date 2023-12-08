package credentials

import (
	"fmt"
	"penguinCredentialManager/crypto"
)

// ParseArgsInit Create vault file
func ParseArgsInit(passphrase string, vaultFilePath string) {
	err := crypto.CreateVaultFile(vaultFilePath, passphrase)
	fmt.Printf("The vault file created on %s\n", vaultFilePath)
	if err != nil {
		return
	}

}
