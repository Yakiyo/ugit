package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

var ErrNotGItDir error = errors.New("cwd is not a working ugit directory, create one with `ugit init`")

// exists returns whether the given file or directory exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func CwdIsRepo(args ...string) {
	var cwd string
	if len(args) == 0 {
		wd, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("Failed to get cwd, err = %v", err))
		}
		cwd = wd
	} else {
		cwd = args[0]
	}
	if !PathExists(filepath.Join(cwd, data.GIT_DIR)) {
		log.Error(ErrNotGItDir)
		os.Exit(1)
	}
}

// assert if `n` arguments where provided
func NArgs(args []string, n int) {
	if len(args) != n {
		log.Errorf("argument error, require %v arguments to be provided", n)
		os.Exit(1)
	}
}

func ShouldSkip(path string) bool {
	skips := []string{".git", data.GIT_DIR}
	for _, skip := range skips {
		if strings.Contains(path, skip) {
			return true
		}
	}
	return false
}

// recursively scan `dir` and return the file names (full path relative to `dir`)
func ScanDir(dir string) ([]string, error) {
	files := []string{}
	err := scan(dir, &files)
	return files, err
}

// scan populates the `files` slice with path.
// we use a separate func from `ScanDir` so that we can initialize
// a single slice and just keep appending to it later on
func scan(dir string, files *[]string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var path string
	for _, entry := range entries {
		path = filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			err = scan(path, files)
			if err != nil {
				return err
			}
		} else {
			*files = append(*files, path)
		}
	}
	return nil
}
