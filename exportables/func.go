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
	PREDICT_DATA = "predict-data"
	Kinds = []string{SENTIMENT_PILOT, SAMPLE_ANALYTIC, ANSWERED_SAMPLE, PREDICT_DATA}
)

func Run(kind string, path string, no_date bool) {
	var outputValue []byte

	switch kind {
		case SENTIMENT_PILOT:
			outputValue = SentimentPilot()
		case SAMPLE_ANALYTIC:
			outputValue = SampleAnalytic()
		case ANSWERED_SAMPLE:
			outputValue = AnsweredSample()
		case PREDICT_DATA:
			outputValue = PredictData()
	}

	outputName := ""

	if no_date {
		outputName = fmt.Sprintf("%s/%s.json", path, kind)
	} else {
		date := time.Now().Format("2006-01-02-15-04")
		outputName = fmt.Sprintf("%s/%s-%s.json", path, kind, date)

	}

    ioutil.WriteFile(outputName, outputValue, 0644)
}