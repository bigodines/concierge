package types

import (
	"github.com/nlopes/slack"
)

type IMessage interface {
	Type() string
	Channel() Channel
	User() User
}

type Message struct {
	Type    string
	Channel Channel
	User    User
}

type User struct {
	ID   string
	Name string
}

type Channel struct {
	ID   string
	Name string
}

type Event struct {
}

// Input channel
type InputInfo struct {
	Channel *slack.Channel
	//Event   *slack.MessageEvent
	UserID string
}

// Output channel
type OutputInfo struct {
	Channel *slack.Channel
	Message *slack.Msg
}
