package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Index represents the staging area
type Index struct {
	Entries map[string]string
}

// AddToStaging stages files for the next commit
func (r *Repository) AddToStaging(paths []string) error {
	index, err := r.readIndex()
	if err != nil {
		index = &Index{Entries: make(map[string]string)}
	}

	for _, path := range paths {
		// Compute relative apth from repo root
		relPath, err := filepath.Rel(r.RootPath, path)
		if err != nil {
			return err
		}

		// Read file content
		content, err := r.ReadFile(path)
		if err != nil {
			return err
		}

		// Update index
		index.Entries[relPath] = blobHash
	}

	return r.writeIndex(index)
}