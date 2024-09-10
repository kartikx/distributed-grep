package main

// Maps the machine identifiers to machine information.
var MachineMap = map[string]MachineInfo{
	"machine-1": {"a@server", "machine.1.log"},
	"machine-2": {"b@server", "machine.2.log"},
	"machine-3": {"b@server", "machine.3.log"},
	"machine-4": {"b@server", "machine.4.log"},
	"machine-5": {"b@server", "machine.5.log"},
}

// The number of machines in the network.
const NumMachines int = 5

// The path to the folder containing the log files.
const LogFolderPath = "/Users/kartik/Code/CS425/MP1/Logs"
