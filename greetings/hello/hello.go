package main

import (
	"fmt"

	"github.com/MikHup/learn-go/greetings/greetings"
)

func main() {
	message := greetings.Hello("Michael")
	fmt.Println(message)
}
