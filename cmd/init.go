package cmd

import (
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/data"
)

func Init() error {
	err := os.MkdirAll(data.GIT_DIR, os.ModePerm)
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(data.GIT_DIR, "objects"), os.ModePerm)
}
