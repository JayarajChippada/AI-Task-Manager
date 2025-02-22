package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const openAIURL = "https://api.openai.com/v1/chat/completions"

type AIRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// CallOpenAI - Sends request to OpenAI API
func CallOpenAI(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing OpenAI API key")
	}

	requestBody, _ := json.Marshal(AIRequest{
		Model: "gpt-4",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{"system", "You are an AI assistant helping with task management."},
			{"user", prompt},
		},
	})

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var aiResp AIResponse
	if err := json.Unmarshal(body, &aiResp); err != nil {
		return "", err
	}

	if len(aiResp.Choices) > 0 {
		return aiResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from AI")
}

// GetTaskSuggestions - Generate tasks from project description
func GetTaskSuggestions(projectDescription string) (string, error) {
	prompt := fmt.Sprintf("Generate a list of tasks for the following project: %s", projectDescription)
	return CallOpenAI(prompt)
}

// ImproveTaskDescription - Improve an existing task description
func ImproveTaskDescription(taskDescription string) (string, error) {
	prompt := fmt.Sprintf("Rewrite and improve the clarity of this task description: %s", taskDescription)
	return CallOpenAI(prompt)
}

// AssignTaskPriority - Categorize a task into Low, Medium, or High priority
func AssignTaskPriority(taskDescription string) (string, error) {
	prompt := fmt.Sprintf("Analyze the following task and suggest its priority (Low, Medium, High): %s", taskDescription)
	return CallOpenAI(prompt)
}
