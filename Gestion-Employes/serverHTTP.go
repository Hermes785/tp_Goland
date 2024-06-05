package Employes

import (
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./././exercic2/index.html")
}

func ServerHTTP() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(port, nil)
}
