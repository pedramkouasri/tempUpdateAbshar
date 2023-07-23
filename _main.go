package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pedramkousari/update-abshar/helpers"
)

const (
	// tempDirectory string = "./tmp"
	// packagePathFile string = "./tmp/package.json"
	messagePath string = "Path Of Tar File:"
)


func main() {
	logFile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	err = godotenv.Load(".env")
	handleError(err)

	file, err := os.Create("ooo.zip");
	handleError(err)

	if err:= helpers.Tar("/home/pedram/Downloads/111.pdf", file); err != nil {
		panic(err)
	}

	// services.CreatePatch()

	// simpleEncrypt()


	// path,err := getPathFile()
	// handleError(err)

	// patchFile, err := os.Open(path)
	// handleError(err)

	// err = helpers.Untar("./tmp", patchFile)
	// handleError(err)

	// // baadbaanDir := os.Getenv("BAADBAAN_DIRECTORY")
	// // _, err = logFile.Write([]byte(fmt.Sprintf("%s\n", baadbaanDir)))

	// _,err = os.Stat(packagePathFile)
	// handleError(err)


	// file, err := os.Open(packagePathFile)
	// handleError(err)

	// pkg := []types.Packages{}

	// decoder := json.NewDecoder(file)
	// err = decoder.Decode(&pkg)
	// handleError(err)

	// latsVersion := pkg[len(pkg)-1]
	// fmt.Printf("%+v \n", latsVersion)
}

func getPathFile() (string,error) {
	fmt.Print(messagePath)	
	var path string
	if _, err := fmt.Scan(&path); err!=nil {
		return "", err
	}

	return path, nil
}

func handleError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func simpleEncrypt(){
	text := "Hello, World!"
	key := "e10adc3949ba59abbe56e057f20f883e"
	// password := "mysecretpassword"

	cmd := exec.Command("/usr/bin/openssl ","enc", "-aes-256-cbc", "-a", "-salt", "-k", key)
	cmd.Stdin = strings.NewReader(text)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error:", err)
		return
	}

	encryptedText := string(output)
	fmt.Println("Encrypted Text:", encryptedText)
}