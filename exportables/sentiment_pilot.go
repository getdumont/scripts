package exportables

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

type resumedTweet struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	User bson.ObjectId `bson:"_user" json:"user"`
	CleanSentiment Sentiment `bson:"clean_sentiment" json:"clean_sentiment"`
	RawSentiment Sentiment `bson:"raw_sentiment" json:"raw_sentiment"`
}

func SentimentPilot() []byte {
	var _tweets []resumedTweet
	tweets, tweetConnClose := ConnectAndGetCollection(LocalConfig, "tweets")

	defer tweetConnClose()
	tweets.Find(nil).All(&_tweets)

	resp, _ := json.Marshal(_tweets)
	return resp
}