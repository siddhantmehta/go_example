package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("Error: Can not open pid file %s", err)
	}
	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return fmt.Errorf("Error: bad process ID, %s", err)
	}
	fmt.Printf("Killing server with pid=%d\n", pid)
	return nil
}
func main() {
	if err := killServer("abcd.txt"); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
