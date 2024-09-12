package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"net"
	"log"
	"bytes"
	"strings"
	"golang.org/x/crypto/ssh"
)

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

// Generate a random log message
func generateLogMessage(level string) string {
	switch level {
		case LogLevel:
			return "This is a normal log message."
		case InfoLevel:
			return "This is an informational message."
		case WarnLevel:
			return "This is a warning message."
		case ErrorLevel:
			return "This is an error message."
		case CriticalErrorLevel:
			return "This is a critical error message."
		default:
			return "Unknown log level."
	}
}

func executeCommandOverSsh(user string, addr string, privateKeyLocation string, cmd []string) (string, error) {
    
	// read privateKey from a file
    pvtKeyBts, err := os.ReadFile(privateKeyLocation)
    key, err := ssh.ParsePrivateKey(pvtKeyBts)
    if err != nil {
        return "ParsePrivateKey error", err
    }

    // SSH Authentication
    config := &ssh.ClientConfig{
        User: user,
        // https://github.com/golang/go/issues/19767 
        // as clientConfig is non-permissive by default 
        // you can set ssh.InsercureIgnoreHostKey to allow any host 
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(key),
        },
    }

    // Setup an SSH connection
    client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
    if err != nil {
        return "", err
    }

    // Create a session over the SSH connection.
    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    var cmdOutputBuffer bytes.Buffer
    session.Stdout = &cmdOutputBuffer // get output of running command to the buffer

    // Run the command
    err = session.Run(strings.Join(cmd, " "))

    return cmdOutputBuffer.String(), err
}

func makeLogFiles(machines map[string]MachineInfo) {
	// Seed random number generator
	rand.Seed(42)

	for _, info := range machines {

		full_log := ""
		// Log messages with random levels for some iterations
		for i := 0; i < 100; i++ { // Adjust 100 to control how many messages you want to log
			level := getRandomLogLevel()
			message := generateLogMessage(level)
			message = fmt.Sprintf("%s : [%s] %s\n", time.Now(), level, message)
			full_log += message
		}
	
		// Make the command to populate the logfile with the constructed log message
		command_make_log := fmt.Sprintf("cat > ~/%s << EOL\n%s\nEOL", info.logFileName, full_log)
		command_make_log_split := strings.Split(command_make_log, " ")
	
		// Execute the command to populate the logfile
		output, err := executeCommandOverSsh(info.user, info.address, "/home/sriramdvt/.ssh/id_ecdsa", command_make_log_split)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added logs to: %s", info.address)
		fmt.Println(string(output))

	}



}

