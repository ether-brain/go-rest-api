package main

import (
	"fmt"

	handlers "github.com/ether-brain-go/go-rest-api/internal/handlers"
)

func main() {
	fmt.Println("Hello World!")
	handlers.Handle()
}
