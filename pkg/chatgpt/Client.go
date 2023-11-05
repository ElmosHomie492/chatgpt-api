package chatgpt

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

func Init() (*OpenAIClient, error) {
	var openAIClient OpenAIClient
	openAIClient.Client = resty.New()

	openAIClient.ApiKey = os.Getenv("CHATGPT_API_KEY")
	if openAIClient.ApiKey == "" {
		return nil, errors.New("API key not found")
	}

	return &openAIClient, nil
}

func (openAIClient *OpenAIClient) AskGPT(query string) (*string, error) {
	response, err := openAIClient.Client.R().
		SetAuthToken(openAIClient.ApiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": query}},
			"max_tokens": 3500,
		}).
		Post(API_ENDPOINT)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while sending send the request: %v", err))
	}

	body := response.Body()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while decoding JSON response: %v", err))
	}

	content := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return &content, nil
}

func (openAIClient *OpenAIClient) CheckAPIKey(apiKey string) bool {
	providedApiKey, _ := b64.StdEncoding.DecodeString(apiKey)
	if strings.TrimSuffix(string(providedApiKey), "\n") != openAIClient.ApiKey {
		return false
	}

	return true
}
