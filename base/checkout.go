package base

import "github.com/Yakiyo/ugit/data"

func Checkout(id string) error {
	c, err := GetCommit(id)
	if err != nil {
		return err
	}
	err = ReadTree(c.Tree, ".")
	if err != nil {
		return err
	}
	return data.SetHEAD(c.Id)
}
