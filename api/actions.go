package api

import "fmt"

type Handler struct {
	Repo Repo
}

func (h *Handler) Create(r *Request) (string, error) {
	if r.Subject == "" {
		return "", fmt.Errorf("Entry Body cannot be blank")
	}
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
	todos, err := h.Repo.Query(meta)

	return todos.String(), err
}
func (h Handler) Edit(r *Request) (string, error) {
	todo, err := h.Repo.Get(r.Id)
	if err != nil {
		return "", err
	}
	if todo == nil {
		return "", fmt.Errorf("Unable to find todo '%s'", r.Id)
	}
	if r.Subject != "" {
		todo.Subject = r.Subject
	}
	if r.Context != "" {
		todo.Meta.Context = r.Context
	}
	todo.Meta.AddTags(r.Tags)
	todo.Meta.RemoveTags(r.TagsToRemove)

	err = h.Repo.Save(todo)
	return todo.String() + "\n", err
}
func (h *Handler) Close(r *Request) (string, error) {
	todo, err := h.Repo.Get(r.Id)
	if err != nil {
		return "", err
	}
	if todo == nil {
		return "", fmt.Errorf("Unable to find todo '%s'", r.Id)
	}
	todo.Status = StatusClosed
	err = h.Repo.Save(todo)
	return todo.String() + "\n", err
}
