# Input Device Capture MacOS

![Static Badge](https://img.shields.io/badge/Python-=>3.11.4-blue)
![Static Badge](https://img.shields.io/badge/Golang-=>1.19.4-blue)
![Static Badge](https://img.shields.io/badge/License-MIT-yellow)


## Description

This is a collection of Golang scripts that capture a variety of different inputs through MacOS.

### Clipboard Capture
Captures the contents of a user's clipboard and then relays the information to a server where it is stored in a JSON file.

The Python server (receive_server.py) needs to be launched on the host machine waiting to receive information from the target. The target needs to be running reader.go

Before launch, generate an API key and apply to both of the scripts and ensure the serverAddress for the reader.go points towards the receiver server.

### Keystroke Capture

Captures the keystrokes in the background and then sends the information to the Python server.

The python server needs to be launched before the keystroke script has begun.

### Features to be added

- Encrypt the JSON data
- Use a Proxy to obfuscate the true location of the receiving server


## Disclaimer

This project is for research purposes only. This project is NOT to be used maliciously and is only for educational purposes.

## Prerequisites

- Python >= 3.11.4
- Golang >= 1.19.4

## Dependancies

- Python
    - Flask = 2.3.2
- Golang
    - github.com/go-ole/go-ole v1.2.6
	- github.com/lufia/plan9stats v0.0.0
	- github.com/power-devops/perfstat v0.0.0
	- github.com/robotn/gohook v0.40.0
	- github.com/shirou/gopsutil/v3 v3.23.7
	- github.com/shoenig/go-m1cpu v0.1.6
	- github.com/tklauser/go-sysconf v0.3.11
	- github.com/tklauser/numcpus v0.6.0
	- github.com/vcaesar/keycode v0.10.0
	- github.com/yusufpapurcu/wmi v1.2.3
	- golang.org/x/sys v0.10.0

## Install

```
git clone https://github.com/BenjaminN117/Input-Device-Capture-MacOS.git
```

### Server

```
python3 -m pip install -r requirements.txt
```

- Modify src/server/env.py.

### Capture Tools

- Include the target server in each script before compiling.

```
make build clipboard
```
```
make build keystroke
```


## Usage

Python Server

```
python3 src/server/receive_server.py
```

Golang Scripts

```
./bin/clipbaord_reader
```
```
./bin/keystroke_reader
```