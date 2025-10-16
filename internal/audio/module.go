package audio

import (
	"context"
	"log"
	"time"

	"cluely/internal/config"
)

// Module управляет захватом и транскрипцией аудио
type Module struct {
	cfg         config.AudioConfig
	transcriber Transcriber
	transcripts chan string
	stopCh      chan struct{}
	isRunning   bool
}

func NewModule(cfg config.AudioConfig) *Module {
	return &Module{
		cfg:         cfg,
		transcripts: make(chan string, 10),
		stopCh:      make(chan struct{}),
	}
}

func (m *Module) Start(ctx context.Context) error {
	if !m.cfg.Enabled {
		log.Println("⏭️  Audio module disabled")
		return nil
	}

	// Создаем транскрибер на основе конфига
	transcriber, err := NewTranscriber(m.cfg.TranscriberType, m.cfg.TranscriberConfig)
	if err != nil {
		return err
	}

	if err := transcriber.Initialize(); err != nil {
		return err
	}

	m.transcriber = transcriber
	m.isRunning = true

	// Запускаем горутину для симуляции аудиоввода (в mock режиме)
	go m.simulateAudioCapture(ctx)

	log.Printf("✅ Audio Module started (transcriber: %s)", m.cfg.TranscriberType)
	return nil
}

func (m *Module) simulateAudioCapture(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stopCh:
			return
		case <-ticker.C:
			// Симулируем захват аудио
			if transcript, err := m.transcriber.Transcribe(ctx, nil); err == nil && transcript != "" {
				select {
				case m.transcripts <- transcript:
				case <-ctx.Done():
					return
				}
			}
		}
	}
}

func (m *Module) TranscriptChannel() <-chan string {
	return m.transcripts
}

func (m *Module) Stop() {
	if !m.isRunning {
		return
	}

	m.isRunning = false
	close(m.stopCh)

	if m.transcriber != nil {
		m.transcriber.Close()
	}

	close(m.transcripts)
	log.Println("🛑 Audio Module stopped")
}
