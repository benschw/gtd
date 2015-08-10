package api

import (
	"fmt"
	"strconv"
)

func NewMemRepo() *MemRepo {
	return &MemRepo{
		todos: make(map[string]*Todo, 0),
	}
}

type MemRepo struct {
	todos map[string]*Todo
}

func (r *MemRepo) Save(todo *Todo) error {
	if todo.Id == "" {
		todo.Id = strconv.Itoa(len(r.todos))
	}
	r.todos[todo.Id] = todo
	return nil
}

func (r *MemRepo) Get(id string) (*Todo, error) {
	if _, ok := r.todos[id]; !ok {
		return nil, fmt.Errorf("id '%s' not found", id)
	}
	return r.todos[id], nil
}

func (r *MemRepo) Query(meta *Meta) (TodoCollection, error) {
	todos := make([]*Todo, 0)
	for _, todo := range r.todos {
		if todo.Meta.Context == meta.Context {
			todos = append(todos, todo)
		}
	}
	var c TodoCollection = todos
	return c, nil
}
