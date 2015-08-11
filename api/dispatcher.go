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

	handler, ok := handlers[req.Action]
	if !ok {
		return "", fmt.Errorf("action not supported")
	}

	return handler(req)
}
