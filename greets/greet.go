package greets

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hi, %s. Welcome!", name)
}
