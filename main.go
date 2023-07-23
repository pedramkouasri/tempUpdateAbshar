package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/plumbing"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {
	repoPath := "<path/to/repository>"
	branch1Name := "branch1"
	branch2Name := "branch2"

	// Open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	

	// Get the references for the branches
	branch1Ref, err := repo.Reference(x, true)
	if err != nil {
		log.Fatal(err)
	}

	branch2Ref, err := repo.Reference(plumbing.ReferenceName("refs/heads/" + branch2Name), true)
	if err != nil {
		log.Fatal(err)
	}

	// Get the commit objects for the branch references
	branch1Commit, err := repo.CommitObject(branch1Ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	branch2Commit, err := repo.CommitObject(branch2Ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// Get the diff between the commits
	diff, err := branch1Commit.Diff(branch2Commit)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the file changes in the diff
	err = diff.ForEach(func(fileDiff *object.Change) error {
		fmt.Printf("File: %s\n", fileDiff.To.Name)

		patch, err := fileDiff.Patch()
		if err != nil {
			return err
		}

		fmt.Printf("Patch:\n%s\n", patch.String())

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
