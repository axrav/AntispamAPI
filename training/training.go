package training

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/navossoc/bayesian"
)

// classifying the types
const (
	HAM  bayesian.Class = "ham"
	SPAM bayesian.Class = "spam"
)

// default classifier
var classifier = bayesian.NewClassifier(SPAM, HAM)

// read the dataset
func ReadDataset() [][]string {
	file, _ := os.Open("dataset.csv")
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rawCSVdata

}

// training with the dataset
func TrainModel(dataset [][]string) {
	var spamDataset []string
	var hamDataset []string
	for _, data := range dataset {
		if data[0] == "spam" {
			spamWords := strings.Split(data[1], " ")
			spamDataset = append(spamDataset, spamWords...)
		} else {
			hamWords := strings.Split(data[1], " ")
			hamDataset = append(hamDataset, hamWords...)
		}
	}
	classifier.Learn(spamDataset, SPAM)
	classifier.Learn(hamDataset, HAM)

}

// predict the scores
func PredictScores(words string) ([]float64, bool) {
	wordSlice := strings.Split(words, " ")
	scores, likely, _ := classifier.ProbScores(
		wordSlice)
	spamScore := math.Ceil(scores[0] * 100)
	if len(words) < 20 { // if the message is too short, it is unlikely to be spam
		spamScore -= 25
	}
	hamScore := 100 - spamScore
	var isSpam bool
	if likely == 0 && spamScore > 60 {
		isSpam = true
	} else {
		isSpam = false
	}
	return []float64{spamScore, hamScore}, isSpam

}
