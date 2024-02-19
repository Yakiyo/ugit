package base

func ReadTree(treeid, cwd string) error {
	_, err := getTree(treeid, cwd)
	if err != nil {
		return err
	}
	return nil
}

func getTree(id, base string) (map[string]string, error) {
	dict := map[string]string{}
	_, err := iterTreeItems(id)
	if err != nil {
		return dict, err
	}
	return dict, nil
}

// create a slice from a tree's entries
func iterTreeItems(id string) ([]objItem, error) {
	items := []objItem{}

	return items, nil
}
