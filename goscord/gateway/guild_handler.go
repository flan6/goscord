package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type PresenceUpdateHandler struct{}

func (*PresenceUpdateHandler) Handle(s *Session, Data []byte) {
	ev, err := event.NewPresenceUpdate(s.rest, Data)

	if err != nil {
		return
	}

	// ToDo : Add a method in state to track presences

	s.Bus().Publish("presenceUpdate", ev.Data)
}

type GuildCreateHandler struct{}

func (*GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddGuild(ev.Data)

	s.Bus().Publish("guildCreate", ev.Data)
}

type GuildUpdateHandler struct{}

func (*GuildUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddGuild(ev.Data)

	s.Bus().Publish("guildUpdate", ev.Data)
}

type GuildDeleteHandler struct{}

func (*GuildDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildDelete(s.rest, data)

	if err != nil {
		return
	}

	s.State().RemoveGuild(ev.Data)

	s.Bus().Publish("guildDelete", ev.Data)
}

type GuildBanAddHandler struct{}

func (*GuildBanAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildBanAdd(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.State().Guild(ev.Data.GuildId)
	user := ev.Data.User

	if err != nil {
		return
	}

	s.Bus().Publish("guildBanAdd", guild, user)
}

type GuildBanRemoveHandler struct{}

func (*GuildBanRemoveHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildBanRemove(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.State().Guild(ev.Data.GuildId)
	user := ev.Data.User

	if err != nil {
		return
	}

	s.Bus().Publish("guildBanRemove", guild, user)
}

type GuildEmojisUpdateHandler struct{}

func (*GuildEmojisUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildEmojisUpdate(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.State().Guild(ev.Data.GuildId)
	if err != nil {
		return
	}

	guild.Emojis = ev.Data.Emojis

	s.Bus().Publish("guildEmojisUpdate", ev.Data)
}

type GuildStickersUpdateHandler struct{}

func (*GuildStickersUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildStickersUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Cache stickers?

	s.Bus().Publish("guildStickersUpdate", ev.Data)
}

type GuildIntegrationsUpdateHandler struct{}

func (*GuildIntegrationsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildIntegrationsUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("guildIntegrationsUpdate", ev.Data)
}

type GuildMemberAddHandler struct{}

func (*GuildMemberAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMemberAdd(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Implement Member count?

	s.State().AddMember(ev.Data.GuildId, ev.Data)

	s.Bus().Publish("guildMemberAdd", ev.Data)
}

type GuildMemberRemoveHandler struct{}

func (*GuildMemberRemoveHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMemberRemove(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Implement Member count?

	s.State().RemoveMember(ev.Data.GuildId, ev.Data.User.Id)

	s.Bus().Publish("guildMemberRemove", ev.Data)
}

type GuildMemberUpdateHandler struct{}

func (*GuildMemberUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMemberUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddMember(ev.Data.GuildId, ev.Data)

	s.Bus().Publish("guildMemberUpdate", ev.Data)
}

type GuildMembersChunkHandler struct{}

func (*GuildMembersChunkHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMembersChunk(s.rest, data)

	if err != nil {
		return
	}

	for _, member := range ev.Data.Members {
		s.State().AddMember(ev.Data.GuildId, member)
	}

	s.Bus().Publish("guildMembersChunk", ev.Data)
}

type GuildRoleCreateHandler struct{}

func (*GuildRoleCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildRoleCreate(s.rest, data)

	if err != nil {
		return
	}

	err = s.State().AddRole(ev.Data.GuildId, ev.Data.Role)

	if err != nil {
		return
	}

	s.Bus().Publish("guildRoleCreate", ev.Data)
}

type GuildRoleUpdateHandler struct{}

func (*GuildRoleUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildRoleUpdate(s.rest, data)

	if err != nil {
		return
	}

	err = s.State().AddRole(ev.Data.GuildId, ev.Data.Role)

	if err != nil {
		return
	}

	s.Bus().Publish("guildRoleUpdate", ev.Data)
}

type GuildRoleDeleteHandler struct{}

func (*GuildRoleDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildRoleDelete(s.rest, data)

	if err != nil {
		return
	}

	err = s.State().RemoveRole(ev.Data.GuildId, ev.Data.RoleId)

	if err != nil {
		return
	}

	s.Bus().Publish("guildRoleDelete", ev.Data)
}
