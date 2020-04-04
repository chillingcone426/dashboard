package cache

import (
	"encoding/json"
	"github.com/apex/log"
)

type TicketCloseMessage struct {
	Uuid   string
	User   uint64
	Reason string
}

func (c *RedisClient) PublishTicketClose(ticket string, userId uint64, reason string) {
	settings := TicketCloseMessage{
		Uuid:   ticket,
		User:   userId,
		Reason: reason,
	}

	encoded, err := json.Marshal(settings); if err != nil {
		log.Error(err.Error())
		return
	}

	c.Publish("tickets:close", string(encoded))
}

