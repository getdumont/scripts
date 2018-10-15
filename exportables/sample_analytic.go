package exportables

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

type resumedSampleUser struct {
	Id					bson.ObjectId	`bson:"_id" json:"id"`
	FollowersCount 		int 			`bson:"followers_count" json:"followers_count"`
	FriendsCount 		int 			`bson:"friends_count" json:"friends_count"`
	FavouritesCount 	int 			`bson:"favourites_count" json:"favourites_count"`
}

type resumedSampleTweet struct {
	Id 			   bson.ObjectId `bson:"_id" json:"id"`
	Text 		   string 		 `bson:"clean_text" json:"text"`
	User 		   bson.ObjectId `bson:"_user" json:"user"`
	CleanSentiment Sentiment 	 `bson:"clean_sentiment" json:"clean_sentiment"`
}

type resumedSample struct {
	Id 				bson.ObjectId 		 `bson:"_id" json:"id"`
	User 			resumedSampleUser 	 `bson:"user"`
	OtherTweets		[]resumedSampleTweet `bson:"other_tweets"`
	NegativeTweets  []resumedSampleTweet `bson:"negative_tweets"`
}

func SampleAnalytic() []byte {
	var _samples []resumedSample
	samples, sampleConnClose := ConnectAndGetCollection(LocalConfig, "samples")

	defer sampleConnClose()

	samples.Find(nil).All(&_samples)

	resp, _ := json.Marshal(_samples)
	return resp
}