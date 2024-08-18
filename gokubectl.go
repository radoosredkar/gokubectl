package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var podName string
	flag.StringVar(&podName, "pn", "", "partial of full podname in the format of -pn=nginx")
	var ports string
	flag.StringVar(&ports, "p", "", "ports in the format of -p=toPort:fromPort")
	flag.Parse()

	if podName == "" {
		fmt.Println("Please specify the pod name as pn=<podname>")
		return
	}
	if !isPortParamValid(ports) {
		return
	}

	// Create the kubectl command
	kubectlCmd := exec.Command("kubectl", "get", "pods")

	// Create a pipe for the grep command
	grepCmd := exec.Command("grep", podName)

	var kubectlOut bytes.Buffer
	var kubectlErr bytes.Buffer
	var grepOut bytes.Buffer
	var grepErr bytes.Buffer

	kubectlCmd.Stdout = &kubectlOut
	kubectlCmd.Stderr = &kubectlErr

	err := kubectlCmd.Run()
	if err != nil {
		fmt.Println("Error executing kubectl command:", err)
		fmt.Println("Kubectl Stderr:", kubectlErr.String())
		return
	}

	// Set the input for the grep command to the output of the kubectl command
	grepCmd.Stdin = &kubectlOut
	grepCmd.Stdout = &grepOut
	grepCmd.Stderr = &grepErr

	err = grepCmd.Run()
	if err != nil {
		fmt.Printf("Pod not found: %s\n", podName)
		fmt.Println("Error executing grep command:", err)
		fmt.Println("Grep Stderr:", grepErr.String())
		return
	}
	filteredOutput := grepOut.String()
	lines := strings.Split(filteredOutput, "\n")
	if len(lines) == 0 {
		fmt.Println("Pod not found")
		return
	}

	if len(lines) > 0 {
		firstPodName := strings.Fields(lines[0])[0]
		commandString := fmt.Sprintf("kubectl port-forward %s %s", firstPodName, ports)
		// Print the command string
		fmt.Print(commandString)

		// Wait for the user to press ENTER
		fmt.Println("\nPress ENTER to execute the command...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		// Execute the command string
		execCmd := exec.Command("/bin/bash", "-c", commandString)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		err = execCmd.Run()
		if err != nil {
			fmt.Println("Error executing command string:", err)
			return
		}

		fmt.Println("Command executed successfully.")
	}

}

func isPortParamValid(ports string) bool {
	//verify ports
	if len(ports) == 0 || !strings.Contains(ports, ":") || (len(strings.Split(ports, ":")) != 2) {
		fmt.Println("Please specify the ports in the format of p=toPort:fromPort")
		return false
	}
	return true
}
