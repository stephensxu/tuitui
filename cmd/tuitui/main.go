package main

import (
	"os"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"fmt"
	"github.com/stephensxu/tuitui"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to your twitter account",
			Action: func(c *cli.Context) error {
				_, err := tuitui.Login()

				if err != nil {
					fmt.Println(err)	
				}

				color.Green("Logged in!")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
