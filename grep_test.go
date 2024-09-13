package main

import (
	"testing"
)

// Models a test case for Grep.
type grepTest struct {
	queryString         string
	matchCountByMachine map[string]int
}

// TODO @sriramdvt Replace query strings and results with what's indicated for the demos.
// I.e. one frequent, one infrequent, one regex.
// List of test cases.
var grepTests = []grepTest{
	{
		"ERROR",
		map[string]int{
			"machine-1": 2,
			"machine-2": 2,
		},
	},
	{
		"DEBUG",
		map[string]int{
			"machine-1": 12,
			"machine-2": 10,
		},
	},
	{
		"NON-EXISTENT",
		map[string]int{},
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
