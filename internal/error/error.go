package error

import (
	"fmt"
	"os"
)

/**
* Parameters error
*/
func ParametersError() {
	fmt.Println(" Wrong usage!")
	fmt.Println("    twitch <stream>")
	os.Exit(1)
}
