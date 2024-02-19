package base

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

// WriteTree - this is the directory equivalent of `data.CreateObject`
func WriteTree(dir string) error {
	// dirs to skip
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			// we skip directories like `.git` and `.ugit`
			if !shouldSkip(entry.Name()) {
				WriteTree(filepath.Join(dir, entry.Name()))
			}
			continue
		}
		path := filepath.Join(dir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		objname, err := data.CreateObject(content, "blob")
		if err != nil {
			return err
		}
		log.Info(objname, "path", path)
	}
	return nil
}

func shouldSkip(path string) bool {
	skips := []string{".git", data.GIT_DIR}
	for _, skip := range skips {
		if strings.Contains(path, skip) {
			return true
		}
	}
	return false
}
