package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pedramkousari/update-abshar/helpers"
)

func main() {
	files := []string{
		"/home/pedram/Projects/Go/UpdateAbshar/tmp/111.pdf",
		"/home/pedram/Projects/Go/UpdateAbshar/tmp/package.json",
	}

	outputFile := "archive.tar.gz"

	err := createTarGz(files, outputFile)
	if err != nil {
		log.Fatal(err)
	}

	key := "e10adc3949ba59abbe56e057f20f883e"

	data,err := ioutil.ReadFile(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	helpers.Encrypt(data,[]byte(key))
	

	log.Printf("Successfully created %s", outputFile)
}


// Function to encrypt a file using OpenSSL
func encryptFileWithOpenSSL(filePath string) error {
	password := "e10adc3949ba59abbe56e057f20f883e"// Generate a random password

	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-in", filePath, "-out", filePath+".enc", "-k", password)
	err := cmd.Run()

	if err != nil {
		return err
	}

	// Remove the original unencrypted tar.gz file
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func encryptFile(key []byte, inputFile string, outputFile string) error {
	plaintext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ioutil.WriteFile(outputFile, ciphertext, 0644)
}

func decryptFile(key []byte, inputFile string, outputFile string) error {
	ciphertext, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(ciphertext) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ioutil.WriteFile(outputFile, ciphertext, 0644)
}

func createTarGz(files []string, outputFile string) error {
	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}


	// Create a gzip writer
	gw := gzip.NewWriter(outFile)	
	defer gw.Close()

	// Create a tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Iterate over the input files
	for _, file := range files {
		err = addFileToTar(file, tw)
		if err != nil {
			return err
		}
	}

	return nil
}

func addFileToTar(file string, tw *tar.Writer) (error) {
	// Open the input file
	inFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Get the file information
	info, err := inFile.Stat()
	if err != nil {
		return err
	}

	// Create a tar header based on the file info
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	// Set the name of the file within the tar archive
	header.Name = filepath.Base(file)

	// Write the header the tar writer
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Copy the file content to the tar writer
	_, err = io.Copy(tw, inFile)
	if err != nil {
		return err
	}

	return nil
}