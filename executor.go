package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Executes grep for the given user input across machines.
// The user input is expected as an array of strings. For example, ["grep", "-i", "query"]
// The output is returns as a map of file names and grep output.
func ExecuteGrepOnMachines(input []string, machines map[string]MachineInfo) (map[string]string, error) {
	finalOutput := make(map[string]string)

	outChan := make(chan GrepOutput)

	for name, info := range machines {
		go executeGrepOnMachine(input, name, info, outChan)
	}

	// TODO Check whether this is more optimal than a buffered channel.
	// TODO Either perform self ssh, or skip iterating over sself.
	for i := 0; i < NumMachines; i++ {
		// fmt.Println("Requesting data")
		select {
		case val := <-outChan:
			finalOutput[val.machineName] = val.output
		// TODO this is just for safety. How do I select an optimal number?
		case <-time.After(60 * time.Second):
			return finalOutput, nil
		}
	}

	return finalOutput, nil
}

// Executes grep for the given user input on a single machine.
// The user input is expected as an array of strings. For example, ["grep", "-i", "query"]
// The output for the operation is written on the channel.
func executeGrepOnMachine(input []string, machineName string, machineInfo MachineInfo,
	outChan chan GrepOutput) {
	// fmt.Println("Executing grep: ", machineName)

	grepOutput, err := executeGrepOverSsh(input, machineInfo)

	if err != nil {
		// fmt.Println("Executed grep: ", machineName)
		outChan <- GrepOutput{machineName, ""}
		return
	}

	// TODO Eventually remove sleep. Current present for testing concurrency.
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)

	fmt.Println("Executed grep on ", machineName)
	outChan <- GrepOutput{machineName, grepOutput}
}

// TODO Use pointers. How are arrays passed?
func executeGrepOverSsh(input []string, machineInfo MachineInfo) (string, error) {

	// TODO replace with a function that invokes ssh.
	// TODO this will fail for commands such as grep -i "error failed", due to the double quotes
	// grepCommand := exec.Command("grep", input[1:]...)

	out, err := executeCommandOverSsh(machineInfo.user, machineInfo.address, "/home/sriramdvt/.ssh/id_ecdsa", input)
	fmt.Println(string(out))

	if err != nil {
		// This could happen with invalid file (status 2), or if no results (status 1).
		// TODO check status code.
		return "", err
	}

	// TODO could this fail? Add nil check.
	return string(out), nil
}
