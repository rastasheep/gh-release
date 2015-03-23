package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var version string

var Commands = []cli.Command{
	cmdCreate,
	cmdDestroy,
}

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:   "token, t",
		Usage:  "github token",
		EnvVar: "GITHUB_TOKEN",
	},
}

var cmdCreate = cli.Command{
	Name:      "create",
	ShortName: "c",
	Usage:     "Create a new release",
	Action:    Create,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "release_path",
			Value: "release",
			Usage: "Directory with assets",
		}},
}

var cmdDestroy = cli.Command{
	Name:      "destroy",
	ShortName: "d",
	Usage:     "Destroy a release",
	Action:    Destroy,
}

func main() {
	app := cli.NewApp()

	app.Name = "gh-release"
	app.Version = version
	app.Author = "Aleksandar Diklic - https://github.com/rastasheep"
	app.Email = "rastasheep@gmail.com"

	app.Usage = "fight the loneliness!"
	app.Action = cli.ShowAppHelp
	app.Commands = Commands
	app.Flags = Flags

	app.Run(os.Args)
}
