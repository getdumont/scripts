package utilities

import (
	"time"
	"github.com/globalsign/mgo/bson"
)

var (
	NEGATIVE_TWEET_MEDIA = 32
	NEGATIVE_PERCENT = 20
)

type TextObject struct {
	Raw 		string 	 `bson:"rawText"`
	PreClear 	string 	 `bson:"preClear"`
	Clear 		string 	 `bson:"clearText"`
	Emojis 		[]string `bson:"emojis"`
}

type Sentiment struct {
	Score 	  float64 `bson:"score"`
	Magnitude float64 `bson:"magnitude"`
}

type Token struct {
	Value	string	`bson:"value"`
	Kind	string	`bson:"pos"`
}

type Tweet struct {
	Id 					bson.ObjectId 	`bson:"_id"`
	ProcessingVersion 	int 			`bson:"processing_version"`
	CreatedAt 			*time.Time 		`bson:"created_at"`
	TweetId 			string 			`bson:"id"`
	TweetStrId 			string 			`bson:"id_str"`
	Text 				string 			`bson:"text"`
	User 				bson.ObjectId 	`bson:"_user"`
	Entities 			bson.M 			`bson:"entities"`
	TextObject 			TextObject 		`bson:"text_object"`
	CleanSentiment 		Sentiment 		`bson:"clean_sentiment"`
	CleanText 			string 			`bson:"clean_text"`
	CleanTextTree 		[]Token 		`bson:"clean_tree"`
	RawSentiment 		Sentiment 		`bson:"raw_sentiment"`
	RawTextTree 		[]Token 		`bson:"raw_tree"`
	WithoutRt 			string 			`bson:"without_rt"`
}

type AnswerResumed struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Questions []QuestionArray `bson:"questions" json:"questions"`
}

func (t *Tweet) GetSentimentScoreAverage() float64 {
	return (t.CleanSentiment.Score + t.RawSentiment.Score) / 2
}

type User struct {
	Id					bson.ObjectId	`bson:"_id"`
	Protected 			bool 			`bson:"protected"`
	ProcessingVersion 	int 			`bson:"processing_version"`
	Description 		string 			`bson:"description"`
	FollowersCount 		int 			`bson:"followers_count"`
	FriendsCount 		int 			`bson:"friends_count"`
	FavouritesCount 	int 			`bson:"favourites_count"`
	TweetId 			string 			`bson:"id"`
	TweetStrId 			string 			`bson:"id_str"`
	ScreenName			string			`bson:"screen_name"`
	CreatedAt 			*time.Time 		`bson:"created_at"`
	ProfileLinkColor 	string 			`bson:"profile_link_color"`
	ProfileSideBorderC 	string 			`bson:"profile_sidebar_border_color"`
	ProfileSideFillC 	string 			`bson:"profile_sidebar_fill_color"`
	ProfileTextColor 	string 			`bson:"profile_text_color"`
}

type Sample struct {
	Id 				 	 bson.ObjectId `bson:"_id"`
	User 			 	 User 		   `bson:"user"`
	OtherTweets		  	 []Tweet 	   `bson:"other_tweets"`
	NegativeTweets  	 []Tweet 	   `bson:"negative_tweets"`
}

type List struct {
	Id			bson.ObjectId   `bson:"_id"`
	Tweets 		[]bson.ObjectId `bson:"tweets"`
	Specialists	[]bson.ObjectId `bson:"specialists_done"`
}

type Question struct {
	Index int `bson:"question_index" json:"question_index"`
	Impact int `bson:"impact" json:"impact"`
}

type QuestionArray []Question

type Answer struct {
	Tweet bson.ObjectId `bson:"to_tweet" json:"to_tweet"`
	Question []Question `bson:"question" json:"question"`
}

func (s *Sample) IsValid() bool {
	otherTweetsQtd := len(s.OtherTweets)
	negativeTweetsQtd := len(s.NegativeTweets)

	totalTweets := otherTweetsQtd + negativeTweetsQtd
	negativePercentage := (float64(negativeTweetsQtd) * 100) / float64(totalTweets)

	if negativeTweetsQtd > NEGATIVE_TWEET_MEDIA && negativePercentage >= float64(NEGATIVE_PERCENT) {
		return true
	}

	return false
}