package main

import (
	"github.com/urfave/cli"
	"os"
	"time"
)

func initCLI() {

	todaysDate := time.Now().Format("20060102")
	app := cli.NewApp()

	app.Name = "Openshift Hellper"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mangirdas Judeikis",
			Email: "info@judeikis.lt",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "imagestream",
			Aliases: []string{"is"},
			Usage:   "modify ImageStream with new version of images",
			Subcommands: []cli.Command{
				{
					Name:  "edit",
					Usage: "pass imageStream details to modify",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file",
							Usage: "--file=path/to/your/file.json",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "--name=name_of_is",
						},
						cli.BoolFlag{
							Name:  "latest",
							Usage: "--latest=false",
						},
						cli.StringFlag{
							Name:  "tag",
							Value: todaysDate,
							Usage: "--tag=" + todaysDate,
						},
					},
					Action: func(c *cli.Context) error {
						UpdateISFile(c.String("file"), c.String("name"), c.String("tag"), c.Bool("latest"))
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
