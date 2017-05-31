package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// Input channel
type CommandChannel struct {
	Channel *slack.Channel
	Event   *slack.MessageEvent
	UserID  string
}

// Output channel
type ReplyChannel struct {
	Channel      *slack.Channel
	Attachment   *slack.Attachment
	DisplayTitle string
}

var (
	api *slack.Client
	//	userMessages      Messages
	botID             string
	botCommandChannel chan *CommandChannel
	botReplyChannel   chan ReplyChannel
)

func handleBotCommands(c chan ReplyChannel) {
	var rc ReplyChannel

	for {
		botChannel := <-botCommandChannel
		rc.Channel = botChannel.Channel
		//reply := &slack.Msg{}
		//commandArray := strings.Fields(botChannel.Event.Text)
		// switch commandArray[1] {

	}
}

func handleBotReply() {
	for {
		ac := <-botReplyChannel
		params := slack.PostMessageParameters{}
		params.AsUser = true
		params.Attachments = []slack.Attachment{*ac.Attachment}
		_, _, errPostMessage := api.PostMessage(ac.Channel.Name, ac.DisplayTitle, params)
		if errPostMessage != nil {
			log.Fatal(errPostMessage)
		}
	}
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "slack" {
		fmt.Fprintf(os.Stderr, "usage: concierge slack bot-token\n")
		os.Exit(1)
	}
	api := slack.New(os.Args[2])
	rtm := api.NewRTM()

	botCommandChannel = make(chan *CommandChannel)
	botReplyChannel = make(chan ReplyChannel)

	//userMessages = make(Messages, 0)

	go rtm.ManageConnection()
	go handleBotCommands(botReplyChannel)
	go handleBotReply()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.MessageEvent:
				channelInfo, err := api.GetChannelInfo(ev.Channel)
				if err != nil {
					log.Fatalln(err)
				}

				command := &CommandChannel{
					Channel: channelInfo,
					Event:   ev,
					UserID:  ev.User,
				}

				if ev.Type == "message" && strings.HasPrefix(ev.Text, "<@"+botID+">") {
					botCommandChannel <- command
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}

}
