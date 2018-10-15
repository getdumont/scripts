package exportables

import (
	"time"
	"encoding/json"
    "fmt"
	"io/ioutil"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

type resumedTweet struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	User bson.ObjectId `bson:"_user" json:"user"`
	CleanSentiment Sentiment `bson:"clean_sentiment" json:"clean_sentiment"`
	RawSentiment Sentiment `bson:"raw_sentiment" json:"raw_sentiment"`
}

func SentimentPilot(path string) {
	var _tweets []resumedTweet
	date := time.Now().Format("2006-01-02-15-04")
	tweets, tweetConnClose := ConnectAndGetCollection(LocalConfig, "tweets")

	defer tweetConnClose()

	tweets.Find(nil).All(&_tweets)

	tweetJson, _ := json.Marshal(_tweets)
	outputName := fmt.Sprintf("%s/sentiment-pilot-%s.json", path, date)
    ioutil.WriteFile(outputName, tweetJson, 0644)
}