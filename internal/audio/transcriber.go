package audio

import (
	"context"
	"fmt"
)

type Transcriber interface {
	Transcribe(ctx context.Context, audioData []byte) (string, error)
	Initialize() error
	Close() error
}

func NewTranscriber(transcriberType string, config map[string]string) (Transcriber, error) {
	switch transcriberType {
	case "azure":
		return NewAzureTranscriber(
			config["subscription_key"],
			config["region"],
			config["language"],
		), nil
	case "mock":
		return NewMockTranscriber(), nil
	default:
		return nil, fmt.Errorf("unknown transcriber type: %s", transcriberType)
	}
}
