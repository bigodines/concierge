package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// Input channel
type InputInfo struct {
	Channel *slack.Channel
	Event   *slack.MessageEvent
	UserID  string
}

// Output channel
type OutputInfo struct {
	Channel *slack.Channel
	Message *slack.Msg
}

type Msg *slack.Msg

var (
	api             *slack.Client
	botID           string
	botInputChannel chan *InputInfo
	botReplyChannel chan *OutputInfo
)

func handleBotCommands(c chan *OutputInfo) {
	var rc OutputInfo

	for {
		command := <-botInputChannel

		reply, err := buildReply(command)
		if err != nil {
			log.Error("Can't parse command.")
			log.Debug("%+v", reply)
		}
		rc.Channel = command.Channel
		rc.Message = reply
		c <- &rc
	}
}

func buildReply(command *InputInfo) (Msg, error) {
	who := command.Event.Msg.User
	msg := command.Event.Msg.Text
	if strings.Contains(Text, "giants game tonight") {
		// ..
		// ..
		reply = &slack.Msg{
			User: "Concierge",
			Text: "I can get you two tickets to the giants game for $9 dollars",
			File: File{
				"id": "abc",
				"title": "G",
				"url": "http://bigode.me,"
			}
		}
		
	}
	return &slack.Msg{
		Text: "Hello",
		User: "Concierge",
	}, nil

}

func handleBotReply(rtm *slack.RTM) {
	for {
		ac := <-botReplyChannel
		rtm.SendMessage(rtm.NewOutgoingMessage(ac.Message.Text, ac.Channel.ID))
	}
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "slack" {
		fmt.Fprintf(os.Stderr, "usage: concierge slack bot-token\n")
		os.Exit(1)
	}
	api := slack.New(os.Args[2])
	rtm := api.NewRTM()

	botInputChannel = make(chan *InputInfo)
	botReplyChannel = make(chan *OutputInfo)

	go rtm.ManageConnection()
	go handleBotCommands(botReplyChannel)
	go handleBotReply(rtm)

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID

			case *slack.MessageEvent:
				fmt.Printf("Got message: %s\n", msg.Data)
				channelInfo, err := api.GetChannelInfo(ev.Channel)
				if err != nil {
					log.Fatalln(err)
				}

				command := &InputInfo{
					Channel: channelInfo,
					Event:   ev,
					UserID:  ev.User,
				}

				fmt.Printf("Type: %s, Text: %s, botID: %s\n", ev.Type, ev.Text, botID)

				if ev.Type == "message" && strings.HasPrefix(ev.Text, "<@"+botID+">") {
					fmt.Printf("Valid message\n")
					botInputChannel <- command
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
