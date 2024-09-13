package main

const userNetId = "kartikr2"
const sshKeyPath = "/Users/kartik/.ssh/id_ecdsa"

// Maps the machine identifiers to machine information.
var MachineMap = map[string]MachineInfo{
	"machine-1":  {1, userNetId, "fa24-cs425-6401.cs.illinois.edu", "machine.1.log"},
	"machine-2":  {2, userNetId, "fa24-cs425-6402.cs.illinois.edu", "machine.2.log"},
	"machine-3":  {3, userNetId, "fa24-cs425-6403.cs.illinois.edu", "machine.3.log"},
	"machine-4":  {4, userNetId, "fa24-cs425-6404.cs.illinois.edu", "machine.4.log"},
	"machine-5":  {5, userNetId, "fa24-cs425-6405.cs.illinois.edu", "machine.5.log"},
	"machine-6":  {6, userNetId, "fa24-cs425-6406.cs.illinois.edu", "machine.6.log"},
	"machine-7":  {7, userNetId, "fa24-cs425-6407.cs.illinois.edu", "machine.7.log"},
	"machine-8":  {8, userNetId, "fa24-cs425-6408.cs.illinois.edu", "machine.8.log"},
	"machine-9":  {9, userNetId, "fa24-cs425-6409.cs.illinois.edu", "machine.9.log"},
	"machine-10": {10, userNetId, "fa24-cs425-6410.cs.illinois.edu", "machine.10.log"},
}

// The path to the folder containing the log files.
const LogFolderPath = "~/"
