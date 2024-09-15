package main

import (
	"testing"
)

// Models a test case for Grep.
type grepTest struct {
	queryString         string
	matchCountByMachine map[string]int
}

// List of test cases.
var grepTests = []grepTest{
	{
		"LOG",
		map[string]int{
			"machine-1" : 38,
			"machine-2" : 45,
			"machine-3" : 43,
			"machine-4" : 36,
			"machine-5" : 43,
			"machine-6" : 41,
			"machine-7" : 36,
			"machine-8" : 30,
			"machine-9" : 46,
			"machine-10": 43,
		},
	},
	{
		"WARNING",
		map[string]int{
			"machine-1": 18,
			"machine-2": 22,
			"machine-3": 21,
			"machine-4": 22,
			"machine-5": 18,
			"machine-6": 21,
			"machine-7": 15,
			"machine-8": 24,
			"machine-9": 17,
			"machine-10": 12,
		},
	},
	{
		"CERROR", // does not exist
		map[string]int{},
	},
}

// Executes grep for the test cases above and asserts equality of expected and received values.
func TestGrep(t *testing.T) {

	// makeLogFiles(TestMachineMap)

	for _, grepTest := range grepTests {
		grepOutput, err := ExecuteGrepOnMachines(grepTest.queryString, TestMachineMap)

		if grepOutput == nil || err != nil {
			t.Errorf("Unexpected error: %s with grep output: %s", err, grepOutput)
		}

		if len(grepTest.matchCountByMachine) != len(grepOutput) {
			t.Errorf("Expected number of matched files: %d, Received: %d", len(grepTest.matchCountByMachine), len(grepOutput))
		}

		for machineName, expectedCount := range grepTest.matchCountByMachine {
			if grepOutput[machineName] == nil || len(grepOutput[machineName]) != expectedCount {
				t.Errorf("Machine: %s Expected: %d, Received: %d", machineName, expectedCount, len(grepOutput[machineName]))
			}
		}
	}
}
