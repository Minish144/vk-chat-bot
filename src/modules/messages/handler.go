package messages

import (
	"context"
	"log"

	"github.com/SevereCloud/vksdk/v2/events"
)

func Handler(ctx context.Context, obj events.MessageNewObject) {
	log.Print(obj.Message.Text)
}
