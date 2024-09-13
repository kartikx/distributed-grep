package main

import (
	"fmt"
	"os"
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

// Prints the grep result to the console.
// Stores individual VM results on the client machine.
func PrintGatheredOutput(gatheredOutput map[string][]string) {
	lineCount := 0
	fmt.Println("===OUTPUT===")

	for machine, output := range gatheredOutput {
		if len(output) == 0 {
			continue
		}

		fmt.Printf("%s: [%d lines]\n", machine, len(output))

		writeOutputToFile(machine, output)

		lineCount += len(output)
	}

	fmt.Printf("Total number of matched lines: %d\n\n", lineCount)

	if lineCount > 0 {
		fmt.Printf("\nLine matches were written to named files on this machine.\n\n")
	}
}

func SanitizeUserInput(input string) string {
	return strings.TrimSpace(input)
}

// Writes the output of the queried machine to a file on the client machine for easy querying.
func writeOutputToFile(machine string, output []string) {
	filename := machine + ".out"
	file, err := os.Create(filename)

	if err != nil {
		fmt.Printf("Unable to create file %s: %v\n", filename, err)
		return
	}

	content := strings.Join(output, "\n")
	content += "\n"

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Printf("Unable to write output to file %s: %v\n", filename, err)
		return
	}
}
