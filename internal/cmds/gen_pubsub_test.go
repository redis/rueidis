// Code generated DO NOT EDIT

package cmds

import "testing"

func pubsub0(s Builder) {
	s.Psubscribe().Pattern("1").Pattern("1").Build()
	s.Publish().Channel("1").Message("1").Build()
	s.PubsubChannels().Pattern("1").Build()
	s.PubsubChannels().Build()
	s.PubsubHelp().Build()
	s.PubsubNumpat().Build()
	s.PubsubNumsub().Channel("1").Channel("1").Build()
	s.PubsubNumsub().Build()
	s.PubsubShardchannels().Pattern("1").Build()
	s.PubsubShardchannels().Build()
	s.PubsubShardnumsub().Channel("1").Channel("1").Build()
	s.PubsubShardnumsub().Build()
	s.Punsubscribe().Pattern("1").Pattern("1").Build()
	s.Punsubscribe().Build()
	s.Spublish().Channel("1").Message("1").Build()
	s.Ssubscribe().Channel("1").Channel("1").Build()
	s.Subscribe().Channel("1").Channel("1").Build()
	s.Sunsubscribe().Channel("1").Channel("1").Build()
	s.Sunsubscribe().Build()
	s.Unsubscribe().Channel("1").Channel("1").Build()
	s.Unsubscribe().Build()
}

func TestCommand_InitSlot_pubsub(t *testing.T) {
	var s = NewBuilder(InitSlot)
	pubsub0(s)
}

func TestCommand_NoSlot_pubsub(t *testing.T) {
	var s = NewBuilder(NoSlot)
	pubsub0(s)
}
