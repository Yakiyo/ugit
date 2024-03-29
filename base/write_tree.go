package base

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/data"
	"github.com/Yakiyo/ugit/utils"
	"github.com/charmbracelet/log"
)

type objItem struct {
	name, id, ftype string
}

// WriteTree - this is the directory equivalent of `data.CreateObject`
func WriteTree(dir string) (string, error) {
	objects := []objItem{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			// we skip directories like `.git` and `.ugit`
			if !utils.ShouldSkip(entry.Name()) {
				id, err := WriteTree(filepath.Join(dir, entry.Name()))
				if err != nil {
					return "", err
				}
				objects = append(objects, objItem{entry.Name(), id, "tree"})
			}
			continue
		}
		path := filepath.Join(dir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		id, err := data.CreateObject(content, data.BlobType)
		if err != nil {
			return "", err
		}
		objects = append(objects, objItem{entry.Name(), id, data.BlobType})
		log.Debug(id, "path", path)
	}
	var str string
	for _, i := range objects {
		str += fmt.Sprintf("%v %v %v\n", i.name, i.id, i.ftype)
	}
	return data.CreateObject([]byte(str), data.TreeType)
}
