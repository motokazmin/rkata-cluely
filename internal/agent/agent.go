package agent

import (
	"context"
	"log"
	"sync"

	"cluely/internal/ai"
	"cluely/internal/audio"
	"cluely/internal/config"
	"cluely/internal/ui"
	"cluely/internal/vision"
)

type Agent struct {
	cfg          *config.Config
	audioModule  *audio.Module
	visionModule *vision.Module
	aiModule     *ai.Module
	uiServer     *ui.Server
	wg           sync.WaitGroup
}

func New(cfg *config.Config) *Agent {
	return &Agent{
		cfg:          cfg,
		audioModule:  audio.NewModule(cfg.Audio),
		visionModule: vision.NewModule(cfg.Vision),
		aiModule:     ai.NewModule(cfg.AI),
		uiServer:     ui.NewServer(cfg.UI),
	}
}

func (a *Agent) Start(ctx context.Context) error {
	// Запускаем UI сервер
	if a.cfg.UI.Enabled {
		if err := a.uiServer.Start(); err != nil {
			return err
		}
		log.Println("✅ UI Server started")
	}

	// Запускаем аудио модуль
	if a.cfg.Audio.Enabled {
		if err := a.audioModule.Start(ctx); err != nil {
			log.Printf("⚠️  Audio module failed: %v (continuing without audio)", err)
		} else {
			log.Println("✅ Audio Module started")
		}
	}

	// Запускаем визуальный модуль
	if a.cfg.Vision.Enabled {
		if err := a.visionModule.Start(ctx); err != nil {
			log.Printf("⚠️  Vision module failed: %v (continuing without vision)", err)
		} else {
			log.Println("✅ Vision Module started")
		}
	}

	// Проверяем AI
	if err := a.aiModule.Health(ctx); err != nil {
		log.Printf("⚠️  AI module health check failed: %v", err)
	} else {
		log.Println("✅ AI Module ready")
	}

	// Запускаем обработку событий
	a.wg.Add(1)
	go a.processingLoop(ctx)

	return nil
}

func (a *Agent) processingLoop(ctx context.Context) {
	defer a.wg.Done()

	log.Println("🔄 Processing loop started")

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Processing loop stopped")
			return

		case transcript, ok := <-a.audioModule.TranscriptChannel():
			if !ok {
				continue
			}
			a.handleTranscript(ctx, transcript)

		case screenshot, ok := <-a.visionModule.ScreenshotChannel():
			if !ok {
				continue
			}
			a.handleScreenshot(ctx, screenshot)
		}
	}
}

func (a *Agent) handleTranscript(ctx context.Context, text string) {
	log.Printf("🎤 Transcript: %s", text)

	input := ai.AnalysisInput{
		TranscriptText: text,
		Type:           "audio",
	}

	result, err := a.aiModule.Analyze(ctx, input)
	if err != nil {
		log.Printf("❌ AI analysis error: %v", err)
		return
	}

	log.Printf("🤖 AI Hint: %s", result.Hint)

	if a.cfg.UI.Enabled {
		a.uiServer.SendHint(result.Hint)
	}
}

func (a *Agent) handleScreenshot(ctx context.Context, data []byte) {
	log.Printf("📸 Screenshot captured: %d bytes", len(data))

	ocrText, err := a.visionModule.ExtractText(ctx, data)
	if err != nil {
		log.Printf("❌ OCR error: %v", err)
		return
	}

	log.Printf("📝 OCR Text: %s", ocrText)

	input := ai.AnalysisInput{
		OCRText: ocrText,
		Type:    "vision",
	}

	result, err := a.aiModule.Analyze(ctx, input)
	if err != nil {
		log.Printf("❌ AI analysis error: %v", err)
		return
	}

	log.Printf("🤖 AI Hint: %s", result.Hint)

	if a.cfg.UI.Enabled {
		a.uiServer.SendHint(result.Hint)
	}
}

func (a *Agent) Stop() {
	log.Println("🛑 Stopping modules...")

	a.audioModule.Stop()
	a.visionModule.Stop()
	a.uiServer.Stop()

	a.wg.Wait()
	log.Println("✅ All modules stopped")
}
