package gtd

import (
	"strconv"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Repo interface {
	Save(*Todo) error
	Get(string) *Todo
	Query(*Meta) []*Todo
}

func NewGhRepo(user string, repo string, token string) *GhRepo {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	return &GhRepo{
		Client: client,
		Owner:  user,
		Repo:   repo,
	}
}

type GhRepo struct {
	Client *github.Client
	Owner  string
	Repo   string
}

func (r *GhRepo) Save(todo *Todo) error {
	if todo.Id == "" {
		labels := append(todo.Meta.Tags, todo.Meta.Context)
		issue, _, err := r.Client.Issues.Create(r.Owner, r.Repo, &github.IssueRequest{
			Title:  &todo.Subject,
			Labels: &labels,
		})

		if err != nil {
			return err
		}
		todo.Id = strconv.Itoa(*issue.Number)
	}
	return nil
}

func (r *GhRepo) Get(id string) *Todo {
	return nil
}

func (r *GhRepo) Query(meta *Meta) []*Todo {
	todos := make([]*Todo, 0)

	labels := append(meta.Tags, meta.Context)
	issues, _, err := r.Client.Issues.ListByRepo(r.Owner, r.Repo, &github.IssueListByRepoOptions{
		Labels: labels,
	})
	if err != nil {
		return todos
	}

	for _, issue := range issues {
		todos = append(todos, parseTodoFromIssue(issue))
	}

	return todos
}
func parseTodoFromIssue(issue github.Issue) *Todo {
	meta := parseMetaFromIssue(issue)

	todo := &Todo{
		Id:      strconv.Itoa(*issue.Number),
		Subject: *issue.Title,
		Meta:    meta,
	}
	return todo
}
func parseMetaFromIssue(issue github.Issue) *Meta {
	meta := &Meta{}
	for _, label := range issue.Labels {
		name := *label.Name
		if name[0:1] == ContextPrefix {
			meta.Context = name
		}
		if name[0:1] == TagPrefix {
			meta.Tags = append(meta.Tags, name)
		}
	}
	return meta
}
