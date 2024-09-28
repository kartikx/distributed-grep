package main

const userNetId = "kartikr2"
const sshKeyPath = "/home/kartikr2/.ssh/id_ecdsa"

// Maps the machine identifiers to machine information.
var MachineMap = map[string]MachineInfo{
	"machine-1":  {1, userNetId, "fa24-cs425-6401.cs.illinois.edu", "machine.log"},
	"machine-2":  {2, userNetId, "fa24-cs425-6402.cs.illinois.edu", "machine.log"},
	"machine-3":  {3, userNetId, "fa24-cs425-6403.cs.illinois.edu", "machine.log"},
	"machine-4":  {4, userNetId, "fa24-cs425-6404.cs.illinois.edu", "machine.log"},
	"machine-5":  {5, userNetId, "fa24-cs425-6405.cs.illinois.edu", "machine.log"},
	"machine-6":  {6, userNetId, "fa24-cs425-6406.cs.illinois.edu", "machine.log"},
	"machine-7":  {7, userNetId, "fa24-cs425-6407.cs.illinois.edu", "machine.log"},
	"machine-8":  {8, userNetId, "fa24-cs425-6408.cs.illinois.edu", "machine.log"},
	"machine-9":  {9, userNetId, "fa24-cs425-6409.cs.illinois.edu", "machine.log"},
	"machine-10": {10, userNetId, "fa24-cs425-6410.cs.illinois.edu", "machine.log"},
}

// Maps the machine identifiers to machine information.
var TestMachineMap = map[string]MachineInfo{
	"machine-1":  {1, userNetId, "fa24-cs425-6401.cs.illinois.edu", "machine.1.log.test"},
	"machine-2":  {2, userNetId, "fa24-cs425-6402.cs.illinois.edu", "machine.2.log.test"},
	"machine-3":  {3, userNetId, "fa24-cs425-6403.cs.illinois.edu", "machine.3.log.test"},
	"machine-4":  {4, userNetId, "fa24-cs425-6404.cs.illinois.edu", "machine.4.log.test"},
	"machine-5":  {5, userNetId, "fa24-cs425-6405.cs.illinois.edu", "machine.5.log.test"},
	"machine-6":  {6, userNetId, "fa24-cs425-6406.cs.illinois.edu", "machine.6.log.test"},
	"machine-7":  {7, userNetId, "fa24-cs425-6407.cs.illinois.edu", "machine.7.log.test"},
	"machine-8":  {8, userNetId, "fa24-cs425-6408.cs.illinois.edu", "machine.8.log.test"},
	"machine-9":  {9, userNetId, "fa24-cs425-6409.cs.illinois.edu", "machine.9.log.test"},
	"machine-10": {10, userNetId, "fa24-cs425-6410.cs.illinois.edu", "machine.10.log.test"},
}

// The path to the folder containing the log files.
const LogFolderPath = "~/mp2-g64"
