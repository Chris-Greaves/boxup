package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: 1.0.0")
}

func CreateBox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	
	fmt.Fprintf(w, "Box created %v", name)
}