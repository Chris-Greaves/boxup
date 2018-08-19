package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func getBox(host, port, output, name string) error {
	fmt.Println("Getting box from server")
	resp, err := http.Get(generateGetBoxURL(host, port, name))
	if err != nil {
		return fmt.Errorf("Error occured contacting server: %v", err)
	} else if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("Error occured contacting server: Status code returned was not 200: %v (%v) Message: '%v'", resp.StatusCode, resp.Status, bodyString)
	}
	defer resp.Body.Close()

	fmt.Println("Uncompressing stream")
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("Error occured creating gzip reader: %v", err)
	}
	defer gzipReader.Close()

	if output == "" {
		output, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("Error getting working directory: %v", err)
		}
	}

	fmt.Println("Unzipping stream")
	err = untar(gzipReader, output)
	if err != nil {
		return fmt.Errorf("Error occured unzipping: %v", err)
	}
	return nil
}

func generateGetBoxURL(host, port, name string) string {
	return fmt.Sprintf("http://%v:%v/GetBox/%v", host, port, name)
}

func untar(reader io.Reader, target string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Error reading next: %v", err)
		}

		path := filepath.Join(target, header.Name)
		fmt.Printf("Creating %v\n", path)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return fmt.Errorf("Error making directory: %v", err)
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return fmt.Errorf("Error opening file: %v", err)
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return fmt.Errorf("Error copying file: %v", err)
		}
	}

	return nil
}
