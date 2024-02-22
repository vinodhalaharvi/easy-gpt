package assistant

import (
	"bufio"
	"bytes"
	"easy-gpt/config"
	"easy-gpt/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Message Define the message structure for the chat API
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GPTAssistant encapsulates interactions with the GPT-4 Assistant via OpenAI's chat API.
type GPTAssistant struct {
	APIKey   string
	Messages []Message
	Client   *http.Client
	Config   *config.Config
}

// NewGPTAssistant creates a new instance of GPTAssistant with initialized values.
func NewGPTAssistant(apiKey string, config *config.Config) *GPTAssistant {
	assistant := GPTAssistant{
		APIKey:   apiKey,
		Messages: make([]Message, 0),
		Client:   &http.Client{},
		Config:   config,
	}
	return &assistant
}

// AddMessage adds a new message to the conversation.
func (g *GPTAssistant) AddMessage(role, content string) {
	g.Messages = append(g.Messages, Message{Role: role, Content: content})
}

// Execute sends the conversation to OpenAI's API and returns the assistant's response.
func (g *GPTAssistant) Execute() (string, error) {
	requestData := map[string]interface{}{
		"model":    g.Config.Assistant.Model,
		"messages": g.Messages,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// Use the method and URL from the config
	req, err := http.NewRequest(
		g.Config.Assistant.RequestMethod,
		g.Config.Assistant.RequestURL,
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+g.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := g.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response ChatCompletionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 && len(response.Choices[0].Message.Content) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		LogProbs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func (g *GPTAssistant) ExecuteStdinInstructions() {
	g.AddMessage("system", g.Config.Assistant.SystemMessage)

	g.GetUserQueryFromStdin()

	// Execute the request
	response, err := g.Execute()

	utils.ProcessJSONResponse(response, err)
}

func (g *GPTAssistant) GetUserQueryFromStdin() {
	// Read user queries from stdin
	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("Enter gcloud operation queries (press CTRL+D to submit):")
	for scanner.Scan() {
		userQuery := scanner.Text()
		g.AddMessage("user", userQuery)
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}
