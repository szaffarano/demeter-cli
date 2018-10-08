package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	prompt "github.com/c-bata/go-prompt"
	"github.com/goburrow/modbus"
	"github.com/spf13/cobra"
	"github.com/szaffarano/demeter-cli/demeter"
)

type config struct {
	method   string
	port     string
	stopbits int
	bytesize int
	parity   string
	baudrate int
	timeout  int
}

var cfg = config{}
var handler *modbus.RTUClientHandler
var verbose bool

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "demeter-cli",
	Short: "Command Line Interface to manage demeter irrigatin system",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("demeter-cli %s (rev-%s)\n", demeter.Version, demeter.Revision)
		fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
		defer fmt.Println("Bye!")

		handler = modbus.NewRTUClientHandler(cfg.port)

		handler.BaudRate = cfg.baudrate
		handler.Parity = cfg.parity
		handler.StopBits = cfg.stopbits
		handler.DataBits = cfg.bytesize
		handler.SlaveId = 0x03
		handler.Timeout = time.Duration(cfg.timeout) * time.Second
		if verbose {
			handler.Logger = log.New(os.Stdout, "demeter-cli: ", log.LstdFlags)
		}

		err := handler.Connect()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		defer handler.Close()

		p := prompt.New(
			demeter.Executor(handler),
			demeter.Completer,
			prompt.OptionTitle("demeter-cli: cli for demeter irrigation system"),
			prompt.OptionPrefix("> "),
			prompt.OptionInputTextColor(prompt.Yellow),
		)

		p.Run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.method, "method", "m", "rtu", "Connection method")
	rootCmd.PersistentFlags().StringVarP(&cfg.port, "device", "d", "/dev/rfcomm0", "Path to the bluetooth device to use")
	rootCmd.PersistentFlags().IntVarP(&cfg.stopbits, "stop", "s", 2, "Stop bits")
	rootCmd.PersistentFlags().IntVarP(&cfg.bytesize, "size", "S", 8, "Byte Size")
	rootCmd.PersistentFlags().StringVarP(&cfg.parity, "paritty", "p", "N", "Should use parity?")
	rootCmd.PersistentFlags().IntVarP(&cfg.baudrate, "bauds", "b", 9600, "Baud rate")
	rootCmd.PersistentFlags().IntVarP(&cfg.timeout, "timeout", "t", 1, "Timeout")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose")
}
