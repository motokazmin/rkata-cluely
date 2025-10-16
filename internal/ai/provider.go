package ai

import "context"

type AIProvider interface {
	Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error)
	Health(ctx context.Context) error
}
