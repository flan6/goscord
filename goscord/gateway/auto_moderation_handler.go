package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type AutoModerationRuleCreateHandler struct{}

func (*AutoModerationRuleCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationRuleCreate", ev.Data)
}

type AutoModerationRuleDeleteHandler struct{}

func (*AutoModerationRuleDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleDelete(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationRuleDelete", ev.Data)
}

type AutoModerationRuleUpdateHandler struct{}

func (*AutoModerationRuleUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationRuleUpdate", ev.Data)
}

type AutoModerationActionExecutionHandler struct{}

func (*AutoModerationActionExecutionHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationActionExecution(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationActionExecution", ev.Data)
}
