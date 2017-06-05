package adaptor

import (
	"concierge/pkg/types"
)

func (m types.Imessage) getType() string {
	return m.Type
}
func (m types.Imessage) getChannel() types.Channel {
	return m.Channel
}
