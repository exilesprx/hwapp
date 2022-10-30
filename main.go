package main

import (
	"fmt"
	"go-hwapp/messages"
)

func main() {
	welcome := messages.Welcome{Message: "Welcome"}

	fmt.Println(welcome.Message)
}
