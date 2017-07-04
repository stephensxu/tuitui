package main

import (
	"os"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"fmt"
	"github.com/stephensxu/tuitui"
)

const (
	donaldTrump = "realDonaldTrump"
	elonMusk    = "elonmusk"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to your twitter account",
			Action: func(c *cli.Context) error {
				fmt.Println("trying to login")

				client := tuitui.NewTuituiClient()
				_, err := tuitui.Login(client)
				
				if err != nil {
					color.Red(err.Error())
					return err
				}

				color.Green("Logged in!")
				return nil
			},
		},
		{
			Name: "elon",
			Usage: "Get latest timeline of Elon Musk",
			Action: func(c *cli.Context) error {
				client, err := tuitui.NewAuthenticatedClient()

				if err != nil {
					fmt.Println(err)	
				}

				tweet, err := client.GetTimeline(elonMusk)
				
				if err != nil {
					fmt.Println(err)	
				}

				color.Green(tweet)
				return nil
			},
		},
		{
			Name: "trump",
			Usage: "Get latest timeline of Donald Trump",
			Action: func(c *cli.Context) error {
				client, err := tuitui.NewAuthenticatedClient()
				
				if err != nil {
					fmt.Println(err)	
				}

				tweet, err := client.GetTimeline(donaldTrump)
				
				if err != nil {
					fmt.Println(err)	
				}

				color.Green(tweet)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
