package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

var ErrNotGItDir error = errors.New("cwd is not a working ugit directory, create one with `ugit init`")

// assert if min `n` arguments where provided
func nArgs(args []string, n int) {
	if len(args) != n {
		log.Errorf("argument error, require %v arguments to be provided", n)
		os.Exit(1)
	}
}

func cwdIsRepo(args ...string) {
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

// exists returns whether the given file or directory exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}