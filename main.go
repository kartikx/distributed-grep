package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print(WelcomePrompt)

	for {
		fmt.Print("Enter grep regex pattern (in single/double quotes):")

		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line = SanitizeUserInput(line)

		output, _ := ExecuteGrepOnMachines(line, MachineMap)

		PrintGatheredOutput(output)
	}
}
