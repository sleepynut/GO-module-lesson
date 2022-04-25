package main

import (
	"fmt"
	"log"

	"github.com/sleepynut/GO-module-lesson/greets"
)

func main() {
	message, _ := greets.Hello("sleepynut")
	fmt.Println(message)

	message, err := greets.Hello("")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(message)
}
