package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benschw/gtd/api"
	"github.com/benschw/gtd/ghissues"
)

func newRepo() (api.Repo, error) {
	// Build Github Issues Repo (configured with ENV)
	switch os.Getenv("GTD_REPO") {
	case "ghissues":
		return ghissues.New()
	}

	return nil, fmt.Errorf("Backend Repo not set or unknown")
}

func main() {
	flag.Parse()
	args := flag.Args()

	// Parse cli args into Request
	defaultCtx := os.Getenv("GTD_CONTEXT")
	if defaultCtx == "" {
		defaultCtx = "@global"
	}

	r, err := api.ParseArgs(args, defaultCtx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	// Build Backend Repo to store todos to
	repo, err := newRepo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	// Dispatch Request
	out, err := api.Dispatch(r, repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(3)
	}

	// Display output from processing request
	fmt.Print(out)
}
