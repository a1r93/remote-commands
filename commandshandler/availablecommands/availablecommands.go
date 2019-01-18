package availablecommands

type macCommand struct {
	id      string
	command string
}

var possibleCommands = []macCommand{
	macCommand{
		id:      "utc-time",
		command: "date -u",
	},
	macCommand{
		id:      "cpu-usage",
		command: `ps -A -o %cpu | awk '{s+=$1} END {print s "%"}'`,
	},
	macCommand{
		id:      "available-ram",
		command: `ps -A -o pmem | awk '{s+=$1} END {print s "%"}'`,
	},
	macCommand{
		id:      "say-something",
		command: `say "I just said something mate"`,
	},
	macCommand{
		id:      "capture",
		command: "screencapture /Users/arnaud/test.jpg",
	},
}

// findCommand looks in the possible commandIDs and returns the
// according command string if it has been found
func findCommand(commandID string) (macCommand, bool) {
	for _, command := range possibleCommands {
		if commandID == command.id {
			return command, true
		}
	}
	var comm macCommand
	return comm, false
}

// commandExists checks whether a given command id is supported
func commandExists(commandID string) bool {
	_, ok := findCommand(commandID)
	return ok
}

// CheckCommandsList takes an array of strings as parameter and checks whether
// the given commands are supported
func CheckCommandsList(commands []string) bool {
	for _, command := range commands {
		if !commandExists(command) {
			return false
		}
	}
	return true
}

// GetCommand takes a command id as parameter and returns the according command
func GetCommand(commandID string) string {
	command, _ := findCommand(commandID)
	return command.command
}
