package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	// The text to be encrypted
	// plainText ,err := ioutil.ReadFile("ooo.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	plainText := "Hello, World!"

	// file, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	encryptedText ,err := ioutil.ReadFile("log.txt")


	// encryptedText, err:= base64.StdEncoding.DecodeString("zQIUl7cJP3DuhYrPgE4pKw==")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Generate a random 32-byte key
	key := []byte("e10adc3949ba59abbe56e057f20f883e")

	// Create a new AES cipher block using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Generate a random IV (initialization vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}

	// Encrypt the plaintext
	ciphertext := encrypt(block, iv, []byte(plainText))
	// _,err = file.WriteString(string(ciphertext))
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Encrypted:", base64.StdEncoding.EncodeToString(ciphertext))

	// Decrypt the ciphertext
	decryptedText := decrypt(block, iv, encryptedText)
	fmt.Println("Decrypted:", string(decryptedText))
}

func encrypt(block cipher.Block, iv, plaintext []byte) []byte {
	// Create a new AES cipher block mode for encryption
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad the plaintext to a multiple of the block size
	paddedText := pad(plaintext, aes.BlockSize)

	// Create a buffer for the ciphertext
	ciphertext := make([]byte, len(paddedText))

	// Encrypt the padded plaintext
	mode.CryptBlocks(ciphertext, paddedText)

	return ciphertext
}

func decrypt(block cipher.Block, iv, ciphertext []byte) []byte {
	// Create a new AES cipher block mode for decryption
	mode := cipher.NewCBCDecrypter(block, iv)

	// Create a buffer for the padded plaintext
	plaintext := make([]byte, len(ciphertext))

	// Decrypt the ciphertext
	mode.CryptBlocks(plaintext, ciphertext)

	// Unpad the decrypted plaintext
	unpaddedText := unpad(plaintext)

	return unpaddedText
}

func pad(text []byte, blockSize int) []byte {
	padding := blockSize - (len(text) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

func unpad(text []byte) []byte {
	padding := int(text[len(text)-1])
	return text[:len(text)-padding]
}
