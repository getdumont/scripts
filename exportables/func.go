package exportables

import (
	"time"
    "fmt"
	"io/ioutil"
)

var (
	SENTIMENT_PILOT = "sentiment-pilot"
	SAMPLE_ANALYTIC = "sample-analytic"
	Kinds = []string{SENTIMENT_PILOT, SAMPLE_ANALYTIC}
)

func Run(kind string, path string) {
	var outputValue []byte

	switch kind {
		case SENTIMENT_PILOT:
			outputValue = SentimentPilot()
		case SAMPLE_ANALYTIC:
			outputValue = SampleAnalytic()
	}

	date := time.Now().Format("2006-01-02-15-04")
	outputName := fmt.Sprintf("%s/%s-%s.json", path, kind, date)
    ioutil.WriteFile(outputName, outputValue, 0644)
}