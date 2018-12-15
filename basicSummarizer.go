package textsummary

import (
	"regexp"
	"strings"
)

var wordWeights = make(map[string]int)

type sent struct {
	text   string
	weight int
}

// BasicSummarizer implements the Summarizer interface
type BasicSummarizer struct{}

// Summary of the given text will be returned. The summary will contain 'n' lines
func (bs *BasicSummarizer) Summary(text string, noOfLines int) string {
	ppPara := preProcess(text)
	sentWeights := []sent{}

	for _, word := range strings.Split(ppPara, " ") {
		word = strings.TrimSpace(word)
		wordWeights[word] = strings.Count(ppPara, word)
	}

	for _, s := range strings.Split(text, ".") {
		s = strings.TrimSpace(s)
		if s != "" {
			sentWeights = append(sentWeights, sent{text: s, weight: getSentWeight(s)})
		}
	}
	sentWeights = sortSent(sentWeights)

	if len(sentWeights) < noOfLines {
		noOfLines = len(sentWeights)
	}

	summary := ""
	for i := 0; i < noOfLines; i++ {
		summary += sentWeights[i].text + ". "
	}

	return strings.TrimSpace(summary)
}

func preProcess(text string) string {
	re := regexp.MustCompile(`(?m)[a-zA-Z ]+`)
	stopWords := []string{"a", "about", "above", "after", "again", "against", "all", "am", "an", "and", "any", "are", "as", "at", "be", "because", "been", "before", "being", "below", "between", "both", "but", "by", "could", "did", "do", "does", "doing", "down", "during", "each", "few", "for", "from", "further", "had", "has", "have", "having", "he", "he'd", "he'll", "he's", "her", "here", "here's", "hers", "herself", "him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm", "i've", "if", "in", "into", "is", "it", "it's", "its", "itself", "let's", "me", "more", "most", "my", "myself", "nor", "of", "on", "once", "only", "or", "other", "ought", "our", "ours", "ourselves", "out", "over", "own", "same", "she", "she'd", "she'll", "she's", "should", "so", "some", "such", "than", "that", "that's", "the", "their", "theirs", "them", "themselves", "then", "there", "there's", "these", "they", "they'd", "they'll", "they're", "they've", "this", "those", "through", "to", "too", "under", "until", "up", "very", "was", "we", "we'd", "we'll", "we're", "we've", "were", "what", "what's", "when", "when's", "where", "where's", "which", "while", "who", "who's", "whom", "why", "why's", "with", "would", "you", "you'd", "you'll", "you're", "you've", "your", "yours", "yourself", "yourselves"}

	str := ""
	for _, match := range re.FindAllString(text, -1) {
		str += match
	}

	for _, word := range stopWords {
		str = strings.Replace(str, " "+word+" ", " ", -1)
	}

	return strings.TrimSpace(str)
}

func getSentWeight(sent string) int {
	weight := 0
	biggest := 0
	for _, word := range strings.Split(sent, " ") {
		w, ok := wordWeights[word]
		if ok {
			weight += w
			if w > biggest {
				biggest = w
			}
		}
	}
	return weight / biggest
}

func sortSent(sents []sent) []sent {
	for i := 0; i < len(sents); i++ {
		for j := 0; j < len(sents)-1; j++ {
			if sents[j].weight < sents[j+1].weight {
				temp := sents[j]
				sents[j] = sents[j+1]
				sents[j+1] = temp
			}
		}
	}
	return sents
}
