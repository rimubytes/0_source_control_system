package repository

import (
	"fmt"
	"path/filepath"
	"github.com/rimubytes/0_source_control_system/pkg/objects"
)

// Commit creates a new commit in the repository
func (r *Repository) Commit(message string) (string, error) {
	// Read current index
	index, err := r.readIndex()
	if err != nil {
		return "", err
	}

	// Create tree object from index
	treeHash, err := r.createTreefromIndex(index)
	if err != nil {
		return "", err
	}

	// Get parent commit hash
	parentHash, err := r.getCurrentBranchHead()
	var parents []string
	if err == nil {
		parents = []string{parentHash}
	}

	// Create commit object
	commit := objects.NewCommit(
		treeHash,
		parents,
		"User <user@example.com>",
		message,
	)

	// Serialize commit
	commitData, err := commit.Serialize()
	if err != nil {
		return "", err
	}

	// Update branch reference
	err = r.updateBranchRef(r.CurrentBranch, commitHash)
	if err != nil {
		return "", err 
	}

	// Clear index after commit
	return commitHash, r.ClearIndex()
}

// getCurrentBranchHead retrieves the current branch's head commit hash
func (r *Repository) getCurrentBranchHead() (string, error) {
	branchPath := filepath.Join(r.GitPath, "refs", "heads", r.CurrentBranch)
	headContent, err := os.ReadFile(branchPath)
	if err != nil {
		return "", err
	}

	return string(headContent), nil
}

// updateBranchRef updates the reference for a specific branch
func (r *Repository) updateBranchRef(branchName, commitHash string) error {
	branchPath := filepath.Join(r.GitPath, "refs", "heads", branchName)
	return os.WriteFile(branchPath, []byte(commitHash), 0644)
}

// createTreeFromIndex creates a tree object from the current index
func (r *Repository) createTreeFromIndex(index *Index) (string, error) {
	// This is a simplified implementation
	// In a real git implementation, this would create a hierarchical tree object
	treeEntries := make([]interface{}, 0, len(index.Entries))
	
	for path, hash := range index.Entries {
		treeEntries = append(treeEntries, map[string]string{
			"path": path,
			"hash": hash,
			"type": "blob",
		})
	}

	treeData, err := json.Marshal(treeEntries)
	if err != nil {
		return "", err
	}

	return r.WriteObject(treeData)
}