package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serves static files i.e CSS + JS
	serveStaticFiles()

	// Routes located in routes.go
	http.HandleFunc("/", getLanding)

	//Server
	port := ":8200"
	fmt.Printf("Live at: http://localhost%v", port)
	err := http.ListenAndServe(port, nil)
	if err == nil {
		panic(err)
	}
}
