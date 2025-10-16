package audio

import (
	"context"
	"log"
)

type MockTranscriber struct{}

func NewMockTranscriber() *MockTranscriber {
	return &MockTranscriber{}
}

func (m *MockTranscriber) Initialize() error {
	log.Println("🎭 Mock transcriber initialized")
	return nil
}

func (m *MockTranscriber) Transcribe(ctx context.Context, audioData []byte) (string, error) {
	return "mock transcription", nil
}

func (m *MockTranscriber) Close() error {
	return nil
}
