package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func PrettyPrintJSON(response string) {
	// Pretty print the JSON response
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(response), "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(prettyJSON.Bytes()))
}

func ExecuteCommandInBash(commandStr string) []byte {
	// Use build.sh to execute the command string
	out, err := exec.Command("bash", "-c", commandStr).CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return nil
	}

	// Return the output
	return out
}

// ExecuteCommand executes a shell command and prints its output
func ExecuteCommand(commandStr string) []byte {
	// Split the command into command and arguments
	parts := strings.Fields(commandStr)
	cmd := parts[0]
	args := parts[1:]

	// Execute the command
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return nil
	}

	// Print the output
	return out
}
func ShowResponse(response string, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(response)
}

func ProcessJSONResponse(response string, err error) {
	// Strip the Markdown code block syntax
	jsonResponse := strings.TrimPrefix(response, "```json")
	jsonResponse = strings.TrimSuffix(jsonResponse, "```")
	jsonResponse = strings.TrimSpace(jsonResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal the JSON into a slice of strings
	var commands []string
	if err := json.Unmarshal([]byte(jsonResponse), &commands); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Execute each command
	for _, cmd := range commands {
		fmt.Printf("Executing command: %s\n", cmd)
		var command = ExecuteCommandInBash(cmd)
		if command != nil {
			fmt.Println(string(command))
		}
	}
}
