package ai

import "context"

type AnalysisInput struct {
	TranscriptText string // Текст из аудиотранскрипции
	OCRText        string // Текст из OCR скриншотов
	Type           string // "audio", "vision", или "combined"
}

type AnalysisOutput struct {
	Hint       string   // Краткая подсказка для пользователя
	Tasks      []string // Структурированные задачи
	Warnings   []string // Предупреждения о рисках
	Confidence float64  // Уверенность AI (0.0 - 1.0)
}

type AIProvider interface {
	Analyze(ctx context.Context, input AnalysisInput) (AnalysisOutput, error)
	Health(ctx context.Context) error
}
