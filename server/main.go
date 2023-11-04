package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sam0307/golang-react-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
