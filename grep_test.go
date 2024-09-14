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
		"GET",
		map[string]int{
			"machine-1": 170372,
			"machine-2": 160563,
			"machine-3": 161169,
			"machine-4": 162590,
			"machine-5": 162714,
			"machine-6": 161189,
			"machine-7": 161050,
			"machine-8": 164776,
			"machine-9": 161852,
			"machine-10": 159130,
		},
	},
	{
		"DELETE",
		map[string]int{
			"machine-1": 27976,
			"machine-2": 26754,
			"machine-3": 26854,
			"machine-4": 27144,
			"machine-5": 27293,
			"machine-6": 26846,
			"machine-7": 27031,
			"machine-8": 27294,
			"machine-9": 26898,
			"machine-10": 26437,
		},
	},
	{
		"400*", // '40' and 0 or more '0'
		map[string]int{
			"machine-1": 45543,
			"machine-2": 43031,
			"machine-3": 42820,
			"machine-4": 43413,
			"machine-5": 43410,
			"machine-6": 43447,
			"machine-7": 42846,
			"machine-8": 44019,
			"machine-9": 43506,
			"machine-10": 42846,
		},
	},
}

// Executes grep for the test cases above and asserts equality of expected and received values.
func TestGrep(t *testing.T) {
	for _, grepTest := range grepTests {
		grepOutput, err := ExecuteGrepOnMachines(grepTest.queryString, MachineMap)

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
