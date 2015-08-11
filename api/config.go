package api

import (
	"fmt"
	"os"
)

func EnvHelp() {
	fmt.Println("You must configure your environment before using `gtd`:")
	fmt.Println("  export GTD_GH_TOKEN=XXXXX")
	fmt.Println("  export GTD_GH_USER=benschw")
	fmt.Println("  export GTD_GH_REPO=gtd")
	fmt.Println("And optionally set a default context:")
	fmt.Println("  export GTD_CONTEXT=@work")
}

type Config struct {
	Token   string
	User    string
	Repo    string
	Context string
}

func NewDefaultConfig() (*Config, error) {
	c := &Config{}

	c.Token = os.Getenv("GTD_GH_TOKEN")
	if c.Token == "" {
		return c, fmt.Errorf("Required ENV var missing")
	}
	c.User = os.Getenv("GTD_GH_USER")
	if c.User == "" {
		return c, fmt.Errorf("Required ENV var missing")
	}
	c.Repo = os.Getenv("GTD_GH_REPO")
	if c.Repo == "" {
		return c, fmt.Errorf("Required ENV var missing")
	}
	c.Context = os.Getenv("GTD_CONTEXT")
	if c.Context == "" {
		c.Context = "@global"
	}
	return c, nil
}
