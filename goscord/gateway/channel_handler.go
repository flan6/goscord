package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type ChannelCreateHandler struct{}

func (*ChannelCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("channelCreate", ev.Data)
}

type ChannelUpdateHandler struct{}

func (*ChannelUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("channelUpdate", ev.Data)
}

type ChannelDeleteHandler struct{}

func (*ChannelDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelDelete(s.rest, data)

	if err != nil {
		return
	}

	ev.Data, _ = s.State().Channel(ev.Data.Id)

	s.State().RemoveChannel(ev.Data)

	s.Bus().Publish("channelDelete", ev.Data)
}

type ChannelPinsUpdateHandler struct{}

func (*ChannelPinsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelPinsUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("channelPinsUpdate", ev.Data)
}

type ThreadCreateHandler struct{}

func (*ThreadCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("threadCreate", ev.Data)
}

type ThreadUpdateHandler struct{}

func (*ThreadUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("threadUpdate", ev.Data)
}

type ThreadDeleteHandler struct{}

func (*ThreadDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadDelete(s.rest, data)

	if err != nil {
		return
	}

	ev.Data, _ = s.State().Channel(ev.Data.Id)

	s.State().RemoveChannel(ev.Data)

	s.Bus().Publish("threadDelete", ev.Data)
}

type ThreadListSyncHandler struct{}

func (*ThreadListSyncHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadListSync(s.rest, data)

	if err != nil {
		return
	}

	for _, thread := range ev.Data.Threads {
		s.State().AddChannel(thread)
	}

	s.Bus().Publish("threadListSync", ev.Data)
}

type ThreadMemberUpdateHandler struct{}

func (*ThreadMemberUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadMemberUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Update thread member?

	s.Bus().Publish("threadMemberUpdate", ev.Data)
}

type ThreadMembersUpdateHandler struct{}

func (*ThreadMembersUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadMembersUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Update thread members?

	s.Bus().Publish("threadMembersUpdate", ev.Data)
}
