package api

import "strings"

const (
	ContextPrefix = "@"
	TagPrefix     = "#"
	ActionNew     = "a"
	ActionClose   = "c"
	ActionEdit    = "m"
	ActionList    = "l"
)

func ParseArgs(args []string, defaultCtx string) (*Request, error) {
	r := &Request{}

	r.Action, args = extractAction(args)
	if r.Action == ActionNew {
		r.Context = defaultCtx
	}
	if r.Action == ActionEdit || r.Action == ActionClose {
		r.Id = args[0]
		args = args[1:]
	}

	context, tags, toRem, args := extractMeta(args)
	if context != "" {
		r.Context = context
	}
	r.Tags = tags
	r.TagsToRemove = toRem
	r.Subject = strings.Join(args, " ")

	return r, nil
}

func extractAction(args []string) (string, []string) {
	if len(args) > 0 {
		switch args[0] {
		case ActionNew:
			return ActionNew, args[1:]
		case ActionClose:
			return ActionClose, args[1:]
		case ActionEdit:
			return ActionEdit, args[1:]
		case ActionList:
			return ActionList, args[1:]
		}
	}
	return ActionList, args
}

func extractMeta(args []string) (string, []string, []string, []string) {
	var context string
	tags := make([]string, 0)
	toRem := make([]string, 0)

	metaComplete := false
	rem := make([]string, 0)
	for i := 0; i < len(args); i++ {
		if !metaComplete {
			if strings.HasPrefix(args[i], ContextPrefix) {
				context = args[i]
			} else if strings.HasPrefix(args[i], TagPrefix) {
				tags = append(tags, args[i])
			} else if strings.HasPrefix(args[i], "-"+TagPrefix) {
				toRem = append(toRem, args[i])
			} else {
				metaComplete = true
			}
		}
		if metaComplete {
			rem = append(rem, args[i])
		}
	}

	return context, tags, toRem, rem
}
