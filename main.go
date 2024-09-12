package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(WelcomePrompt)

	// makeLogFiles(MachineMap)

	for {
		fmt.Print("Enter grep regex pattern (in single/double quotes):")

		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		parsedInput := ParseUserInput(line)

		// TODO Add validations for user input. Use function from IOInput.

		output, _ := ExecuteGrepOnMachines(parsedInput, MachineMap)

		PrintGatheredOutput(output)
	}
}
