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

