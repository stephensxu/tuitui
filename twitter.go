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

func (t *TuituiClient) GetRecentTweets(screenName string) ([]twitter.Tweet, error) {
	timeLineParams := &twitter.UserTimelineParams{
		ScreenName: screenName,
	}

	tweets, _, err := t.client.Timelines.UserTimeline(timeLineParams)

	if err != nil {
		return []twitter.Tweet{}, err
	}

	return tweets, nil
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

	_, _, err = twitterClient.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		fmt.Println(err)
		return false	
	}

	t.client = twitterClient
	t.credentials = credentials

	return true
}
