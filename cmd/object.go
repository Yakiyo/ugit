package cmd

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

func HashObj(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid arguments provided, must provide only one argument")
	}
	content, err := os.ReadFile(args[0])
	if err != nil {
		log.Error("failure when reading file", "path", args[0])
		return err
	}
	h, err := createObject(content, "blob")
	if err != nil {
		return err
	}
	fmt.Println(h)
	return nil
}

func CatFile(args []string) error {
	data, err := getObject(args[0], "blob")
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

// content is the file's content, ftype is the file's type (default should be blob)
func createObject(content []byte, ftype string) (string, error) {
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
	_, err = file.Write(data)
	if err != nil {
		log.Error("Error when writing content to object file", "err", err)
	}
	return objname, err
}

// id is the object's id, expected is the expected ftype.
// expected can be an empty string if we do not want to check the type
func getObject(id, expected string) (string, error) {
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
