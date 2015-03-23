package cmd

import (
	"errors"

	"github.com/codegangsta/cli"
)

type Cmd struct {
	Name        string
	Context     *cli.Context
	Slug        string
	Version     string
	ReleasePath string
	Token       string
}

func Create(c *cli.Context) (*Cmd, error) {
	args := c.Args()
	if len(args) < 2 {
		return nil, errors.New("You must provide repository name and release version.")
	}

	return &Cmd{
		Name:        "create",
		Context:     c,
		Slug:        args.First(),
		Version:     args.Get(1),
		ReleasePath: c.String("release_path"),
		Token:       c.GlobalString("token"),
	}, nil
}
