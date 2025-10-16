package ai

type AnalysisInput struct {
	TranscriptText string
	OCRText        string
	Type           string
}

type AnalysisOutput struct {
	Hint       string
	Tasks      []Task
	Confidence float64
}

type Task struct {
	Description string
	Assignee    string
	Priority    string
}
