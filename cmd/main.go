package main

import (
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/LuisHenriqueFA14/twitch-chat-cli/internal/message"
	"github.com/LuisHenriqueFA14/twitch-chat-cli/internal/error"
)

var stream string

func main() {
	if len(os.Args) != 2 {
		error.ParametersError()
	} else {
		stream = os.Args[1]
	}

	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(message.ShowMessage)

	client.Join(stream)
	err := client.Connect()

	if err != nil {
		panic(err)
	}
}
