package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/getsentry/raven-go"
)

func main() {
	dsn := flag.String("dsn", "", "Sentry dsn to report errors")
	flag.Parse()
	command := strings.Join(flag.Args(), " ")
	if *dsn != "" {
		fmt.Println("Using sentry dsn", *dsn)
		raven.SetDSN(*dsn)
	}
	if command != "" {
		cmd := exec.Command(command)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			if *dsn != "" {
				raven.CaptureErrorAndWait(err, nil)
			} else {
				fmt.Println(err.Error())
			}
		}
	} else {
		fmt.Println("No command provided")
	}
}
