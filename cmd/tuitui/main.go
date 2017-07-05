package main

import (
	"os"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"fmt"
	"github.com/stephensxu/tuitui"
)

const (
	donaldTrump   = "realDonaldTrump"
	elonMusk      = "elonmusk"
	maxTweetCount = 10
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to your twitter account",
			Action: func(c *cli.Context) error {
				client := tuitui.NewTuituiClient()
				_, err := tuitui.Login(client)

				if err != nil {
					color.Red(err.Error())
				}

				color.Green("Logged in!")
				return nil
			},
		},
		{
			Name: "elon",
			Usage: "Get latest 10 tweets of Elon Musk",
			Action: func(c *cli.Context) error {
				client, err := tuitui.NewAuthenticatedClient()

				if err != nil {
					fmt.Println(err)
				}

				tweets, err := client.GetRecentTweets(elonMusk)
				
				if err != nil {
					color.Red(err.Error())
				}

				for i, tweet := range tweets {
					if i < maxTweetCount {
						color.Green(tweet.Text)
						fmt.Println("\n")
					}
				}
				return nil
			},
		},
		{
			Name: "trump",
			Usage: "Get latest 10 tweets of Donald Trump",
			Action: func(c *cli.Context) error {
				client, err := tuitui.NewAuthenticatedClient()
				
				if err != nil {
					color.Red(err.Error())
				}
				
				tweets, err := client.GetRecentTweets(donaldTrump)
				
				if err != nil {
					color.Red(err.Error())
				}

				for i, tweet := range tweets {
					if i < maxTweetCount {
						color.Green(tweet.Text)
						fmt.Println("\n")
					}
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
