# XGO-API

## Overview

XGO-Api is a client implemention for The Experience API (or xAPI) written in Golang.
## The Experience API (xAPI)

The Experience API (or xAPI) is a new specification for learning technology that makes it possible to collect data about the wide range of experiences a person has (online and offline). This API captures data in a consistent format about a person or group’s activities from many technologies. Very different systems are able to securely communicate by capturing and sharing this stream of activities using xAPI’s simple vocabulary.

## Installation
	go get github.com/bkaraceylan/xapi-go@latest
## Module Usage
	lrs, err := client.NewRemoteLRS(
		"https://cloud.scorm.com/ScormEngineInterface/TCAPI/public/",
		"1.0.0",
		"Basic VGVzdFVzZXI6cGFzc3dvcmQ=",
	)

	if err != nil {
		panic(err)
	}

	statement, err := lrs.GetStatement("b1893eed-14e6-4ac2-b154-3c6e828c2297")

	if err != nil {
		panic(err)
	}

	pretty, err := statement.ToJson(true)

	if err != nil {
		panic(err)
	}

	fmt.Println(pretty)

## CLI Usage
	Usage:
	xapi-go [flags]
	xapi-go [command]

	Available Commands:
	completion   Generate the autocompletion script for the specified shell
	getStatement 
	help         Help about any command

	Flags:
		--auth string       Authentication header (Basic, Bearer etc...)
		--endpoint string   URL of the API endpoint
	-h, --help              help for xapi-go
		--password string   API user's password
		--username string   API user's username
		--version string    API version

	Use "xapi-go [command] --help" for more information about a command.

## TODO
### Module
- [x] About Resource
- [*] State Resource
- [ ] Documents Resource
- [ ] Agents Resource
- [ ] Activities Resource
- [ ] Agent Profile Resource
- [ ] Activity Profile Resource
### Module
- [x] About Resource
- [*] State Resource
- [ ] Documents Resource
- [ ] Agents Resource
- [ ] Activities Resource
- [ ] Agent Profile Resource
- [ ] Activity Profile Resource
