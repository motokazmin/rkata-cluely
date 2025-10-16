package vision

import (
	"context"
	"fmt"
	"log"
	"time"

	"cluely/internal/config"
)

type Module struct {
	cfg            config.VisionConfig
	screenshotChan chan []byte
	ocrEngine      OCREngine
	stopChan       chan struct{}
}

func NewModule(cfg config.VisionConfig) *Module {
	return &Module{
		cfg:            cfg,
		screenshotChan: make(chan []byte, 10),
		stopChan:       make(chan struct{}),
	}
}

func (m *Module) Start(ctx context.Context) error {
	if !m.cfg.Enabled {
		log.Println("⏭️  Vision module disabled")
		return nil
	}

	// Создаем OCR engine
	ocrEngine, err := NewOCREngine(m.cfg.OCREngine, m.cfg.OCRConfig)
	if err != nil {
		return err
	}

	if err := ocrEngine.Initialize(); err != nil {
		return err
	}

	m.ocrEngine = ocrEngine
	log.Printf("✅ OCR Engine initialized: %s", m.cfg.OCREngine)

	// Симулируем захват скриншотов
	go m.simulateScreenshotLoop(ctx)

	return nil
}

func (m *Module) simulateScreenshotLoop(ctx context.Context) {
	log.Println("📸 Screenshot simulation started")
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stopChan:
			return
		case <-ticker.C:
			// Симулируем скриншот
			fakeScreenshot := []byte("fake screenshot data")
			log.Println("📸 Simulated screenshot taken")
			m.screenshotChan <- fakeScreenshot
		}
	}
}

func (m *Module) ScreenshotChannel() <-chan []byte {
	return m.screenshotChan
}

func (m *Module) ExtractText(ctx context.Context, imageData []byte) (string, error) {
	if m.ocrEngine == nil {
		return "", fmt.Errorf("OCR engine not initialized")
	}
	return m.ocrEngine.ExtractText(ctx, imageData)
}

func (m *Module) Stop() {
	close(m.stopChan)

	if m.ocrEngine != nil {
		m.ocrEngine.Close()
	}

	close(m.screenshotChan)
}
