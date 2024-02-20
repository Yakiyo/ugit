// commands related to creating & reading
// objects for single files

package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/ugit/data"
	"github.com/Yakiyo/ugit/utils"
	"github.com/charmbracelet/log"
)

// hash a file and create object with the hash as file name
func HashObject(args []string) error {
	utils.CwdIsRepo()
	utils.NArgs(args, 1)
	content, err := os.ReadFile(args[0])
	if err != nil {
		log.Error("failure when reading file", "path", args[0])
		return err
	}
	h, err := data.CreateObject(content, data.BlobType)
	if err != nil {
		return err
	}
	fmt.Println(h)
	return nil
}

// read a object with the filename as the given hash
func CatFile(args []string) error {
	utils.CwdIsRepo()
	utils.NArgs(args, 1)
	data, err := data.GetObject(args[0], "")
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
