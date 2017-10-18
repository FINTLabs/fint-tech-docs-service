package controller // import "github.com/FINTprosjektet/fint-tech-docs-service/controller"

import "net/http"

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func assets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/assets/")
}

func files(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/")
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

