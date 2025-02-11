package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getLanding)
	port := ":8200"

	fmt.Printf("Live at: http://localhost%v", port)
	err := http.ListenAndServe(port, nil)
	if err == nil {

	}
}
