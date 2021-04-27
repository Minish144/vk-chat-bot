package vk

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/spf13/viper"
)

func Init() *api.VK {
	groupToken := viper.GetString("vk.token")
	return api.NewVK(groupToken)
}
