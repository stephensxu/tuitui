package tuitui

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TuituiClient struct {
	client *twitter.Client
	credentials map[string]string
}

func NewTuituiClient() *TuituiClient {
	return &TuituiClient{}
}

// Authenticate : Authenticates using the fields passed in
func (t *TuituiClient) Authenticate() bool {
	credentials, err := Load()
	if err != nil {
		fmt.Println(err)	
	}

	config := oauth1.NewConfig(credentials["consumerKey"], credentials["consumerSecret"])
	token := oauth1.NewToken(credentials["accessToken"], credentials["accessSecret"])

	httpClient := config.Client(oauth1.NoContext, token)

	twitterClient := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := twitterClient.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		fmt.Println(err)
		return false	
	}

	fmt.Println("ScreenName: ", user.ScreenName)

	t.client = twitterClient
	t.credentials = credentials

	return true
}
