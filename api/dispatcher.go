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
	return todo.Id + "\n", err
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
	for _, tag := range r.Tags {
		todo.Meta.Tags = append(todo.Meta.Tags, tag)
	}
	n := make([]string, 0)
	for _, tag := range todo.Meta.Tags {
		found := false
		for _, rem := range r.TagsToRemove {
			if tag == rem {
				found = true
			}
		}
		if !found {
			n = append(n, tag)
		}
	}
	todo.Meta.Tags = n
	err := h.Repo.Save(todo)
	return todo.Id, err
}
func (h *Handler) Close(r *Request) (string, error) {
	todo := h.Repo.Get(r.Id)
	if todo == nil {
		return "", fmt.Errorf("Unable to find todo '%s'", r.Id)
	}
	todo.Status = StatusClosed
	h.Repo.Save(todo)
	return todo.Id, nil
}
