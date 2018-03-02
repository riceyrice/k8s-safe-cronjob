package main

import (
	"flag"
	"fmt"
	"errors"
	"os/exec"

	"github.com/getsentry/raven-go"
)

func main() {
	dsn := flag.String("dsn", "", "Sentry dsn to report errors")
	flag.Parse()
	command := flag.Arg(0)
	args := flag.Args()[1:]
	if *dsn != "" {
		fmt.Println("Using sentry dsn", *dsn)
		raven.SetDSN(*dsn)
	}
	if command != "" {
		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		fmt.Printf("%s", string(output))
		if err != nil {
			if *dsn != "" {
				wrapper := errors.New(fmt.Sprintf("Command exited with non-zero status: %s, reported: %s, output: %s", command, err.Error(), output))
				raven.CaptureErrorAndWait(wrapper, nil)
			} else {
				fmt.Println(err.Error())
			}
		}
	} else {
		fmt.Println("No command provided")
	}
}
