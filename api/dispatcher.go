package api

import "fmt"

type Action func(*Request) (string, error)

func Dispatch(req *Request, repo Repo) (string, error) {
	h := &Handler{Repo: repo}

	handlers := map[string]Action{
		ActionNew:   h.Create,
		ActionClose: h.Close,
		ActionEdit:  h.Edit,
		ActionList:  h.List,
	}

	return handlers[req.Action](req)
}

type Handler struct {
	Repo Repo
}

func (h *Handler) Create(r *Request) (string, error) {
	todo := &Todo{
		Meta: &Meta{
			Context: r.Context,
			Tags:    r.Tags,
		},
		Subject: r.Subject,
		Status:  StatusOpen,
	}
	err := h.Repo.Save(todo)
	return todo.String() + "\n", err
}
func (h Handler) List(r *Request) (string, error) {
	meta := &Meta{Context: r.Context, Tags: r.Tags}
	todos := h.Repo.Query(meta)

	var c TodoCollection = todos
	return c.String(), nil
}
func (h Handler) Edit(r *Request) (string, error) {
	todo := h.Repo.Get(r.Id)
	if todo == nil {
		return "", fmt.Errorf("Unable to find todo '%s'", r.Id)
	}
	if r.Subject != "" {
		todo.Subject = r.Subject
	}
	if r.Context != "" {
		todo.Meta.Context = r.Context
	}

	tags := append(todo.Meta.Tags, r.Tags...)
	n := make([]string, 0)
	for _, tag := range tags {
		if !stringInSlice(tag, r.TagsToRemove) {
			n = append(n, tag)
		}
	}
	todo.Meta.Tags = n
	err := h.Repo.Save(todo)
	return todo.String() + "\n", err
}
func (h *Handler) Close(r *Request) (string, error) {
	todo := h.Repo.Get(r.Id)
	if todo == nil {
		return "", fmt.Errorf("Unable to find todo '%s'", r.Id)
	}
	todo.Status = StatusClosed
	err := h.Repo.Save(todo)
	return todo.String() + "\n", err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
