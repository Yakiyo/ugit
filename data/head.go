package data

import (
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/utils"
	"github.com/charmbracelet/log"
)

var (
	_PathHEAD = filepath.Join(GIT_DIR, "HEAD")
)

// set `id` as HEAD
func SetHEAD(id string) error {
	file, err := os.Create(_PathHEAD)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(id))
	log.Info("writing HEAD", "id", id)
	return err
}

// get commit from HEAD
func GetHEAD() (string, error) {
	// If HEAD doesn't exists, it means we don't have any commit
	// for example its a newly created repository
	if !utils.PathExists(_PathHEAD) {
		return "", nil
	}
	content, err := os.ReadFile(_PathHEAD)
	return string(content), err
}