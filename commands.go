package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/rastasheep/gh-release/cmd"
	"github.com/rastasheep/gh-release/github"
)

func Create(c *cli.Context) {
	cmd, err := cmd.Create(c)
	FatalIf(cmd, err)

	release, err := github.CreateRelease(cmd.Slug, cmd.Version, cmd.ReleasePath, cmd.Token)
	FatalIf(cmd, err)

	err = release.Deploy()
	FatalIf(cmd, err)

	println("Created release", release)
}

func Destroy(c *cli.Context) {
	println("Destroyed release")
}

func FatalIf(c *cmd.Cmd, err error) {
	if err != nil {
		if c != nil {
			cli.ShowCommandHelp(c.Context, c.Name)
		}
		
		println("ERROR:\n   ", err.Error())
		os.Exit(1)
	}
}
