package main

import (
	"net/http"
)

func getLanding(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./View/index.html")

}
