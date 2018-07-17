package main

import (
	"fmt"
	"syscall"
)

func Host(hostname string) error {
	fmt.Printf("Setting hostname to: %s\n", hostname)
	err := syscall.Sethostname([]byte(hostname))
	if err != nil {
		return err
	}
	return nil
}
