package vectorsrepresentation

// this file will convert text data into vector representation using TF-IDF algorithm

import (
	"math"
	"strings"
)

// tokenization - split the text into words and calculate the frequency of each word in each documentation
// TF = (Number of occurences of term in document) / (Total number of terms in document)
// Measure the importance of the term within a specific document
func Tokenization(documentation map[int]string) map[int]map[string]float64 { // parameters: an array of documentations
	termFrequencies := make(map[int]map[string]float64, len(documentation))
	for i, doc := range documentation {
		terms := strings.Fields(doc) // seperates docs into terms
		termFrequencies[i] = make(map[string]float64)
		totalTerms := float64(len(terms))
		for _, term := range terms {
			termFrequencies[i][term] += 1.0 / totalTerms // calculate the term frequency inside this document
		}
	}
	return termFrequencies
}

// Calculate the inverse document frequency (IDF) -> the log(total number of documents / Number of documents containing the term) -> measure of how unique or rare a term is across a collection of documents
func CalculateIDF(documentation map[int]string, termFrequencies map[int]map[string]float64) map[string]float64 {
	docCount := len(documentation)
	inverseDocFreq := make(map[string]float64)
	for _, freqMap := range termFrequencies {
		for term := range freqMap {
			inverseDocFreq[term]++
		}
	}

	for term, df := range inverseDocFreq {
		inverseDocFreq[term] = math.Log(float64(docCount) / df)
	}

	return inverseDocFreq
}

// Calculate the TF-IDF values for each document
func CalculateTF_IDF(documentation map[int]string, termFrequencies map[int]map[string]float64, inverseDocFreq map[string]float64) map[int]map[string]float64 {
	tfidf := make(map[int]map[string]float64, len(documentation))
	for i, freqMap := range termFrequencies {
		tfidf[i] = make(map[string]float64)
		for term, tf := range freqMap {
			tfidf[i][term] = tf * inverseDocFreq[term] // the numerical values of the frequency of this term in a documentation
		}
	}
	return tfidf
}

// Calculate the TF-IDF vectors for each document with respect to other documentations
func CalculateTFIDFVectors(tfidfMap map[int]map[string]float64, documentation map[int]string) map[int][]float64 {
	tfidfVector := make(map[int][]float64, len(documentation))
	for i, _ := range documentation {
		terms := strings.Fields(documentation[i])
		tfidfVector[i] = make([]float64, len(terms)) // create a list of with size of each documentation for each vector
	}

	for i, _ := range documentation {
		terms := strings.Fields(documentation[i])
		for j, term := range terms {
			tfidfVector[i][j] = tfidfMap[i][term]
		}
	}

	return tfidfVector
}
