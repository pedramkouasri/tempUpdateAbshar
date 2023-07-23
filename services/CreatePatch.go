package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pedramkousari/update-abshar/types"
)

const (
	technicalRris string = "git@10.10.10.236:root/technical-risk-micro-service.git"
	baadbaan      string = "git@10.10.10.26:root/baadbaan_new.git"
)

func CreatePatch() {
	packages := []types.Packages{}

	getDiffComposer("22.1.2", "22.2")
	createTarFile()
	gzipTarFile()
	moveFile()
	encryptFile()
	destroyFiles()

	reader := strings.NewReader(`[
		{
			"version": "10",
			"package_version": {
			"baadbaan": "v11.10.0",
			"technical": "v1.1.0"
			}
		},
		{
			"version": "11",
			"package_version": {
			"baadbaan": "v12.10.0"
			}
		},
		{
			"version": "12",
			"package_version": {
			"baadbaan": "v12.13.0"
			}
		}
	]
`)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&packages); err != nil {
		log.Fatalf("Error Is %v", err)
	}

	fmt.Println(packages)

}

func getDiffComposer(current_tag string, lasts_tag string) {
	// git diff --name-only --diff-filter=ACMR {lastTag} {current_tag} > diff.txt'
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git --git-dir %s/.git  diff --name-only --diff-filter ACMR %s %s > %s/diff.txt", os.Getenv("BAADBAAN_DIRECTORY"), lasts_tag, current_tag, os.Getenv("BAADBAAN_DIRECTORY")))

	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func createTarFile() {
	// tar -cf patch.tar --files-from=diff.txt
	command := fmt.Sprintf("cd %s &&  tar -cf patch.tar --files-from=diff.txt", os.Getenv("BAADBAAN_DIRECTORY"))
	cmd := exec.Command("sh", "-c", command)

	_, err := cmd.Output()
	if err != nil && err.Error() != "exit status 2" {
		log.Fatal(err)
	}
}

func gzipTarFile() {
	// cd {baadbaan_path} && gzip -f patch.tar
	cmd := exec.Command("gzip", "-f", fmt.Sprintf("%s/patch.tar", os.Getenv("BAADBAAN_DIRECTORY")))

	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func addPackageToTarFile() {}
func moveFile()            {}
func encryptFile()         {}
func destroyFiles()        {}
