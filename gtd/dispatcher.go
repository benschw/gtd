package gtd

const (
	ContextPrefix = "@"
	TagPrefix     = "#"
	ActionNew     = "+"
	ActionDone    = "-"
	ActionList    = "l"
	ActionSave    = "m"
)

type Action func(*Request) (string, error)

func Dispatch(req *Request, repo Repo) (string, error) {
	h := &Handler{Repo: repo}

	handlers := map[string]Action{
		ActionNew:  h.Create,
		ActionDone: h.Close,
		ActionList: h.List,
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
	return todo.Id, err
}
func (h Handler) List(r *Request) (string, error) {
	meta := &Meta{Context: r.Context, Tags: r.Tags}
	todos := h.Repo.Query(meta)

	var c TodoCollection = todos
	return c.String(), nil
}
func (h *Handler) Close(r *Request) (string, error) {
	todo := h.Repo.Get(r.Id)
	todo.Status = StatusClosed
	h.Repo.Save(todo)
	return todo.Id, nil
}
