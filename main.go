package main

import (
	"fmt"
	"log"

	"github.com/sleepynut/GO-module-lesson/greets"
)

func main() {
	fmt.Println(greets.Hello("sleepynut"))

	message, err := greets.Hello("")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(message)
}
