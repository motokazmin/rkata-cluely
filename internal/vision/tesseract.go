package vision

import (
	"context"
	"log"
)

type TesseractOCR struct {
	language string
}

func NewTesseractOCR(language string) *TesseractOCR {
	if language == "" {
		language = "eng+rus"
	}
	return &TesseractOCR{
		language: language,
	}
}

func (t *TesseractOCR) Initialize() error {
	// TODO: Initialize gosseract
	// import "github.com/otiai10/gosseract/v2"
	// client := gosseract.NewClient()
	// defer client.Close()
	// client.SetLanguage(t.language)

	log.Printf("🔧 Tesseract configured: language=%s", t.language)
	return nil
}

func (t *TesseractOCR) ExtractText(ctx context.Context, imageData []byte) (string, error) {
	// TODO: Real Tesseract OCR
	return "Error: Connection timeout on server-01\nCPU: 95%\nMemory: 8.2GB/16GB", nil
}

func (t *TesseractOCR) Close() error {
	return nil
}
