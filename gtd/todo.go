package gtd

import (
	"fmt"
	"strings"
)

const (
	StatusOpen   = "open"
	StatusClosed = "closed"
)

type Todo struct {
	Id      string
	Meta    *Meta
	Subject string
	Body    string
	Status  string
}

func (t *Todo) String() string {
	return fmt.Sprintf(
		"[%s] %s| %s",
		t.Id,
		t.Meta.String(),
		t.Subject,
	)
}

type Meta struct {
	Context string
	Tags    []string
}

func (m *Meta) String() string {
	return fmt.Sprintf(
		"%s %s ",
		m.Context,
		strings.Join(m.Tags, " "),
	)
}
