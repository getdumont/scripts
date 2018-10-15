package exportables

var (
	SENTIMENT_PILOT = "sentiment-pilot"
	Kinds = []string{SENTIMENT_PILOT}
)

func Run(kind string, path string) {
	switch kind {
		case SENTIMENT_PILOT:
			SentimentPilot(path)
	}
}