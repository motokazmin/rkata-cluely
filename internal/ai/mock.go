package ai

import (
	"context"
	"fmt"
	"log"
)

type MockProvider struct{}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error) {
	hint := fmt.Sprintf("Mock AI Analysis for %s: This looks important!", input.Type)

	return AnalysisOutput{
		Hint:       hint,
		Confidence: 0.9,
	}, nil
}

func (m *MockProvider) Health(ctx context.Context) error {
	log.Println("🎭 Mock AI provider is healthy")
	return nil
}
