package vision

import (
	"context"
	"log"
	"time"

	"cluely/internal/config"
)

// Module управляет захватом скриншотов и OCR обработкой
type Module struct {
	cfg         config.VisionConfig
	ocrEngine   OCREngine
	screenshots chan []byte
	stopCh      chan struct{}
	isRunning   bool
}

func NewModule(cfg config.VisionConfig) *Module {
	return &Module{
		cfg:         cfg,
		screenshots: make(chan []byte, 10),
		stopCh:      make(chan struct{}),
	}
}

func (m *Module) Start(ctx context.Context) error {
	if !m.cfg.Enabled {
		log.Println("⏭️  Vision module disabled")
		return nil
	}

	// Создаем OCR движок на основе конфига
	ocrEngine, err := NewOCREngine(m.cfg.OCREngine, m.cfg.OCRConfig)
	if err != nil {
		return err
	}

	if err := ocrEngine.Initialize(); err != nil {
		return err
	}

	m.ocrEngine = ocrEngine
	m.isRunning = true

	// Запускаем горутину для симуляции захвата скриншотов
	go m.simulateScreenshotCapture(ctx)

	log.Printf("✅ Vision Module started (OCR engine: %s)", m.cfg.OCREngine)
	return nil
}

func (m *Module) simulateScreenshotCapture(ctx context.Context) {
	// В mock режиме отправляем скриншоты каждые 10 секунд
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-m.stopCh:
			return
		case <-ticker.C:
			// Симулируем захват скриншота (в реальной версии это был бы bytes от скриншота)
			dummyScreenshot := []byte("mock_screenshot_data")
			select {
			case m.screenshots <- dummyScreenshot:
				log.Println("📸 Mock screenshot captured")
			case <-ctx.Done():
				return
			}
		}
	}
}

func (m *Module) ExtractText(ctx context.Context, imageData []byte) (string, error) {
	if m.ocrEngine == nil {
		return "", nil
	}
	return m.ocrEngine.ExtractText(ctx, imageData)
}

func (m *Module) ScreenshotChannel() <-chan []byte {
	return m.screenshots
}

func (m *Module) Stop() {
	if !m.isRunning {
		return
	}

	m.isRunning = false
	close(m.stopCh)

	if m.ocrEngine != nil {
		m.ocrEngine.Close()
	}

	close(m.screenshots)
	log.Println("🛑 Vision Module stopped")
}
