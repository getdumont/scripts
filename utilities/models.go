package utilities

import (
	"time"
	"github.com/globalsign/mgo/bson"
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
	CleanTextTree 		bson.M 			`bson:"clean_tree"`
	RawSentiment 		Sentiment 		`bson:"raw_sentiment"`
	RawTextTree 		bson.M 			`bson:"raw_tree"`
	WithoutRt 			string 			`bson:"without_rt"`
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

