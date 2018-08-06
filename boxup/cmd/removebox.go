package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func generateRemoveBoxURL(host, port, name string) string {
	return fmt.Sprintf("http://%v:%v/RemoveBox/%v", host, port, name)
}

func removeBox(host, port, name string) error {
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest(http.MethodDelete, generateRemoveBoxURL(host, port, name), nil)
	if err != nil {
		return fmt.Errorf("Error occured creating request: %v", err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error occured calling Server: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode != 204 {
		return fmt.Errorf("Error occured calling Server: Status code returned was not 204: %v (%v) Message: '%v'", resp.StatusCode, resp.Status, bodyString)
	}
	return nil
}
