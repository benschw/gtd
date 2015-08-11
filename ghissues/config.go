package ghissues

import (
	"fmt"
	"os"
)

type Config struct {
	Token string
	User  string
	Repo  string
}

func DefaultConfig() (*Config, error) {
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
	return c, nil
}
