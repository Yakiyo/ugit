package cmd

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"

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
	h, err := createHash(content, "blob")
	if err != nil {
		return err
	}
	fmt.Println(h)
	return nil
}

func createHash(content []byte, ftype string) ([]byte, error) {
	data := []byte{}
	data = append(data, []byte(ftype)...)
	data = append(data, []byte("\x00")...)
	data = append(data,  content...)

	objname := sha1.Sum(data)
	path := filepath.Join(GIT_DIR, "objects", string(objname[:]))
	if err := os.MkdirAll(filepath.Dir(path), os.ModeDir); err != nil {
		return []byte{}, err
	}
	file, err := os.Create(path)
	if err != nil {
		log.Error("Error creating object file", "err", err)
		return []byte{}, err
	}
	_, err = file.Write(data)
	if err != nil {
		log.Error("Error when writing content to object file", "err", err)
	}
	return objname[:], err
}