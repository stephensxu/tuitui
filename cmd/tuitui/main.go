package main

import (
	"os"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to your twitter account",
			Action: func(c *cli.Context) error {
				color.Green("Logged in!")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
