package base

import (
	"fmt"
	"time"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

// creates the commit object
//
// a commit object is basically formatted as follows
// tree <oid>
// author <username>
// time <time>
// <empty-line>
// <message>
func Commit(message string) (string, error) {
	id, err := WriteTree(".")
	if err != nil {
		return "", err
	}
	head, err := data.GetHEAD()
	if err != nil {
		log.Warn("Unexpected error when trying to read head", "err", err)
		head = ""
	}
	if head != "" {
		head = fmt.Sprintf("parent %v\n", head)
	}
	body := fmt.Sprintf(
		"tree %v%v\ntime %v\n\n%v",
		id, head, time.Now().Format("2006-01-02T15:04:05"),
		message,
	)
	cid, err := data.CreateObject([]byte(body), data.CommitType)
	if err == nil {
		err = data.SetHEAD(cid)
	}
	return cid, err
}
