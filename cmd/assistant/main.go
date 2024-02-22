package main

import (
	"easy-gpt/assistant"
	"easy-gpt/config"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API environment variable not set")
		return
	}

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		fmt.Println("Error: CONFIG_FILE_PATH environment variable not set")
		return
	}

	gptAssistant := assistant.NewGPTAssistant(
		apiKey,
		(&config.Config{}).FromYAMLFile(configFilePath),
	)
	gptAssistant.ExecuteStdinInstructions()
}
