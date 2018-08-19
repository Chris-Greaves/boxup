package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func generateListBoxesURL(host, port string) string {
	return fmt.Sprintf("http://%v:%v/Boxes", host, port)
}

func listBoxes(host, port string) error {
	resp, err := http.Get(generateListBoxesURL(host, port))
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
