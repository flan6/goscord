package packet

import (
	"time"

	"github.com/goccy/go-json"
)

type Hello struct {
	*Packet
	Data struct {
		HeartbeatInterval time.Duration `json:"heartbeat_interval"`
	} `json:"d"`
}

func NewHello(data []byte) (*Hello, error) {
	var packet Hello

	err := json.Unmarshal(data, &packet)

	if err != nil {
		return nil, err
	}

	return &packet, nil
}
