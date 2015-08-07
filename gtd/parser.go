package gtd

import "strings"

func ParseArgs(args []string, defaultCtx string) (*Request, error) {
	r := &Request{}

	r.Action, args := extractAction(args)
	if r.Action == ActionNew {
		r.Context = defaultCtx
	}

	context, tags, args := extractMeta(args)
	if context != "" {
		r.Context = context
	}
	r.Tags = tags

	r.Subject = strings.Join(args, " ")

	return r, nil
}

func extractAction(args []string) (string, []string) {
	if len(args) > 0 {
		switch args[0] {
		case ActionNew:
			return ActionNew, args[1:]
		case ActionDone:
			return ActionDone, args[1:]
		case ActionList:
			return ActionList, args[1:]
		}
	}
	return ActionList, args
}

func extractMeta(args []string) (string, []string, []string) {
	var context string
	tags := make([]string, 0)

	metaComplete := false
	rem := make([]string, 0)
	for i := 0; i < len(args); i++ {
		if !metaComplete {
			if strings.HasPrefix(args[i], ContextPrefix) {
				context = args[i]
			} else if strings.HasPrefix(args[i], TagPrefix) {
				tags = append(tags, args[i])
			} else {
				metaComplete = true
			}
		}
		if metaComplete {
			rem = append(rem, args[i])
		}
	}

	return context, tags, rem
}
