package main

import (
	"net/http"
	"remotecmds/commandshandler"
)

var port = ":8080"

func main() {
	http.HandleFunc("/commands", commandshandler.Handler)
	http.ListenAndServe(port, nil)
}
