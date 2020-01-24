// Package twofer implements a simple function to create a message.
package twofer

import (
	"fmt"
	"strings"
)

const defaultName = "you"

// ShareWith creates a 'two for one' message with the name parameter.
// if the parameters is empty, the default name will be used.
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name =  "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
