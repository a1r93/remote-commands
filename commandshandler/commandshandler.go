package commandshandler

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"remotecmds/commandshandler/availablecommands"
	"remotecmds/commandshandler/httpstatus"
)

type resultObj struct {
	output  string
	command string
}

func (r resultObj) toString() string {
	return r.command + ": " + r.output
}

// handleSingleCommand executes a given command and returns the result to the
// result channel
func handleSingleCommand(result chan resultObj, command string, commandID string) {
	output, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		toReturn := resultObj{
			command: commandID,
			output:  err.Error(),
		}
		result <- toReturn
		return
	}

	var outputString string

	if len(output) == 0 {
		outputString = "Done"
	} else {
		outputString = string(output)
	}

	toReturn := resultObj{
		command: commandID,
		output:  outputString,
	}
	result <- toReturn
}

// Handler function used in order to handle the /commands requests
func Handler(w http.ResponseWriter, r *http.Request) {
	commands, ok := r.URL.Query()["cmds"]

	// Validate the query parameter
	if ok == false {
		httpstatus.SendBadrequest(w, "The cmds parameter is missing")
		return
	}

	valid := availablecommands.CheckCommandsList(commands)
	if valid == false {
		httpstatus.SendBadrequest(w, "One of the given commands does not exist")
		return
	}

	resultChannel := make(chan resultObj)
	length := len(commands)

	// Run a goroutine for every command
	for _, commandID := range commands {
		command := availablecommands.GetCommand(commandID)

		go handleSingleCommand(resultChannel, command, commandID)
	}

	// We know that commands.length goRoutines were created, so we expect result
	// to resolve commands.length times.
	var toReturn []string
	for i := 0; i < length; i++ {
		result, _ := <-resultChannel

		toReturn = append(toReturn, result.toString())
	}

	// Close the result channel and respond to the calling client
	close(resultChannel)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toReturn)
}
