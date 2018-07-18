package main

import (
	"os"
	"fmt"
	"log"
	"syscall"
)

func Host(hostname string) (string, error) {
	err := syscall.Sethostname([]byte(hostname))
	if err != nil {
		return "", err
	}
	hostname, err = os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func main() {
	args := os.Args[1:]
	if len(args) == 2 && args[0] == "test-Host(hostname)"  {
		hostname, err := Host(args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hostname)
	} else {
		fmt.Println("This cli interface is dedicated being executed by unit tests")
	}
}
