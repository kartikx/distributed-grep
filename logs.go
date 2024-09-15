package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Define log levels to make test log files
const (
	LogLevel   = "LOG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARNING"
	ErrorLevel = "ERROR"
	CriticalErrorLevel = "CRITICALERROR"
)

// Probability of each log level in the log files
var probabilities = map[string]float64{
	LogLevel:   0.4,  // 40% chance for Log messages
	InfoLevel:  0.3,  // 30% chance for Info messages
	WarnLevel:  0.2,  // 20% chance for Warning messages
	ErrorLevel: 0.09,  // 9% chance for Error messages
	CriticalErrorLevel: 0.01,  // 1% chance for Error messages
}

// Randomly select log level based on random probability
func getRandomLogLevel() string {
	r := rand.Float64()
	sum := 0.0
	for level, probability := range probabilities {
		sum += probability
		if r <= sum {
			return level
		}
	}
	return LogLevel // Write a default Log message if probabilities have an error
}

// Generate a random word of the specified length
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// Generate a random log message
func generateLogMessage(level string) string {
	switch level {
		case LogLevel:
			return fmt.Sprintf("This is a normal log message: %s", randSeq(100))
		case InfoLevel:
			return fmt.Sprintf("This is an informational message: %s", randSeq(100))
		case WarnLevel:
			return fmt.Sprintf("This is a warning message: %s", randSeq(100))
		case ErrorLevel:
			return fmt.Sprintf("This is an error message: %s", randSeq(100))
		case CriticalErrorLevel:
			return fmt.Sprintf("This is a critical error message: %s", randSeq(100))
		default:
			return "Unknown log level."
	}
}

func makeLogFiles(machines map[string]MachineInfo) {

	// Seed random number generator
	rand.Seed(42)

	for _, info := range machines {

		full_log := ""
		// Log messages with random levels for some iterations
		for i := 0; i < 100; i++ { // Log 100 messages
			level := getRandomLogLevel()
			message := generateLogMessage(level)
			message = fmt.Sprintf("%s : [%s] %s\n", time.Now(), level, message)
			full_log += message
		}
	
		// Make the command to populate the logfile with the constructed log message
		command_make_log := fmt.Sprintf("cat > ~/%s << EOL\n%s\nEOL", info.logFileName, full_log)
		command_make_log_split := strings.Split(command_make_log, " ")
	
		// Execute the command to populate the logfile
		output, err := ExecuteCommandOverSsh(info.user, info.address, "/home/sdevata2/.ssh/id_ecdsa", command_make_log_split)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added logs to: %s", info.address)
		fmt.Println(string(output))
	}
}
