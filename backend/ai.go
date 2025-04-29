package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func IsOllamaRunning() bool {
	client := http.Client{
		Timeout: 2 * time.Second, // Short timeout
	}
	resp, err := client.Get("http://localhost:11434")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

func GenerateAIFortune() string {
	if !IsOllamaRunning() {
		return "⚠️ AI server not available."
	}

	body := map[string]interface{}{
		"model":  "mistral",
		"prompt": "Generate a short, mysterious and wise one-sentence fortune.",
		"stream": false,
	}
	jsonData, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ Ollama Error:", err)
		return "⚠️ Could not connect to AI server."
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var result struct {
		Response string `json:"response"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("❌ Response Error:", err)
		return "⚠️ Invalid AI response."
	}

	return result.Response
}
