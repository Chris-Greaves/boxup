package boxmanagment

import (
	"compress/gzip"
	"io"
	"errors"
	"os"
)

var (
	Boxes          		= map[string]Box{}
	ErrBoxConflict 		= errors.New("a Box with the same name already exists")
	ErrBoxDoesntExist	= errors.New("Box doesnt exist")
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

	Boxes[box.Name] = box

	//Boxes = append(Boxes, box)
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

func GetBox(name string) (io.Reader, error) {
	if _, ok := Boxes[name]; !ok {
		return nil, ErrBoxDoesntExist
	}

	node, err := os.Open(Boxes[name].Location)
	if err != nil {
		return nil, err
	}

	return gzip.NewReader(node)
}