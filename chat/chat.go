package chat

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

// Message structure as defined in assistant.go remains the same
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCompletionResponse to mirror the response structure specific to chat interactions
type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int         `json:"index"`
		Message      Message     `json:"message"`
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

// GPTChat to encapsulate chat interactions
type GPTChat struct {
	APIKey   string
	Messages []Message
	Client   *http.Client
	Config   *config.Config
}

// NewGPTChat creates a new instance of GPTChat with initialized values.
func NewGPTChat(apiKey string, config *config.Config) *GPTChat {
	return &GPTChat{
		APIKey:   apiKey,
		Messages: make([]Message, 0),
		Client:   &http.Client{},
		Config:   config,
	}
}

// AddMessage adds a new message to the conversation.
func (g *GPTChat) AddMessage(role, content string) {
	g.Messages = append(g.Messages, Message{Role: role, Content: content})
}

// Execute sends the conversation to OpenAI's API and returns the chat's response.
func (g *GPTChat) Execute() (string, error) {
	requestData := map[string]interface{}{
		"model":    g.Config.Chat.Model,
		"messages": g.Messages,
	}
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		g.Config.Chat.RequestMethod,
		g.Config.Chat.RequestURL,
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

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from OpenAI")
}

func (g *GPTChat) ExecuteStdinInstructions() {
	for _, message := range g.Config.Chat.Messages {
		g.AddMessage(message.Role, message.Content)
	}

	g.GetUserQueryFromStdin()

	// Execute the request
	response, err := g.Execute()

	utils.ShowResponse(response, err)
}

func (g *GPTChat) GetUserQueryFromStdin() {
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
