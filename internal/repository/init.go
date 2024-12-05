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

// ComputeSHA1 generates a SHA-1 hash for given data
func ComputeSHA1(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// WriteObject stores a git object (blob, tree, commit) in the repository
func (r *Repository) WriteObject(content []byte) (string, error) {
	hash := ComputeSHA1(content)
	objectPath := filepath.Join(r.GitPath, "objects", hash[:2], hash[2:])
	
	// Ensure the subdirectory exists
	err := os.MkdirAll(filepath.Dir(objectPath), 0755)
	if err != nil {
		return "", err
	}

	return hash, os.WriteFile(objectPath, content, 0644)
}

// GetRepositoryRoot finds the root of the current git repository
func GetRepositoryRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		gitDir := filepath.Join(currentDir, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return currentDir, nil
		}

		// Move to parent directory
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// Reached root without finding .git
			return "", fmt.Errorf("not a git repository")
		}
		currentDir = parentDir
	}
}
