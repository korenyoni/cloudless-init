package main

import (
	"strings"
	"os"
	"os/exec"
	"testing"
	"syscall"
)

func TestHost(t *testing.T) {
	oldHostname, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	var hostname string = "test"
	cmd := exec.Command("go", "run", "main.go", "test-Host(hostname)", hostname)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER |
                        syscall.CLONE_NEWUTS,
		UidMappings: []syscall.SysProcIDMap{
                        {
                                ContainerID: 0,
                                HostID:      os.Getuid(),
                                Size:        1,
                        },
                },
                GidMappings: []syscall.SysProcIDMap{
                        {
                                ContainerID: 0,
                                HostID:      os.Getgid(),
                                Size:        1,
                        },
                },
	}
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	outputString := strings.Trim(string(output), "\n")
	if  outputString != hostname {
		t.Fatalf("Expected hostname to be '%s' after Host(%s), got '%s'", hostname, hostname, output)
	}
	t.Logf("Hostname was '%s', successfuly changed it to '%s' in new UTS namespace", oldHostname, outputString)
}
