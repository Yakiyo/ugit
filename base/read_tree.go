package base

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

func ReadTree(treeid, cwd string) error {
	// TODO: clear current dir first
	tree, err := getTree(treeid, cwd)
	if err != nil {
		return err
	}
	log.Debug("finished reading tree", "tree", tree)
	for path, id := range tree {
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}
		content, err := data.GetObject(id, "")
		if err != nil {
			return err
		}
		err = os.WriteFile(path, []byte(content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

// read a tree object
func getTree(id, base string) (map[string]string, error) {
	tree := map[string]string{}
	entries, err := iterTreeItems(id)
	if err != nil {
		return tree, err
	}
	for _, i := range entries {
		if strings.Contains(i.name, "/") || i.name == ".." || i.name == "." {
			return tree, fmt.Errorf("invalid value for name found in entry, name = %v", i.name)
		}
		path := filepath.Join(base, i.name)

		if i.ftype == data.BlobType {
			tree[path] = i.id
		} else if i.ftype == data.TreeType {
			innertree, err := getTree(i.id, path)
			if err != nil {
				return tree, err
			}
			for k, v := range innertree {
				tree[k] = v
			}
		} else {
			log.Fatal("unknown ftype received", "ftype", i.ftype)
		}
	}
	return tree, nil
}

// create a slice from a tree's entries
func iterTreeItems(id string) ([]objItem, error) {
	items := []objItem{}
	if id == "" {
		return items, nil
	}
	tree, err := data.GetObject(id, data.TreeType)
	if err != nil {
		return items, err
	}

	for _, line := range strings.Split(strings.TrimSpace(tree), "\n") {
		s := strings.Split(line, " ")
		if len(s) != 3 {
			return items, fmt.Errorf("unexpected error when parsing line entry in tree object, got entry lenght %v, line %v", len(s), line)
		}
		name, id, ftype := s[0], s[1], s[2]
		items = append(items, objItem{name, id, ftype})
	}
	return items, nil
}

/** copy pasta from cmd/util.go **/

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
