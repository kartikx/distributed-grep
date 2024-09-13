package main

import (
	"fmt"
	"path"
	"strings"
	"time"
)

// Executes grep for the given user input across machines.
// The output is a map of file names and newline separated grep output.
func ExecuteGrepOnMachines(input string, machines map[string]MachineInfo) (map[string][]string, error) {
	finalOutput := make(map[string][]string)

	outChan := make(chan GrepOutput)

	for name, info := range machines {
		go executeGrepOnMachine(input, name, info, outChan)
	}

	for i := 0; i < len(machines); i++ {
		select {
		case val := <-outChan:
			val.output = strings.TrimSpace(val.output)
			if len(val.output) != 0 {
				finalOutput[val.machineName] = strings.Split(val.output, "\n")
			}
		case <-time.After(20 * time.Second):
			fmt.Println("Timeout. Didn't hear back from expected number of machines.")
			return finalOutput, nil
		}
	}

	return finalOutput, nil
}

// Executes grep for the given user input on a single machine.
// The output for the operation is written on the channel.
func executeGrepOnMachine(input string, machineName string, machineInfo MachineInfo,
	outChan chan GrepOutput) {
	grepOutput, err := executeGrepOverSsh(input, machineInfo)

	if err != nil {
		outChan <- GrepOutput{machineName, ""}
		return
	}

	outChan <- GrepOutput{machineName, grepOutput}
}

// Executes the grep command for a machine over ssh.
func executeGrepOverSsh(input string, machineInfo MachineInfo) (string, error) {
	logFilePath := path.Join(LogFolderPath, machineInfo.logFileName)
	grepCommandToRun := []string{"grep", "-E", input, logFilePath}

	out, err := ExecuteCommandOverSsh(machineInfo.user, machineInfo.address, sshKeyPath, grepCommandToRun)

	if err != nil {
		// This could happen if there are no results for grep (status code 1),
		// or if the grep command failed (status code 2)
		return "", err
	}

	return out, nil
}
