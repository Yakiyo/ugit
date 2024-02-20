package base

import (
	"fmt"
	"time"

	"github.com/Yakiyo/ugit/data"
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
	body := fmt.Sprintf(
		"tree %v\ntime %v\n\n%v",
		id, time.Now().Format("2006-01-02T15:04:05"),
		message,
	)
	return data.CreateObject([]byte(body), data.CommitType)
}
