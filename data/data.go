package data

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

const GIT_DIR = ".ugit"

const (
	BlobType = "blob"
	TreeType = "tree"
)

// content is the file's content, ftype is the file's type (default should be blob)
func CreateObject(content []byte, ftype string) (string, error) {
	data := []byte{}
	data = append(data, []byte(ftype)...)
	data = append(data, []byte("\x00")...)
	data = append(data, content...)

	sum := sha1.Sum(data)
	objname := hex.EncodeToString(sum[:])
	path := filepath.Join(GIT_DIR, "objects", string(objname[:]))
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir); err != nil {
		return "", err
	}

	file, err := os.Create(path)
	if err != nil {
		log.Error("Error creating object file", "err", err)
		return "", err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Error("Error when writing content to object file", "err", err)
	}
	return objname, err
}

// id is the object's id, expected is the expected ftype.
// expected can be an empty string if we do not want to check the type
func GetObject(id, expected string) (string, error) {
	content, err := os.ReadFile(filepath.Join(GIT_DIR, "objects", id))
	if err != nil {
		return "", err
	}
	split := strings.SplitN(string(content), "\x00", 2)
	if len(split) != 2 {
		return "", fmt.Errorf("unexpected result, should have received exactly 2 substrings, got %#v", split)
	}
	ftype := split[0]
	if expected != "" && ftype != expected {
		return "", fmt.Errorf("expected != ftype, expected = %v, ftype = %v", expected, ftype)
	}
	return split[1], nil
}
