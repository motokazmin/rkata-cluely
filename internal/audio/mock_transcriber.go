package audio

import (
	"context"
	"log"
	"time"
)

// MockTranscriber имитирует работу транскрибера для тестирования
type MockTranscriber struct {
	counter int
}

func NewMockTranscriber() *MockTranscriber {
	return &MockTranscriber{counter: 0}
}

func (m *MockTranscriber) Initialize() error {
	log.Println("🎤 MockTranscriber initialized")
	return nil
}

func (m *MockTranscriber) Transcribe(ctx context.Context, audioData []byte) (string, error) {
	m.counter++

	// Симулируем различные сценарии инцидентов
	mockTranscripts := []string{
		"У нас критическая проблема с CPU. Нагрузка 95 процентов!",
		"Memory leak обнаружен в последнем деплойе. Нужно откатиться.",
		"Сервер не отвечает. Проверьте логи в /var/log/app.log",
		"Database connection timeout. Возможно, network issue.",
		"API возвращает 500 ошибки последний час.",
	}

	transcript := mockTranscripts[m.counter%len(mockTranscripts)]
	log.Printf("🎤 Mock Transcribe #%d: %s", m.counter, transcript)

	// Симулируем задержку
	select {
	case <-time.After(500 * time.Millisecond):
		return transcript, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (m *MockTranscriber) Close() error {
	log.Println("🎤 MockTranscriber closed")
	return nil
}
