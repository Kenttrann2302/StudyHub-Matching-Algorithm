package vectorsrepresentation

// this file will convert text data into vector representation using TF-IDF algorithm

import (
	"math"
	"strings"
)

// tokenization - split the text into words and calculate the frequency of each word in each documentation
func Tokenization(documentation []string) []map[string]float64 { // parameters: an array of documentations
	termFrequencies := make([]map[string]float64, len(documentation));
	for i, doc := range documentation {
		terms := strings.Fields(doc); // seperates docs into terms
		termFrequencies[i] = make(map[string]float64);
		totalTerms := float64(len(terms));
		for _, term := range terms {
			termFrequencies[i][term] += 1.0 / totalTerms; // calculate the term frequency inside this document
		}
	}
	return termFrequencies;
}

// Calculate the inverse document frequency (IDF)
func CalculateIDF(documentation []string, termFrequencies []map[string]float64)  map[string]float64 {
	docCount := len(documentation);
	inverseDocFreq := make(map[string]float64);
	for _, freqMap := range termFrequencies {
		for term := range freqMap {
			inverseDocFreq[term]++;
		}
	}

	for term, df := range inverseDocFreq {
		inverseDocFreq[term] = math.Log(float64(docCount) / df);
	}

	return inverseDocFreq;
}

// Calculate the TF-IDF 
func CalculateTF_IDF(documentation []string, termFrequencies []map[string]float64, inverseDocFreq map[string]float64) []map[string]float64 {
	tfidf := make([]map[string]float64, len(documentation));
	for i, freqMap := range termFrequencies {
		tfidf[i] = make(map[string]float64);
		for term, tf := range freqMap {
			tfidf[i][term] = tf * inverseDocFreq[term]; // the numerical values of the frequency of this term in a documentation
		}
	}
	return tfidf;
}



