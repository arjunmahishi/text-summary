package textsummary

// Summarizer holds function(s) to summarize the given text
type Summarizer interface {
	Summary(text string, n int) string
}

// CreateSummarizer creates an instance of Summarizer
func CreateSummarizer() Summarizer {
	return &BasicSummarizer{}
}
