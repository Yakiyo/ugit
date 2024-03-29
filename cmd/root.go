// package command contains functions implementing the
// command line interface
package cmd

// this file contains general commands that don't require a separate file
// or those which are pretty short and their actual implementation lives in
// a separate file/package/dir

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yakiyo/ugit/base"
	"github.com/Yakiyo/ugit/data"
	"github.com/Yakiyo/ugit/utils"
)

func Init() error {
	err := os.MkdirAll(data.GIT_DIR, os.ModePerm)
	if err != nil {
		return err
	}
	return os.MkdirAll(filepath.Join(data.GIT_DIR, "objects"), os.ModePerm)
}

func Commit(args []string) error {
	utils.NArgs(args, 1)
	id, err := base.CreateCommit(args[0])
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func Log(args []string) error {
	var id string
	var err error

	if len(args) > 0 {
		id = args[0]
	} else {
		id, err = data.GetHEAD()
		if err != nil {
			return err
		}
	}

	for id != "" {
		commit, err := base.GetCommit(id)
		if err != nil {
			return err
		}
		commit.Id = id
		fmt.Printf(
			"commit: %v\n"+
				"time: %v\n"+
				"message: %v\n\n",
			commit.Id,
			commit.Time.String(),
			commit.Message,
		)

		id = commit.Parent
	}
	return nil
}
