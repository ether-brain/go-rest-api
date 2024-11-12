package main

import (
	"fmt"

	handlers "github.com/ether-brain/go-rest-api/internal/handlers"
)

func main() {
	status, response, err := handlers.Handle()
	fmt.Println(*status, *response, err)
}
