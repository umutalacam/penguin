package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"penguinCredentialManager/credentials"
)

func main() {
	var arguments = os.Args
	ParseArguments(arguments)
}

// ParseArguments
// Command parser that invokes required module to fulfill user request
///**
func ParseArguments(args []string) {
	// Validate arguments
	if len(args) < 2 {
		printGuide()
		return
	}
	// Get user command
	command := args[1]
	var vaultFilePath = ".penguin/vault"
	var passphrase = input_passphrase(vaultFilePath)

	switch command {
	case "get":
		credentials.ParseArgsGetCredentials(passphrase, vaultFilePath, args[1:])
		return
	case "put":
		credentials.ParseArgsPutCredentials(passphrase, vaultFilePath, args[1:])
		return
	case "rm":
		credentials.ParseArgsRmCredentials(passphrase, vaultFilePath, args[1:])
		return
	case "init":
		credentials.ParseArgsInit(passphrase, vaultFilePath)
	case "-h":
		printIntroMessage()
		printGuide()
		return

	default:
		printGuide()
	}
}

func printIntroMessage() {
	fmt.Print(`
   _
 ('v')		Penguin credential store manager, simplifies the management of secrets,
//-=-\\		passwords with high security.
(\_=_/)
 ^^ ^^
`)
}

func printGuide() {
	fmt.Print("Usage: ")
	fmt.Println("penguin [command] [arguments]")
}

func input_passphrase(vaultFilePath string) string {
	fmt.Printf("Enter the passphrase for (%s):", vaultFilePath)
	var passPhrase, _ = terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Print("\n")
	return string(passPhrase)
}
