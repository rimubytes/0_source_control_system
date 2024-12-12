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

// readIndex reads the current staging index
func (r *Repository) readIndex() (*Index, error) {
	indexPath := filepath.Join(r.GitPath, "index")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}

	var index Index
	err = json.Unmarshal(data, &index)
	return &index, err
}

// writeIndex writes the current staging index
func (r *Repository) writeIndex(index *Index) error {
	indexPath := filepath.Join(r.GitPath, "index")
	data, err := json.Marshal(index)
	if err != nil {
		return err
	}
	return os.WriteFile(indexPath, data, 0644)
}

// ClearIndex removes the staging index
func (r *Repository) ClearIndex() error {
	indexPath := filepath.Join(r.GitPath, "index")
	return os.Remove(indexPath)
}