package gtd

import (
	"fmt"
	"strings"
)

const (
	ContextPrefix = "@"
	TagPrefix     = "#"
	ActionNew     = "+"
	ActionDone    = "-"
	ActionList    = "list"
)

type Dispatcher struct {
	Repo Repo
}

func (d *Dispatcher) Dispatch(meta *Meta, args []string) (string, error) {
	action, args := extractAction(args)

	context, tags, args := extractMeta(args)
	if context != "" {
		meta.Context = context
	}
	meta.Tags = tags

	subj := strings.Join(args, " ")

	fmt.Printf("%+v\n%s\n", meta, subj)
	todo := &Todo{Meta: meta, Subject: subj}

	var out string
	var err error
	switch action {
	case ActionNew:
		out, err = d.Create(todo)
	case ActionDone:
		out, err = d.Close(todo)
	case ActionList:
		out, err = d.List(meta)
	}

	return out, err
}

func (d *Dispatcher) Create(todo *Todo) (string, error) {
	todo.Status = StatusOpen
	err := d.Repo.Save(todo)
	return todo.Id, err
}
func (d *Dispatcher) List(meta *Meta) (string, error) {
	todos := d.Repo.Query(meta)
	var out string
	for _, todo := range todos {
		out = fmt.Sprintf("%s%s\n", out, todo)
	}
	return out, nil
}
func (d *Dispatcher) Close(todo *Todo) (string, error) {
	return "", nil
}

func extractAction(args []string) (string, []string) {
	switch args[0] {
	case ActionNew:
		return ActionNew, args[1:]
	case ActionDone:
		return ActionDone, args[1:]
	case ActionList:
		return ActionList, args[1:]
	}
	return ActionNew, args
}

func extractMeta(args []string) (string, []string, []string) {
	var context string
	tags := make([]string, 0)

	metaComplete := false
	rem := make([]string, 0)
	for i := 0; i < len(args); i++ {
		if !metaComplete {
			if strings.HasPrefix(args[i], ContextPrefix) {
				context = args[i]
			} else if strings.HasPrefix(args[i], TagPrefix) {
				tags = append(tags, args[i])
			} else {
				metaComplete = true
			}
		}
		if metaComplete {
			rem = append(rem, args[i])
		}
	}

	return context, tags, rem
}
