package ai

import (
	"context"

	"cluely/internal/config"
)

type Module struct {
	cfg      config.AIConfig
	provider AIProvider
}

func NewModule(cfg config.AIConfig) *Module {
	var provider AIProvider

	switch cfg.Provider {
	case "ollama":
		provider = NewOllamaProvider(cfg.OllamaURL, cfg.Model)
	case "mock":
		provider = NewMockProvider()
	default:
		provider = NewOllamaProvider(cfg.OllamaURL, cfg.Model)
	}

	return &Module{
		cfg:      cfg,
		provider: provider,
	}
}

func (m *Module) Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error) {
	return m.provider.Analyze(ctx, input)
}

func (m *Module) Health(ctx context.Context) error {
	return m.provider.Health(ctx)
}
