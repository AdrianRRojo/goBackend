package main

import (
	"net/http"
)

func getLanding(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./View/index.html")

}

func serveStaticFiles() {
	http.Handle("/View/", http.StripPrefix("/View/", http.FileServer(http.Dir("./View/"))))
}
