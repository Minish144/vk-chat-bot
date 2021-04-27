package messages

import (
	"context"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

var vkApi *api.VK = nil

func Init(vk *api.VK) error {
	vkApi = vk
	params := api.Params{
		"user_ids": []int{1},
	}
	_, err := vkApi.UsersGet(params)
	return err
}

func MessageNewHandler(ctx context.Context, obj events.MessageNewObject) {
	if obj.Message.ID == 0 {
		conversationMessageHandler(&obj)
	} else {
		personalMessageHandler(&obj)
	}

}

func conversationMessageHandler(obj *events.MessageNewObject) {
	// msg := (*obj).Message
	// client := (*obj).ClientInfo
}

func personalMessageHandler(obj *events.MessageNewObject) {
	// TODO
}
