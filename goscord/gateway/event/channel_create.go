package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type ChannelCreate struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelCreate(rest *rest.Client, data []byte) (*ChannelCreate, error) {
	pk := new(ChannelCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
