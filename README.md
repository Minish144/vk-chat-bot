# VK Chat Bot

## How to

### **1.** Getting a token
First of all you have to generate a token of your bot. [Here's the guide](https://vk.com/dev/access_token)

### **2.** Making a config file
Then you create config file, following `config.example.yml`

### **3.** Download all dependencies
Just run `go mod download`

### **4.** Running the project
Run `go run main.go`, if you did everything according to the instructions above you'll something like this:

![go-run-demo](https://i.imgur.com/YpSTsV7.png)

### **5.** Coding message handlers
Now it's time to code your own message handlers. In `src/modules/messages/handler.go` you have everything that you need for it. `vkApi` is a module which provides you VK API methods, [here's the library itself](https://github.com/SevereCloud/vksdk).

Both conversationMessageHandler and personalMessageHandler are function where you handle your messages. Messages, sent to bot from vk conversation, you handle in conversationMessageHandler function, personal messages you handle in personalMessageHandler.

```go
var vkApi *api.VK

. . .

func conversationMessageHandler(obj *events.MessageNewObject) {
	// msg := (*obj).Message
	// client := (*obj).ClientInfo
	fmt.Println("Conversation message received!")
}

func personalMessageHandler(obj *events.MessageNewObject) {
	// msg := (*obj).Message
	// client := (*obj).ClientInfo
	fmt.Println("Personal message received!")
}
```

### **6.** Linking group and server app
So the very last step is to link your project and your group, [that's how you do it](https://vk.com/dev/callback_api)