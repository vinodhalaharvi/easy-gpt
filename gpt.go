package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// OpenAIRequest represents the payload for making a request to OpenAI's API.
type OpenAIRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

// OpenAIResponse represents the structure of the response from OpenAI's API.
type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// GPT4 encapsulates the interaction with GPT-4, including the history and HTTP client.
type GPT4 struct {
	APIKey  string
	History []string
	Client  *http.Client
	OpenAIRequest
	OpenAIResponse
}

// NewGPT4 creates a new instance of GPT4 with initialized values.
func NewGPT4(apiKey string) *GPT4 {
	return &GPT4{
		APIKey:  apiKey,
		Client:  &http.Client{},
		History: make([]string, 0),
	}
}

// AddToHistory adds a new user prompt to the history.
func (g *GPT4) AddToHistory(userPrompt string) {
	g.History = append(g.History, "User: "+userPrompt)
}

// Execute sends the request to OpenAI's API and returns the response.
func (g *GPT4) Execute(newPrompt string) (string, error) {
	// Add the new prompt to the history
	g.AddToHistory(newPrompt)

	// Generate the prompt with history
	prompt := ""
	for _, exchange := range g.History {
		prompt += exchange + "\n"
	}

	// Update the request data
	g.Prompt = prompt

	requestBody, err := json.Marshal(g.OpenAIRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(requestBody))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &g.OpenAIResponse)
	if err != nil {
		return "", err
	}

	aiResponse := g.Choices[0].Text
	g.History = append(g.History, "AI: "+aiResponse)

	return aiResponse, nil
}

func (g *GPT4) ClearHistory() {
	g.History = make([]string, 0)
}

func (g *GPT4) SetHistory(history []string) {
	g.History = history
}
