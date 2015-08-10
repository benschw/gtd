package api

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
