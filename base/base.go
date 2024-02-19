package base

import (
	"io/fs"
	"path/filepath"
)

// WriteTree - this is the directory equivalent of `data.CreateObject`
func WriteTree(dir string) error {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}