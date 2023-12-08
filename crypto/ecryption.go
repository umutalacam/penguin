package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"penguinCredentialManager/model"
)

// CreateVaultFile Creates a vault file on the file system
func CreateVaultFile(vaultFilePath string, passphrase string) error {
	// Create data
	var vault = model.Vault{
		Version:    "0.0.1",
		Type:       "vault",
		Encryption: "aes",
		Items:      []*model.Entity{},
	}
	// Marshall data to json
	jsonData, err := json.Marshal(&vault)
	if err != nil {
		return errors.New("unable to encode vault to JSON")
	}
	// Convert to json string
	err = SaveVaultFileEncrypted(vaultFilePath, passphrase, string(jsonData))
	if err != nil {
		return err
	}
	return nil
}

func SaveVaultFileEncrypted(vaultFilePath string, passphrase string, content string) error {
	// Hash password to get key
	key := hashPassphrase(passphrase)
	// Decrypt data
	encrypted, err := encrypt([]byte(content), key[:])
	if err != nil {
		return err
	}
	// Save
	err = ioutil.WriteFile(vaultFilePath, encrypted, 0644)
	if err != nil {
		return errors.New("unable to save vault file")
	}
	return nil
}

func ReadVaultFile(vaultFilePath string, passphrase string) (string, error) {
	// Read file
	content, err := ioutil.ReadFile(vaultFilePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	// Hash password to get key
	key := hashPassphrase(passphrase)
	// Decrypt
	bytes, err := decrypt(content, key[:])
	if err != nil {
		return "", err
	}
	// Return
	return string(bytes), nil
}

func hashPassphrase(passphrase string) [16]byte {
	return md5.Sum([]byte(passphrase))
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(nonce, ciphertext...)
	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aesGCM.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:aesGCM.NonceSize()], ciphertext[aesGCM.NonceSize():]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
