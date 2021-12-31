package rueidis

import "github.com/rueian/rueidis/internal/cmds"

func NewPubSubHandlers(onConnected PubSubSetup, opt PubSubOption) PubSubHandlers {
	return PubSubHandlers{
		onMessage:      opt.OnMessage,
		onPMessage:     opt.OnPMessage,
		onSubscribed:   opt.OnSubscribed,
		onUnSubscribed: opt.OnUnSubscribed,
		onConnected:    onConnected,
	}
}

type PubSubSetup func(prev error, client DedicatedClient)

type PubSubOption struct {
	OnMessage      func(channel, message string)
	OnPMessage     func(pattern, channel, message string)
	OnSubscribed   func(channel string, active int64)
	OnUnSubscribed func(channel string, active int64)
}

type PubSubHandlers struct {
	onMessage      func(channel, message string)
	onPMessage     func(pattern, channel, message string)
	onSubscribed   func(channel string, active int64)
	onUnSubscribed func(channel string, active int64)
	onConnected    PubSubSetup
}

func (h PubSubHandlers) _install(prev error, builder *cmds.Builder, pick func() conn) {
	if cc := pick(); cc != nil {
		cc.OnDisconnected(func(err error) {
			if err != ErrConnClosing {
				h._install(err, builder, pick)
			}
		})
		go h.onConnected(prev, &dedicatedSingleClient{cmd: builder, wire: cc})
	}
}
func (h PubSubHandlers) installHook(builder *cmds.Builder, pick func() conn) {
	if h.onConnected != nil {
		h._install(nil, builder, pick)
	}
}
