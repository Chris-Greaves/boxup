package boxmanagment

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	Boxes             = map[string]Box{}
	ErrBoxConflict    = errors.New("a Box with the same name already exists")
	ErrBoxDoesntExist = errors.New("Box doesnt exist")
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

func GetBoxes() []Box {
	var boxes []Box
	for _, box := range Boxes {
		boxes = append(boxes, box)
	}

	return boxes
}

func GetBoxZip(name string, writer io.Writer) error {
	if _, ok := Boxes[name]; !ok {
		return ErrBoxDoesntExist
	}

	gzipwriter := gzip.NewWriter(writer)
	defer gzipwriter.Close()
	tarball := tar.NewWriter(gzipwriter)
	defer tarball.Close()

	info, err := os.Stat(Boxes[name].Location)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(Boxes[name].Location)
	}

	return filepath.Walk(Boxes[name].Location,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, Boxes[name].Location))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
}

func RemoveBox(name string) {
	delete(Boxes, name)
}
