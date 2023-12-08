package vault

import (
	"encoding/json"
	"errors"
	"log"
	"penguinCredentialManager/crypto"
	"penguinCredentialManager/model"
	"strings"
)

func loadCredentialVault(passphrase string, vaultFilepath string) model.Vault {
	content, err := crypto.ReadVaultFile(vaultFilepath, passphrase)
	// Now let's unmarshall the data into `payload`
	var payload model.Vault
	err = json.Unmarshal([]byte(content), &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}

func saveCredentialVault(passPhrase string, vault *model.Vault, vaultFilepath string) error {
	// Marshal the vault structure to JSON
	jsonData, err := json.Marshal(&vault)
	if err != nil {
		return errors.New("unable to encode vault to JSON")
	}
	// Write the JSON data to the vault file
	err = crypto.SaveVaultFileEncrypted(vaultFilepath, passPhrase, string(jsonData))
	return nil
}

func GetCredential(passPhrase string, vaultFilePath string, credentialPath string) (credential string, err error) {
	var vault = loadCredentialVault(passPhrase, vaultFilePath)
	// Trim the slash
	trimmedCredentialPath := strings.Trim(credentialPath, "/")
	credentialPathParts := strings.Split(trimmedCredentialPath, "/")
	var credentialName = credentialPathParts[len(credentialPathParts)-1]
	// We need a root node
	var rootDirectory = model.Entity{
		Name:  "VaultRootEntity",
		Items: vault.Items,
	}
	// Iterate through collections
	var currentDirectory = rootDirectory
	for i := 0; i < len(credentialPathParts); i++ {
		// Find if there is a directory in this level
		directoryName := credentialPathParts[i]
		var entity, isFound = currentDirectory.GetEntity(directoryName)
		if !isFound {
			return credentialName, errors.New("no such entity or directory " + directoryName)
		}
		currentDirectory = *entity
	}
	// Return the collection data
	return dumpEntityJson(&currentDirectory)
}

func PutCredential(passphrase string, vaultFilePath string, credentialPath string, credentialValue string) error {
	var vault = loadCredentialVault(passphrase, vaultFilePath)
	var itemsPtr = &vault.Items
	// Trim the slash
	trimmedCredentialPath := strings.Trim(credentialPath, "/")
	credentialPathParts := strings.Split(trimmedCredentialPath, "/")
	var credentialName = credentialPathParts[len(credentialPathParts)-1]
	// Iterate through collections
	var rootDirectory = &model.Entity{
		Type:  "rootDirectory",
		Name:  "root",
		Value: nil,
		Items: *itemsPtr,
	}
	var currentDirectory = rootDirectory // /dev/ally-bros/superego/db_pass
	for i := 0; i < len(credentialPathParts)-1; i++ {
		// Move to the next level directory
		var directoryName = credentialPathParts[i]
		var entity, isFound = currentDirectory.GetEntity(directoryName)
		if !isFound {
			// Create a new directory is there is no entity
			currentDirectory = currentDirectory.CreateDirectoryEntity(directoryName)
		} else {
			if entity.Type != "directory" {
				return errors.New(directoryName + " not a directory")
			}
			currentDirectory = entity
		}
	}
	// Put the entity to current directory
	var entity = &model.Entity{
		Type:  "entity",
		Name:  credentialName,
		Value: &credentialValue,
		Items: nil,
	}
	vault.Items = rootDirectory.Items
	currentDirectory.CreateCredentialEntity(entity)
	err := saveCredentialVault(passphrase, &vault, vaultFilePath)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCredential(passphrase string, vaultFilePath string, credentialPath string) error {
	var vault = loadCredentialVault(passphrase, vaultFilePath)
	// Trim the slash
	trimmedCredentialPath := strings.Trim(credentialPath, "/")
	credentialPathParts := strings.Split(trimmedCredentialPath, "/")
	var credentialName = credentialPathParts[len(credentialPathParts)-1]
	// Iterate through collections
	currentDirectory := vault.GetRootEntity()
	// /dev/ally-bros/superego/db_pass
	for i := 0; i < len(credentialPathParts)-1; i++ {
		// Find if there is a directory in this level
		directoryName := credentialPathParts[i]
		var entity, isFound = currentDirectory.GetEntity(directoryName)
		if !isFound {
			return errors.New("no such entity or directory " + directoryName)
		}
		currentDirectory = entity
	}
	// Delete entity from current directory
	err := currentDirectory.DeleteEntity(credentialName)
	err = saveCredentialVault(passphrase, &vault, vaultFilePath)
	if err != nil {
		return err
	}
	return nil
}

func dumpEntityJson(entity *model.Entity) (collectionData string, err error) {
	jsonData, err := json.MarshalIndent(entity, "", "  ")
	if err != nil {
		return "", errors.New("unable to decode entity")
	}
	return string(jsonData), nil
}

func dumpVaultJson(vault *model.Vault) (vaultData string, err error) {
	// We found the credential
	jsonData, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return "", errors.New("unable to decode vault")
	}
	return string(jsonData), nil
}
