# tuitui

## Summary

Show the latest tweets from Elon Musk and Donald Trump!

tuitui is a CLI tool that pulls the latest tweets from Elon Musk and Donald Trump to your command line so you can check out what they're saying during your brief 3 second break before moving on to the next task.

This is really inspired to solve a problem of my own. Elon Musk and Donald Trump are the two twitter stars that I constant pay attention to(paying attention doesn't mean I'm big fan, just means it's a source of information I'm always checking, sometimes just to entertain myself)

## Usage

### Installation

Download latest executables for OSX and Linux from [Github releases](https://github.com/stephensxu/tuitui/releases).

So far I've only tested and confirmed it's working on darwin_amd64

After downloading put the executable in your `PATH` and you can use with `tuitui`

### Twitter App

`tuitui login` will promp you to enter twitter credentails, you'll only need to do this once.

This app requires certain credentials from twitter in order for it to work(credential to use twitter API).

You will need a Twitter App in order to obtain credentials from twitter. If you don't have one yet, it takes 5 minutes to create an empty twitter app [here](https://apps.twitter.com/)(no coding required)

Once you've done that, go to `Keys and Access Tokens` tab and obtain the following credentials:

Consumer Key
Consumer Secret
Access Token
Access Token Secret

Enter these credentials as promped.


`tuitui elon`

`tuitui trump`

Enjoy!
