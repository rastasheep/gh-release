package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
)

var repo_name, release_ver string

func fatal(msg string) {
	println("Error:", msg)
	os.Exit(1)
}

func requireArgs(num int, args cli.Args) {
	if len(args) < num {
		fatal("You must provide repository name and release version.")
	}
}

func checkReleaseDir(path string) {
	files, _ := filepath.Glob(path + "/*")
	if len(files) == 0 {
		fatal("Release dir empty.")
	}
}

func Create(c *cli.Context) {
	requireArgs(2, c.Args())
	checkReleaseDir(c.String("release_path"))

	println("Created release", c.Args().First())
}

func Destroy(c *cli.Context) {
	requireArgs(2, c.Args())

	println("Destroyed release: ", c.Args().First())
}

func main() {
	app := cli.NewApp()

	app.Name = "gh-release"
	app.Version = "0.0.1"
	app.Author = "Aleksandar Diklic"
	app.Email = "rastasheep@gmail.com"

	app.Usage = "fight the loneliness!"
	app.Action = cli.ShowAppHelp

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Value:  "",
			Usage:  "GitHub access token",
			EnvVar: "GITHUB_ACCESS_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		{
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
		},
		{
			Name:      "destroy",
			ShortName: "d",
			Usage:     "Destroy a release",
			Action:    Destroy,
		},
	}

	app.Run(os.Args)
}
