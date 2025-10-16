package ai

import (
	"context"
	"log"

	"cluely/internal/config"
)

// Module управляет AI анализом контекста
type Module struct {
	cfg      config.AIConfig
	provider AIProvider
}

func NewModule(cfg config.AIConfig) *Module {
	return &Module{
		cfg: cfg,
	}
}

func (m *Module) Initialize(ctx context.Context) error {
	// Создаем AI провайдера на основе конфига
	var provider AIProvider
	var err error

	switch m.cfg.Provider {
	case "ollama":
		provider = NewOllamaProvider(m.cfg.OllamaURL, m.cfg.Model)
	case "mock":
		provider = NewMockAIProvider()
	default:
		provider = NewMockAIProvider()
		log.Printf("⚠️  Unknown AI provider '%s', using mock", m.cfg.Provider)
	}

	m.provider = provider

	// Проверяем здоровье провайдера
	if err := m.provider.Health(ctx); err != nil {
		log.Printf("⚠️  AI provider health check failed: %v (will try to continue)", err)
		// Не возвращаем ошибку, чтобы система работала даже если AI недоступен
	}

	log.Printf("✅ AI Module initialized (provider: %s, model: %s)", m.cfg.Provider, m.cfg.Model)
	return err
}

func (m *Module) Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error) {
	if m.provider == nil {
		// Возвращаем default output если провайдер не инициализирован
		return AnalysisOutput{
			Hint:       "ℹ️ AI модуль недоступен, используйте стандартные инструменты диагностики.",
			Confidence: 0.0,
		}, nil
	}

	return m.provider.Analyze(ctx, input)
}

func (m *Module) Health(ctx context.Context) error {
	if m.provider == nil {
		// Инициализируем провайдера если еще не инициализирован
		if err := m.Initialize(ctx); err != nil {
			log.Printf("⚠️  Failed to initialize AI provider: %v", err)
			return err
		}
	}

	return m.provider.Health(ctx)
}
