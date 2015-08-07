package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/benschw/ghgtd/gtd"
)

func envHelp() {
	fmt.Println("You must configure your environment before using `gtd`:")
	fmt.Println("  export GTD_GH_TOKEN=XXXXX")
	fmt.Println("  export GTD_GH_USER=benschw")
	fmt.Println("  export GTD_GH_REPO=gtd")
	fmt.Println("And optionally set a default context:")
	fmt.Println("  export GTD_CONTEXT=@work")

}

func getArgs() []string {
	args := make([]string, 0)

	for i := 0; i < flag.NArg(); i++ {
		args = append(args, flag.Arg(i))
	}
	return args
}

func getEnv() (string, string, string, string, error) {
	var err error
	token := os.Getenv("GTD_GH_TOKEN")
	if token == "" {
		err = fmt.Errorf("Required ENV var missing")
	}
	user := os.Getenv("GTD_GH_USER")
	if user == "" {
		err = fmt.Errorf("Required ENV var missing")
	}
	repo := os.Getenv("GTD_GH_REPO")
	if repo == "" {
		err = fmt.Errorf("Required ENV var missing")
	}
	defaultContext := os.Getenv("GTD_CONTEXT")
	if defaultContext == "" {
		defaultContext = "@global"
	}

	return token, user, repo, defaultContext, err
}

func main() {
	token, user, repo, defaultContext, err := getEnv()
	if err != nil {
		envHelp()
		os.Exit(1)
	}

	flag.Parse()
	args := flag.Args()

	r, err := gtd.ParseArgs(args, defaultContext)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	out, err := gtd.Dispatch(r, gtd.NewGhRepo(user, repo, token))
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	fmt.Print(out)

}
