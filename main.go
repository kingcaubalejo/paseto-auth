package main

import (
	"log"
	"net/http"
	"fmt"

	r "qualitech.paseto-auth/route"
)


func main() {
	router := r.InitializeRoutes()
	
	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
