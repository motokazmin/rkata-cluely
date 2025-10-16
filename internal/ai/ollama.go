package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type OllamaProvider struct {
	baseURL string
	model   string
	client  *http.Client
}

func NewOllamaProvider(url, model string) *OllamaProvider {
	if url == "" {
		url = "http://localhost:11434"
	}
	if model == "" {
		model = "llama3.2:latest"
	}

	return &OllamaProvider{
		baseURL: url,
		model:   model,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func (o *OllamaProvider) Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error) {
	prompt := o.buildPrompt(input)

	reqBody := ollamaRequest{
		Model:  o.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return AnalysisOutput{}, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", o.baseURL+"/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return AnalysisOutput{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return AnalysisOutput{}, fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AnalysisOutput{}, fmt.Errorf("ollama returned status %d", resp.StatusCode)
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return AnalysisOutput{}, err
	}

	return AnalysisOutput{
		Hint:       ollamaResp.Response,
		Confidence: 0.85,
	}, nil
}

func (o *OllamaProvider) buildPrompt(input AnalysisInput) string {
	var prompt string

	if input.Type == "audio" {
		prompt = fmt.Sprintf(`Ты - эксперт SRE помощник для IT-тимлида. 
Проанализируй следующую фразу из инцидент-митинга и дай краткую (1-2 предложения) подсказку или действие:

Фраза: "%s"

Ответь кратко и по делу.`, input.TranscriptText)
	} else if input.Type == "vision" {
		prompt = fmt.Sprintf(`Ты - эксперт SRE помощник. 
Проанализируй текст, извлеченный с экрана (логи, метрики):

Текст: "%s"

Дай краткую оценку проблемы и предложи действие (1-2 предложения).`, input.OCRText)
	}

	return prompt
}

func (o *OllamaProvider) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", o.baseURL+"/api/tags", nil)
	if err != nil {
		return err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return fmt.Errorf("ollama health check failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ollama not healthy: status %d", resp.StatusCode)
	}

	log.Printf("✅ Ollama is healthy: %s (model: %s)", o.baseURL, o.model)
	return nil
}
