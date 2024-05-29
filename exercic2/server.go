package exercice2

import (
	"fmt"
	"net/http"
)

const port =":8080"
func Home(w http.ResponseWriter, r *http.Request){

fmt.Fprint( w, "hello word")
}

func Server() {
	http.HandleFunc("/",Home) 
		
	http.ListenAndServe(port, nil)
}