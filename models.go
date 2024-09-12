package main

// Models the information about a Machine.
type MachineInfo struct {
	machineNum	int
	user		string
	address     string
	logFileName string
}

// Models the output of grep from a Machine.
type GrepOutput struct {
	machineName string
	output      string
}
