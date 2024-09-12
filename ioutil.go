package main

import (
	"fmt"
	"strings"
)

// Welcome prompt that appears on program startup.
const WelcomePrompt = `
   _____ ______      ________ ______ ______ ______   _____ _   _  _____ 
  / ____/ __ \ \    / /  ____|  ____|  ____|  ____| |_   _| \ | |/ ____|
 | |   | |  | \ \  / /| |__  | |__  | |__  | |__      | | |  \| | |     
 | |   | |  | |\ \/ / |  __| |  __| |  __| |  __|     | | | .   | |     
 | |___| |__| | \  /  | |    | |____| |    | |____   _| |_| |\  | |____ 
  \_____\____/   \/   |_|    |______|_|    |______| |_____|_| \_|\_____|                                                                        
`

// Prints the aggregated grep result to the console.
func PrintGatheredOutput(gatheredOutput map[string]string) {
	fmt.Println("==== OUTPUT ================")
	for machine, output := range gatheredOutput {
		if len(output) == 0 {
			continue
		}

		fmt.Println("==== " + machine + " ================")
		numLinesInGrepOutput := strings.Count(output, "\n")
		fmt.Printf("%d lines\n", numLinesInGrepOutput)
	}
	fmt.Println("====")
}

func ParseUserInput(input string) []string {
	// Remove trailing new line.
	input = input[:len(input)-1]

	return strings.Split(input, " ")
}

// Validate that the first token is "grep"
// Handle quotes if required
// Optionally, handle injections.
func ValidateUserInput(input []string) error {
	// TODO implement.

	return nil
}
