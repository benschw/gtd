

	
## Usage

	gtd action [id] [meta[,meta[,...]]] [subject]

### Actions

- + new todo

	gtd + @work \#foo Hello World
	1

- - mark todo as done

	gtd - 1
	1

- l list todos (filtered by _meta_)

	gtd l @work
	# list of all @work todos

- m modify todo

	gtd m 1 @home -\#foo #bar
	# set context to @home, remove #foo tag, add #bar tag

	gtd m 1 @work Hello Galaxy
	# set context to @work, update subject

### meta

- @CONTEXT prefix with @ to set context (there can only be one context, if many are supplied, only one will be used)
- @TAG prefix with # to assign a tag to the todo. many can be used
	- when modifying a todo, -#tag will remove a tag if it exists

