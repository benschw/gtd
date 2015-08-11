package api

type Repo interface {
	Save(*Todo) error
	Get(string) (*Todo, error)
	Query(*Meta) (TodoCollection, error)
}
