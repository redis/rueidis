package rueidis

import "github.com/rueian/rueidis/internal/cmds"

func NewPubSubOption(onConnected func(prev error, client DedicatedClient), cbs PubSubHandler) PubSubOption {
	return PubSubOption{
		onMessage:      cbs.OnMessage,
		onPMessage:     cbs.OnPMessage,
		onSubscribed:   cbs.OnSubscribed,
		onUnSubscribed: cbs.OnUnSubscribed,
		onConnected:    onConnected,
	}
}

type PubSubHandler struct {
	OnMessage      func(channel, message string)
	OnPMessage     func(pattern, channel, message string)
	OnSubscribed   func(channel string, active int64)
	OnUnSubscribed func(channel string, active int64)
}

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
			if err != ErrConnClosing {
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
