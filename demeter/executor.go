package demeter

import (
	"fmt"
	"os"
	"strings"

	"github.com/goburrow/modbus"
)

type command struct {
	Name        string
	Description string
	Code        code
}
type code func(modbus.ClientHandler, []string)

var commands = map[string]command{
	"read_coils": {
		Name:        "read_coils",
		Description: "Read the values of all coils",
		Code:        readCoils,
	},
	"read_discrete_inputs": {
		Name:        "read_discrete_inputs",
		Description: "Read the values of all discrete inputs",
		Code:        readDiscreteInputs,
	},
	"read_holding_registers": {
		Name:        "read_holding_registers",
		Description: "Read the values of the holding registers",
		Code:        readHoldingRegisters,
	},
	"read_input_registers": {
		Name:        "read_input_registers",
		Description: "Read the input registers",
		Code:        readInputRegisters,
	},
	"write_coils": {
		Name:        "write_coils",
		Description: "Write coils",
		Code:        writeCoils,
	},
	"write_registers": {
		Name:        "write_registers",
		Description: "Write registers",
		Code:        writeRegisters,
	},
	"get_datetime": {
		Name:        "get_datetime",
		Description: "Get demeter's date and time",
		Code:        getDateTime,
	},
	"set_datetime": {
		Name:        "set_datetime",
		Description: "Set demeter's date and time",
		Code:        setDateTime,
	},
	"get_loginterval": {
		Name:        "get_loginterval",
		Description: "Get the log interval",
		Code:        getLogInterval,
	},
	"set_loginterval": {
		Name:        "set_loginterval",
		Description: "Set the log interval",
		Code:        setLogInterval,
	},
	"read_event": {
		Name:        "read_event",
		Description: "Read an event",
		Code:        readEvent,
	},
	"write_event": {
		Name:        "write_event",
		Description: "Write an event",
		Code:        writeEvent,
	},
	"read_events": {
		Name:        "read_events",
		Description: "Read all events",
		Code:        readEvents,
	},
	"disable_event": {
		Name:        "disable_event",
		Description: "Disable an event",
		Code:        disableEvent,
	},
	"enable_event": {
		Name:        "enable_event",
		Description: "Enable an event",
		Code:        enableEvent,
	},
	"disable_relay": {
		Name:        "disable_relay",
		Description: "Disable a relay",
		Code:        disableRelay,
	},
	"enable_relay": {
		Name:        "enable_relay",
		Description: "Enable a relay",
		Code:        enableRelay,
	},
	"temperature": {
		Name:        "temperature",
		Description: "Get demeter's temperature",
		Code:        getTemperature,
	},
	"humidity": {
		Name:        "humidity",
		Description: "Get demeter's humidity",
		Code:        getHumidity,
	},
	"light": {
		Name:        "light",
		Description: "Get demete's light value",
		Code:        getLight,
	},
}

// Executor run any valid command
func Executor(handler modbus.ClientHandler) func(string) {
	return func(cmd string) {
		fullCmd := strings.Split(strings.ToLower(strings.TrimSpace(cmd)), " ")
		cmdName := fullCmd[0]
		cmdArgs := []string{}
		if len(fullCmd) > 0 {
			cmdArgs = fullCmd[1:]
		}

		if len(cmdName) == 0 {
			return
		} else if cmdName == "quit" || cmdName == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
			return
		}
		if command, ok := commands[cmdName]; ok {
			command.Code(handler, cmdArgs)
		} else {
			fmt.Println(cmdName, "Unkwnown command")
		}
	}

}
