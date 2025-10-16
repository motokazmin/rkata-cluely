package audio

import (
	"context"
	"fmt"
	"log"
)

type AzureTranscriber struct {
	subscriptionKey string
	region          string
	language        string
}

func NewAzureTranscriber(key, region, language string) *AzureTranscriber {
	if language == "" {
		language = "ru-RU"
	}
	return &AzureTranscriber{
		subscriptionKey: key,
		region:          region,
		language:        language,
	}
}

func (a *AzureTranscriber) Initialize() error {
	if a.subscriptionKey == "" {
		return fmt.Errorf("azure subscription key is required")
	}

	// TODO: Initialize Azure Speech SDK
	// import "github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
	// config, err := speech.NewSpeechConfigFromSubscription(a.subscriptionKey, a.region)

	log.Printf("🔧 Azure Speech configured: region=%s, language=%s", a.region, a.language)
	return nil
}

func (a *AzureTranscriber) Transcribe(ctx context.Context, audioData []byte) (string, error) {
	// TODO: Real Azure Speech recognition
	return "azure transcription result", nil
}

func (a *AzureTranscriber) Close() error {
	return nil
}
