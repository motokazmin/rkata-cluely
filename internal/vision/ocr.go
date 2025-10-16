package vision

import (
	"context"
	"fmt"
)

type OCREngine interface {
	ExtractText(ctx context.Context, imageData []byte) (string, error)
	Initialize() error
	Close() error
}

func NewOCREngine(engineType string, config map[string]string) (OCREngine, error) {
	switch engineType {
	case "tesseract":
		// TODO: Implement Tesseract OCR
		// For now, fallback to mock
		return NewMockOCR(), nil
	case "mock":
		return NewMockOCR(), nil
	default:
		return nil, fmt.Errorf("unknown OCR engine: %s", engineType)
	}
}
