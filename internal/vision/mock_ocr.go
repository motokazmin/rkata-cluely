package vision

import (
	"context"
	"log"
	"time"
)

// MockOCR имитирует работу OCR для тестирования
type MockOCR struct {
	counter int
}

func NewMockOCR() *MockOCR {
	return &MockOCR{counter: 0}
}

func (m *MockOCR) Initialize() error {
	log.Println("📸 MockOCR initialized")
	return nil
}

func (m *MockOCR) ExtractText(ctx context.Context, imageData []byte) (string, error) {
	m.counter++

	// Симулируем различные типы логов и метрик
	mockOCRTexts := []string{
		"ERROR: Connection refused\nStack trace: ...\n[ERROR] Database pool exhausted",
		"CPU: 95%\nMemory: 8.2GB/16GB\nDisk I/O: 85%\nNetwork: 450Mbps",
		"2025-10-16 14:23:45 [ERROR] Failed to connect to service\n2025-10-16 14:23:46 [WARN] Retrying...\n2025-10-16 14:23:47 [INFO] Connection restored",
		"Pod Status: CrashLoopBackOff\nRestarts: 5\nLast Error: OutOfMemory",
		"HTTP/1.1 503 Service Unavailable\nRetry-After: 60\nContent-Length: 1234",
	}

	ocrText := mockOCRTexts[m.counter%len(mockOCRTexts)]
	log.Printf("📸 Mock OCR #%d:\n%s", m.counter, ocrText)

	// Симулируем задержку обработки
	select {
	case <-time.After(800 * time.Millisecond):
		return ocrText, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (m *MockOCR) Close() error {
	log.Println("📸 MockOCR closed")
	return nil
}
