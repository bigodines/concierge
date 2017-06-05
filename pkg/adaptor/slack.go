package adaptor

type message struct {
	Type    string
	Channel Channel
	User    User
}

func (m message) getType() string {
	return m.Type
}
func (m message) getChannel() Channel {
	return m.Channel
}
