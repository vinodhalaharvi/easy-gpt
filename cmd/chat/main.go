package main

import (
	"easy-gpt/chat"
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

	gptChat := chat.NewGPTChat(
		apiKey,
		(&config.Config{}).FromYAMLFile(configFilePath),
	)
	gptChat.ExecuteStdinInstructions()

}
