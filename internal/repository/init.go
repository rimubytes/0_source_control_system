package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

// Repository represents the structure of a git-like repository
type Repository struct {
	RootPath string
	GitPath string
	CurrentBranch string
}

// Init initialize a new repository in the specified path
func Init(path string) (*Repository, error) {
	// Ensure the path exists and is absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	gitPath := filepath.Join(absPath, ".git")
	err = os.MkdirAll(filepath.Join(gitPath, "objects"), 0755)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(filepath.Join(gitPath, "refs", "heads"), 0755)
	if err != nil {
		return nil, err
	}

	// Create initial HEAD file pointing to master branch
	headPath := filepath.Join(gitPath, "HEAD")
	err = os.WriteFile(headPath, []byte("ref: refs/heads/master"), 0644)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		RootPath:     absPath,
		GitPath:      gitPath,
		CurrentBranch: "master",
	}

	return repo, nil
}