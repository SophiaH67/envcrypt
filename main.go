package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// # Envcrypt

// This project is a simple tool to encrypt and decrypt your .env files with your primary gpg key. This assumes that your .env file is called `.env` and that it is in the gitignore file. The encrypted file will be called `.env.asc` and should be committed to the repository.

// ## Usage

// After any changes to the `.env` file, run `envcrypt` to encrypt the file. This will create a `.env.asc` file that should be committed to the repository. When you pull the repository, run `envcrypt` to decrypt the file.

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}

	return false
}

func runCommand(args ...string) (string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()

	if err != nil {
		// Print command and args
		fmt.Println("Failed to run command:")
		fmt.Println(args)
		fmt.Println(err.Error())

		return "", err
	}

	return string(out), nil

}

func main() {
	// Check if gpg is installed
	_, err := exec.LookPath("gpg")

	if err != nil {
		fmt.Println("GPG is not installed")
		os.Exit(1)
	}

	envFileExists := exists(".env")
	envAscFileExists := exists(".env.asc")

	if !envFileExists && !envAscFileExists {
		fmt.Println("No .env or .env.asc file found")
		os.Exit(1)
	}

	if !envFileExists {
		// Decrypt
		_, err = runCommand("gpg", "-o", ".env", "-d", ".env.asc")

		if err != nil {
			fmt.Println("Could not decrypt .env.asc")
			os.Exit(1)
		}

		fmt.Println("Decrypted .env.asc")
		return
	}

	// Get the primary key configured in git
	gpgKeyId, err :=
		runCommand("git", "config", "--get", "user.signingkey")
	gpgKeyId = strings.TrimSpace(gpgKeyId)

	if err != nil || gpgKeyId == "" {
		fmt.Println("No signing key configured in git")
		os.Exit(1)
	}

	// Encrypt
	_, err = runCommand("gpg", "-o", ".env.asc", "-e", "-r", gpgKeyId, ".env")

	if err != nil {
		fmt.Println("Could not encrypt .env")
		os.Exit(1)
	}

	fmt.Println("Encrypted .env")
}
