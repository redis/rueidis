package rueidis

import "github.com/rueian/rueidis/internal/cmds"

// NewPubSubOption creates a PubSubOption for client initialization.
// The onConnected callback is called when the connection to a redis node established including auto reconnect.
// One should subscribe to channel in the onConnected callback to have it resubscribe after auto reconnection.
func NewPubSubOption(onConnected func(prev error, client DedicatedClient), cbs PubSubHandler) PubSubOption {
	return PubSubOption{
		onMessage:      cbs.OnMessage,
		onPMessage:     cbs.OnPMessage,
		onSubscribed:   cbs.OnSubscribed,
		onUnSubscribed: cbs.OnUnSubscribed,
		onConnected:    onConnected,
	}
}

// PubSubHandler is called in single thread when receiving corresponding PubSub messages from the redis node.
// These callbacks must not take too long to complete, otherwise it will impact Client performance.
// They usually just put message to other golang channels.
type PubSubHandler struct {
	OnMessage      func(channel, message string)
	OnPMessage     func(pattern, channel, message string)
	OnSubscribed   func(channel string, active int64)
	OnUnSubscribed func(channel string, active int64)
}

// PubSubOption should be created from the NewPubSubOption()
type PubSubOption struct {
	onMessage      func(channel, message string)
	onPMessage     func(pattern, channel, message string)
	onSubscribed   func(channel string, active int64)
	onUnSubscribed func(channel string, active int64)
	onConnected    func(prev error, client DedicatedClient)
}

func (h PubSubOption) _install(prev error, builder *cmds.Builder, pick func() conn) {
	if cc := pick(); cc != nil {
		cc.OnDisconnected(func(err error) {
			if err != ErrClosing {
				h._install(err, builder, pick)
			}
		})
		go h.onConnected(prev, &dedicatedSingleClient{cmd: builder, wire: cc})
	}
}
func (h PubSubOption) installHook(builder *cmds.Builder, pick func() conn) {
	if h.onConnected != nil {
		h._install(nil, builder, pick)
	}
}
