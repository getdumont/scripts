package exportables

import (
	"time"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

type resumedTweet struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	User bson.ObjectId `bson:"_user" json:"user"`
	CreatedAt *time.Time `bson:"created_at" json:"created_at"`
	CleanSentiment Sentiment `bson:"clean_sentiment" json:"clean_sentiment"`
}

func SentimentPilot() []byte {
	var _tweets []resumedTweet
	tweets, tweetConnClose := ConnectAndGetCollection(LocalConfig, "tweets")

	defer tweetConnClose()
	tweets.Find(nil).All(&_tweets)

	resp, _ := json.Marshal(_tweets)
	return resp
}