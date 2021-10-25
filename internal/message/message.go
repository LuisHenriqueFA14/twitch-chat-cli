package message

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

/**
* Return blue text to unix patterns.
*/
func blueText(text string) string {
	return "\033[34m" + text + "\033[0m"
}

/**
* Return white text to unix patterns.
*/
func whiteText(text string) string {
	return "\033[37m" + text + "\033[0m"
}

/**
* Prints formatted text.
*/
func ShowMessage(message twitch.PrivateMessage) {
	fmt.Println(blueText(message.User.Name) + ": " + whiteText(message.Message))
}
