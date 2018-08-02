package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func generateAddBoxURL(host, port string) string {
	return fmt.Sprintf("http://%v:%v/CreateBox", host, port)
}

func addBox(host, port, name, path string) error {
	path = strings.Replace(path, "\\", "\\\\", -1)
	jsonStr := []byte(`{"name": "` + name + `","location": "` + path + `"}`)

	resp, err := http.Post(generateAddBoxURL(host, port), "json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return fmt.Errorf("Error occured calling Server: %v", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error occured calling Server: Status code returned was not 200: %v (%v) Message: '%v'", resp.StatusCode, resp.Status, bodyString)
	}
	fmt.Printf("Response: %v", bodyString)
	return nil
}
