package gtd

type Request struct {
	Action       string
	Id           string
	Context      string
	Tags         []string
	TagsToRemove []string
	Subject      string
}
