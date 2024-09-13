package main

import (
	"bytes"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Executes a command over ssh. The command should be passed as an array of strings.
// privateKeyLocation must point to a valid SSH key which is whitelisted on the user and address being accessed.
// This code  was inspired from: https://stackoverflow.com/a/41528181
func ExecuteCommandOverSsh(user string, addr string, privateKeyLocation string, cmd []string) (string, error) {
	pvtKeyBts, err := os.ReadFile(privateKeyLocation)

	if err != nil {
		return "", err
	}

	key, err := ssh.ParsePrivateKey(pvtKeyBts)
	if err != nil {
		return "", err
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
	session.Stdout = &cmdOutputBuffer

	err = session.Run(strings.Join(cmd, " "))

	return cmdOutputBuffer.String(), err
}
