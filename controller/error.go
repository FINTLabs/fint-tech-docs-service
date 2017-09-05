package controller

import "net/http"

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}
