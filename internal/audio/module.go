package audio

import (
	"context"
	"fmt"
	"log"
	"time"

	"cluely/internal/config"
)

type Module struct {
	cfg            config.AudioConfig
	transcriptChan chan string
	transcriber    Transcriber
	stopChan       chan struct{}
}

func NewModule(cfg config.AudioConfig) *Module {
	return &Module{
		cfg:            cfg,
		transcriptChan: make(chan string, 10),
		stopChan:       make(chan struct{}),
	}
}

func (m *Module) Start(ctx context.Context) error {
	if !m.cfg.Enabled {
		log.Println("⏭️  Audio module disabled")
		return nil
	}

	// Создаем транскрайбер
	transcriber, err := NewTranscriber(m.cfg.TranscriberType, m.cfg.TranscriberConfig)
	if err != nil {
		return fmt.Errorf("failed to create transcriber: %w", err)
	}

	if err := transcriber.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize transcriber: %w", err)
	}

	m.transcriber = transcriber
	log.Printf("✅ Transcriber initialized: %s", m.cfg.TranscriberType)

	// Пока симулируем захват аудио
	go m.simulateCaptureLoop(ctx)

	return nil
}

func (m *Module) simulateCaptureLoop(ctx context.Context) {
	log.Println("🎙️  Audio capture simulation started")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	testPhrases := []string{
		"У нас проблема с продакшн сервером",
		"Нужно откатиться на предыдущую версию",
		"Проверьте логи в Grafana",
		"CPU использование выше 90 процентов",
	}

	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stopChan:
			return
		case <-ticker.C:
			// Симулируем транскрипцию
			phrase := testPhrases[i%len(testPhrases)]
			i++

			log.Printf("🎙️  Simulated audio: %s", phrase)
			m.transcriptChan <- phrase
		}
	}
}

func (m *Module) TranscriptChannel() <-chan string {
	return m.transcriptChan
}

func (m *Module) Stop() {
	close(m.stopChan)

	if m.transcriber != nil {
		m.transcriber.Close()
	}

	close(m.transcriptChan)
}
