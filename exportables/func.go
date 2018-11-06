package exportables

import (
	"time"
    "fmt"
	"io/ioutil"
)

var (
	SENTIMENT_PILOT = "sentiment-pilot"
	SAMPLE_ANALYTIC = "sample-analytic"
	ANSWERED_SAMPLE = "answered-sample"
	Kinds = []string{SENTIMENT_PILOT, SAMPLE_ANALYTIC, ANSWERED_SAMPLE}
)

func Run(kind string, path string) {
	var outputValue []byte

	switch kind {
		case SENTIMENT_PILOT:
			outputValue = SentimentPilot()
		case SAMPLE_ANALYTIC:
			outputValue = SampleAnalytic()
		case ANSWERED_SAMPLE:
			outputValue = AnsweredSample()
	}

	date := time.Now().Format("2006-01-02-15-04")
	outputName := fmt.Sprintf("%s/%s-%s.json", path, kind, date)
    ioutil.WriteFile(outputName, outputValue, 0644)
}