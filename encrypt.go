package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"strings"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

const (
	// secret length must be 16, 24 or 32, corresponding to the AES-128, AES-192 or AES-256 algorithms
	secret = "github.com/adibendahan/mysqlbeat"
)

func main() {

	secretLength := len(secret)
	if secretLength != 16 && secretLength != 24 && secretLength != 32 {
		fmt.Printf("Error: `secret` length must be 16, 24 or 32, corresponding to the AES-128, AES-192 or AES-256 algorithms.\n")
		os.Exit(-1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password to encrypt: ")
	plaintext, _ := reader.ReadBytes('\n')
	plaintext = []byte(strings.Trim(string(plaintext), "\n"))

	// Create the aes encryption algorithm
	c, err := aes.NewCipher([]byte(secret))
	if err != nil {

		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}

	// Encrypted string
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Encrypted password: %x\n", ciphertext)

}
