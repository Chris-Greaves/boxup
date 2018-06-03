package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: 1.0.0")
}

func GetBoxes(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Boxes)
}

func CreateBox(w http.ResponseWriter, r *http.Request) {
	var box Box
	var err error
	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&box)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v", err)
		w.WriteHeader(400)
		return
	}

	err = AddBox(box)
	if err != nil {
		fmt.Fprintf(w, "An Error occured: %v", err)
		if err == ErrBoxConflict {
			w.WriteHeader(409)
			return
		}
		w.WriteHeader(400)
		return
	}

	fmt.Fprintf(w, "%v box created", box.Name)
}
