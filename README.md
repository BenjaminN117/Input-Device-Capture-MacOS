# Input Device Capture MacOS

- Clipboard Capture
- Keystroke Capture

## Disclaimer
This project is for research purposes only. This project is NOT to be used maliciously and is only for educational purposes.

## Clipboard Capture
Captures the contents of a user's clipboard and then relays the information to a server where it is stored in a JSON file.

The Python server (receive_server.py) needs to be launched on the host machine waiting to recieve information from the target. The target needs to be running reader.go

Before launch, generate an API key and apply to both of the scripts and ensure the serverAddress for the reader.go points towards the receiver server.
### Dependancies
Gopsutil - github.com/shirou/gopsutil/v3/host

## Keystroke Capture

Captures the keystrokes in the background and then sends the information to the Python server.

### Dependancies
Gopsutil - github.com/shirou/gopsutil/v3/host
Gohook - github.com/robotn/gohook

## Getting Started

1) Generate an API key and add to the script and receiver_server.py
2) Ensure the reader is relaying information to the correct server address
3) Compile the script
4) Launch the receiver_server
5) Deploy the reader application on the target machine

Build the script into a single package
```
go build keystroke_reader.go
```
or
```
go build clipboard_reader.go
```


## Features to be added

- Encrypt the JSON data
- Use a Proxy to obfuscate the true location of the receiving server