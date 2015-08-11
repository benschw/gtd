package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benschw/gtd/api"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// Get Config from ENV variables
	cfg, err := api.NewDefaultConfig()
	if err != nil {
		api.EnvHelp()
		os.Exit(1)
	}

	// Parse cli args into Request
	r, err := api.ParseArgs(args, cfg.Context)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	// Dispatch Request
	out, err := api.Dispatch(r, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(3)
	}

	// Display output from processing request
	fmt.Print(out)
}
