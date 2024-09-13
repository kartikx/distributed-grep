package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

		executionStartTime := time.Now()

		output, _ := ExecuteGrepOnMachines(line, MachineMap)

		fmt.Printf("grep %s took %v\n", line, time.Since(executionStartTime))

		PrintGatheredOutput(output)
	}
}
