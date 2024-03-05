package base

import (
	"fmt"
	"strings"
	"time"

	"github.com/Yakiyo/ugit/data"
	"github.com/charmbracelet/log"
)

const timeLayout = "2006-01-02T15:04:05"

type Commit struct {
	// commit id
	Id string
	// the id to the tree object this commit represents
	Tree string
	// commit parent (may be an empty string)
	Parent string
	// time of commit
	Time time.Time
	// the message corresponding to the commit
	Message string
}

func (c *Commit) fromString(body string) error {
	split := strings.SplitN(body, "\n\n", 2)
	if len(split) != 2 {
		return fmt.Errorf("unexpected outcome, splitting of commit object did not return 2 entries, len: %v, body: %v", len(split), body)
	}
	c.Message = split[1]
	for _, line := range strings.Split(strings.TrimSpace(split[0]), "\n") {
		pieces := strings.SplitN(line, " ", 2)
		switch pieces[0] {
		case "id":
			c.Id = pieces[1]
		case "tree":
			c.Tree = pieces[1]
		case "parent":
			c.Parent = pieces[1]
		case "time":
			t, err := time.Parse(timeLayout, pieces[1])
			if err != nil {
				return err
			}
			c.Time = t
		}

	}
	return nil
}

// creates the commit object
//
// a commit object is basically formatted as follows
// tree <oid>
// parent <oid> [this may not be present (such as the initial commit)]
// author <username>
// time <time>
// <empty-line>
// <message>
func CreateCommit(message string) (string, error) {
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
		head = fmt.Sprintf("\nparent %v", head)
	}
	body := fmt.Sprintf(
		"tree %v%v\ntime %v\n\n%v",
		id, head, time.Now().Format(timeLayout),
		message,
	)
	cid, err := data.CreateObject([]byte(body), data.CommitType)
	if err == nil {
		err = data.SetHEAD(cid)
	}
	return cid, err
}

// reads a commit object
func GetCommit(id string) (commit Commit, err error) {
	object, err := data.GetObject(id, data.CommitType)
	if err != nil {
		return
	}
	err = commit.fromString(object)
	return
}
