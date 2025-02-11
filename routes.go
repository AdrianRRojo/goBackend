package main

import (
	"net/http"
)

func serveStaticFiles() {
	http.Handle("/View/", http.StripPrefix("/View/", http.FileServer(http.Dir("./View/"))))
}

func getLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./View/index.html")
}
