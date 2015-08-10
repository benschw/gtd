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
	metaStr := ""
	if t.Meta != nil {
		metaStr = t.Meta.String()
	}
	return fmt.Sprintf(
		"[%s] %s| %s",
		t.Id,
		metaStr,
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

func (m *Meta) AddTags(tags []string) {
	m.Tags = append(m.Tags, tags...)
}
func (m *Meta) RemoveTags(tags []string) {
	n := make([]string, 0)
	for _, tag := range m.Tags {
		if !stringInSlice(tag, tags) {
			n = append(n, tag)
		}
	}
	m.Tags = n
}

func (m *Meta) String() string {
	tagStr := ""
	if m.Tags != nil {
		tagStr = strings.Join(m.Tags, " ") + " "
	}
	return m.Context + " " + tagStr
}
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
