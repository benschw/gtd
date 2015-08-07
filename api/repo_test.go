package api

import "strconv"

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

func (r *MemRepo) Get(id string) *Todo {
	return r.todos[id]
}

func (r *MemRepo) Query(meta *Meta) []*Todo {
	todos := make([]*Todo, 0)
	for _, todo := range r.todos {
		if todo.Meta.Context == meta.Context {
			todos = append(todos, todo)
		}
	}
	return todos
}
