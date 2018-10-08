package demeter

import (
	"fmt"
	"strconv"

	"github.com/goburrow/modbus"
)

func unimplemented(handler modbus.ClientHandler, args []string) {
	fmt.Println("The command is not implemented yet")
}

func disableEvent(handler modbus.ClientHandler, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: disable_event <event number>")
		return
	}

	if v, err := strconv.Atoi(args[0]); err != nil {
		fmt.Printf("%s: should be an integer\n", args[0])
		return
	} else {
		var client = modbus.NewClient(handler)
		if _, err := client.WriteSingleRegister(uint16(v*6+5), 0); err != nil {
			fmt.Println("Error disabling event")
		}
	}
}

func disableRelay(handler modbus.ClientHandler, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: disable_relay <relay number>")
		return
	}

	if v, err := strconv.Atoi(args[0]); err != nil {
		fmt.Printf("%s: should be an integer\n", args[0])
		return
	} else if v != 0 && v != 1 {
		fmt.Printf("%d: there are only two relays available\n", v)
	} else {
		var client = modbus.NewClient(handler)
		if _, err := client.WriteSingleCoil(uint16(v), 0x0000); err != nil {
			fmt.Println("Error disabling relay", err)
		}
	}
}

func enableEvent(handler modbus.ClientHandler, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: disable_event <event number>")
		return
	}

	if v, err := strconv.Atoi(args[0]); err != nil {
		fmt.Printf("%s: should be an integer\n", args[0])
		return
	} else {
		var client = modbus.NewClient(handler)
		if _, err := client.WriteSingleRegister(uint16(v*6+5), 1); err != nil {
			fmt.Println("Error disabling event")
		}
	}
}

func enableRelay(handler modbus.ClientHandler, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: enable_relay <relay number>")
		return
	}

	if v, err := strconv.Atoi(args[0]); err != nil {
		fmt.Printf("%s: should be an integer\n", args[0])
		return
	} else if v != 0 && v != 1 {
		fmt.Printf("%d: there are only two relays available\n", v)
	} else {
		var client = modbus.NewClient(handler)
		if _, err := client.WriteSingleCoil(uint16(v), 0xFF00); err != nil {
			fmt.Println("Error enabling relay", err)
		}
	}
}

func getDateTime(handler modbus.ClientHandler, args []string) {
	var client = modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(1, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		var r = expandValues(results)
		fmt.Printf("%d/%02d/%02d %02d:%02d:%02d\n", r[0], r[1], r[2], r[3], r[4], r[5])
	}
}

func getHumidity(handler modbus.ClientHandler, args []string) {
	var client = modbus.NewClient(handler)
	results, err := client.ReadInputRegisters(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		var r = expandValues(results)
		fmt.Printf("%d,%d %%\n", r[0]/10, r[0]%10)
	}
}

func getLight(handler modbus.ClientHandler, args []string) {
	var client = modbus.NewClient(handler)
	results, err := client.ReadInputRegisters(2, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		var r = expandValues(results)
		fmt.Printf("%d\n", r[0])
	}
}

func getLogInterval(handler modbus.ClientHandler, args []string) {

}

func getTemperature(handler modbus.ClientHandler, args []string) {
	var client = modbus.NewClient(handler)
	results, err := client.ReadInputRegisters(0, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		var r = expandValues(results)
		fmt.Printf("%d,%d °C\n", r[0]/10, r[0]%10)
	}
}

func readCoils(handler modbus.ClientHandler, args []string) {

}

func readDiscreteInputs(handler modbus.ClientHandler, args []string) {

}

func readEvent(handler modbus.ClientHandler, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: read_event <event number>")
		return
	}

	if v, err := strconv.Atoi(args[0]); err != nil {
		fmt.Printf("%s: should be an integer\n", args[0])
		return
	} else if v < 0 || v > 9 {
		fmt.Printf("%d: there are only 10 events available\n", v)
	} else {
		var client = modbus.NewClient(handler)
		results, err := client.ReadHoldingRegisters(uint16(v)*6+7, 6)
		if err != nil {
			fmt.Println(err)
		} else {
			printEvents(expandValues(results))
		}
	}
}

func readEvents(handler modbus.ClientHandler, args []string) {
	var client = modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(7, 6*10)
	if err != nil {
		fmt.Println(err)
	} else {
		printEvents(expandValues(results))
	}
}

func readHoldingRegisters(handler modbus.ClientHandler, args []string) {

}

func readInputRegisters(handler modbus.ClientHandler, args []string) {

}

func setDateTime(handler modbus.ClientHandler, args []string) {

}

func setLogInterval(handler modbus.ClientHandler, args []string) {

}

func writeCoils(handler modbus.ClientHandler, args []string) {

}

func writeEvent(handler modbus.ClientHandler, args []string) {

	// write_event <number> <hh> <mm> <ss> <duration> <relay> <1|0>
	if len(args) != 7 {
		fmt.Println("Usage: write_event <event number> <hh>:<mm>:<ss> <duracion> <relay> <1|0>")
		return
	}
	numbers := make([]uint16, len(args))
	for i := 0; i < len(args); i++ {
		if v, err := strconv.Atoi(args[i]); err != nil {
			fmt.Printf("%s: should be an integer\n", args[0])
			return
		} else {
			numbers[i] = uint16(v)
		}
	}

	if numbers[5] != 0 && numbers[5] != 1 {
		fmt.Printf("%d: there are only 2 relays\n", numbers[5])
	} else {
		var client = modbus.NewClient(handler)
		for i := range numbers[1:] {
			// @TODO WriteMultipleRegisters doesn't work!
			_, err := client.WriteSingleRegister(uint16(uint16(numbers[0])*6+7+uint16(i)-1), numbers[i])
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		readEvent(handler, []string{fmt.Sprintf("%d", numbers[0])})
	}

}

func writeRegisters(handler modbus.ClientHandler, args []string) {

}

func expandValues(values []byte) []int {
	newValues := make([]int, len(values)/2, len(values)/2)
	for i := range newValues {
		newValues[i] = (int(values[2*i]) << 8) + int(values[2*i+1])
	}
	return newValues
}

func printEvents(events []int) {
	for i := 0; i < len(events)/6; i++ {
		event := events[i*6 : (i+1)*6]

		state := "disabled"
		if event[5] != 0 {
			state = "enabled"
		}

		date := fmt.Sprintf("%02d:%02d:%02d", event[0], event[1], event[2])
		duration := event[3]
		relay := event[4]

		fmt.Printf("Event #%d [%s]: Begin at %s during %d seconds driven through relay #%d\n", i, state, date, duration, relay)
	}

}
