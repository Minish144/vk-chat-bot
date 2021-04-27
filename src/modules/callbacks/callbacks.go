package callbacks

import (
	"vk-chat-bot/src/modules/messages"

	"github.com/SevereCloud/vksdk/v2/callback"
)

func Init() *callback.Callback {
	callback := callback.NewCallback()
	callback.MessageNew(messages.MessageNewHandler)
	return callback
}
