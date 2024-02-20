package data

import (
	"os"
	"path/filepath"
)

// set `id` as HEAD
func SetHEAD(id string) error {
	file, err := os.Create(filepath.Join(GIT_DIR, "HEAD"))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(id))
	return err
}