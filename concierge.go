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
	Channel *slack.Channel
	Message *slack.Msg
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

		reply := &slack.Msg{
			Text: "Hello",
			User: "Concierge",
		}
		rc.Channel = botChannel.Channel
		rc.Message = reply
		c <- rc
		fmt.Printf("Pushed to channel\n")

		//reply := &slack.Msg{}

		//commandArray := strings.Fields(botChannel.Event.Text)
		// switch commandArray[1] {

	}
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

	botCommandChannel = make(chan *CommandChannel)
	botReplyChannel = make(chan ReplyChannel)

	//userMessages = make(Messages, 0)

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

				command := &CommandChannel{
					Channel: channelInfo,
					Event:   ev,
					UserID:  ev.User,
				}

				fmt.Printf("Type: %s\n", ev.Type)
				fmt.Printf("Text: %s\n", ev.Text)
				fmt.Printf("botID: %s\n", botID)

				if ev.Type == "message" && strings.HasPrefix(ev.Text, "<@"+botID+">") {
					fmt.Printf("Valid message\n")
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
