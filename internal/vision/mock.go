package vision

import (
	"context"
	"log"
)

type MockOCR struct{}

func NewMockOCR() *MockOCR {
	return &MockOCR{}
}

func (m *MockOCR) Initialize() error {
	log.Println("🎭 Mock OCR initialized")
	return nil
}

func (m *MockOCR) ExtractText(ctx context.Context, imageData []byte) (string, error) {
	return "mock ocr result", nil
}

func (m *MockOCR) Close() error {
	return nil
}
