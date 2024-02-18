package main

import (
	"os"
	"path/filepath"
)

const GIT_DIR = ".ugit"

func initCmd() error {
	err := os.MkdirAll(GIT_DIR, os.ModePerm)
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(GIT_DIR, "objects"), os.ModePerm)
}
