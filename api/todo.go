package api

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

type TodoCollection []*Todo

func (t TodoCollection) String() string {
	var out string
	for _, todo := range t {
		out = fmt.Sprintf("%s%s\n", out, todo)
	}
	return out
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
