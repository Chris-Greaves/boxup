package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"boxup/boxup-server/boxmanagment"
)

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: 1.0.0")
}

func GetBoxes(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(boxmanagment.Boxes)
}

func CreateBox(w http.ResponseWriter, r *http.Request) {
	var box boxmanagment.Box
	var err error
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&box)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v", err)
		w.WriteHeader(400)
		return
	}

	err = boxmanagment.AddBox(box)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v", err)
		if err == boxmanagment.ErrBoxConflict {
			w.WriteHeader(409)
			return
		}
		w.WriteHeader(400)
		return
	}

	fmt.Fprintf(w, "%v box created", box.Name)
}

func GetBox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	err := boxmanagment.GetBoxZip(name, w)

	if err != nil {
		if err == boxmanagment.ErrBoxDoesntExist {
			fmt.Fprintf(w, "An Error occured: %v", err)
			w.WriteHeader(400)
			return
		}

		fmt.Fprintf(w, "An Error occured: %v", err)
		w.WriteHeader(500)
		return
	}
}
