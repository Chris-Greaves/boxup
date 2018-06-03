package main

import (
	"errors"
	"os"
)

var (
	Boxes          []Box
	ErrBoxConflict = errors.New("a Box with the same name already exists")
)

func AddBox(box Box) (err error) {
	err = CheckNameIsUnique(box.Name)
	if err != nil {
		return err
	}

	err = CheckDirectoryExists(box.Location)
	if err != nil {
		return err
	}

	Boxes = append(Boxes, box)
	return err
}

func CheckNameIsUnique(name string) (err error) {
	for _, box := range Boxes {
		if box.Name == name {
			return ErrBoxConflict
		}
	}
	return nil
}

func CheckDirectoryExists(dir string) (err error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}
	return nil
}
