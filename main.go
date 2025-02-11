package main

import (
	"fmt"
	"goBackend/View/controllers"
	"net/http"
)

func main() {
	// Serves static files i.e CSS + JS
	serveStaticFiles()

	// GET Routes located in routes.go
	http.HandleFunc("/", getLanding)

	// POST routes in controllers/Register.go
	http.HandleFunc("/submit", controllers.PostForm)

	//Server
	port := ":8200"
	fmt.Printf("Live at: http://localhost%v \n", port)
	err := http.ListenAndServe(port, nil)
	if err == nil {
		fmt.Printf("Error -- %v", err)
	}
}
