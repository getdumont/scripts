package select_sample

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

var (
	LIMIT_SCORE = (float64(0.2) * -1)
	LIST_LIMIT = 200
)

func negativeTweetsProcess(tweets []Tweet) ([]Tweet, []Tweet) {
	negativeTweets := []Tweet{}
	otherTweets := []Tweet{}

	for _, tweet := range tweets {
		if tweet.GetSentimentScoreAverage() < float64(LIMIT_SCORE) {
			negativeTweets = append(negativeTweets, tweet)
		} else {
			otherTweets = append(otherTweets, tweet)
		}
	}

	return negativeTweets, otherTweets
}

func Run(processing_version int16) {
	fmt.Print("Sample Command \n")
	fmt.Print("  > Connections Open\n")

	users, userConnClose := ConnectAndGetCollection(LocalConfig, "users")
	tweets, tweetConnClose := ConnectAndGetCollection(LocalConfig, "tweets")
	samples, sampleConnClose := ConnectAndGetCollection(LocalConfig, "samples")
	lists, listConnClose := ConnectAndGetCollection(LocalConfig, "lists")

	defer userConnClose()
	defer tweetConnClose()
	defer sampleConnClose()
	defer listConnClose()

	var _users []User
	var totalTweets []Tweet

	users.Find(nil).All(&_users)
	fmt.Printf("  > Get %d users\n", len(_users))

	totalSamples := 0
	for _, user := range _users {
		var _tweets []Tweet
		tweets.Find(bson.M{
			"processing_version": processing_version,
			"_user": user.Id,
		}).All(&_tweets)

		negativeTweets, otherTweets := negativeTweetsProcess(_tweets)

		sample := Sample{
			bson.NewObjectId(),
			user,
			otherTweets,
			negativeTweets,
		}

		if sample.IsValid() {
			totalSamples = totalSamples + 1
			totalTweets = append(totalTweets, _tweets...)
			samples.Insert(&sample)
		}
	}

	fmt.Printf("  > Total of %d Samples Created\n", totalSamples)

	totalLists := 0
	listCount := 0
	listPersist := List{
		bson.NewObjectId(),
		[]bson.ObjectId{},
		[]bson.ObjectId{},
	}

	for _, tweetInSample := range totalTweets {
		if listCount > LIST_LIMIT {
			lists.Insert(&listPersist)
			totalLists = totalLists + 1
			listCount = 0

			listPersist = List{
				bson.NewObjectId(),
				[]bson.ObjectId{},
				[]bson.ObjectId{},
			}
		}

		listPersist.Tweets = append(listPersist.Tweets, tweetInSample.Id)
		listCount = listCount + 1
	}

	fmt.Printf("  > Total of %d Lists Created\n", totalLists)
}