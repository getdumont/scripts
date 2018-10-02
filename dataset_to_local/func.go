package dataset_to_local

import (
	"fmt"
	"log"
	. "github.com/getdumont/scripts/utilities"
)

func TransferTweets() () {
	atlasCol, atlasClose := ConnectAndGetCollection(AtlasConfig, "tweets")
	localCol, localClose := ConnectAndGetCollection(LocalConfig, "tweets")

	defer atlasClose()
	defer localClose()

	totalTweets, _ := atlasCol.Count()
	totalProcessed := 0

	fmt.Print(" > Start Tweets\n")
	fmt.Printf("  - Proccess Info: %d/%d\n", totalTweets, totalProcessed)
	for totalTweets >= totalProcessed {
		var tweets []Tweet
		atlasCol.Find(nil).Skip(totalProcessed).Limit(250).All(&tweets)

		for _, tweet := range tweets {
			err := localCol.Insert(&tweet)
			if err != nil {
				fmt.Printf("  - Fail At: %s", tweet.Id)
			}
		}

		totalProcessed = totalProcessed + 250
		fmt.Printf("  - Proccess Info: %d/%d\n", totalTweets, totalProcessed)

	}
	fmt.Print("  - End Pull For Tweets")
}

func TransferUsers() () {
	atlasCol, atlasClose := ConnectAndGetCollection(AtlasConfig, "users")
	localCol, localClose := ConnectAndGetCollection(LocalConfig, "users")

	defer atlasClose()
	defer localClose()

	totalUsers, _ := atlasCol.Count()
	totalProcessed := 0

	fmt.Print(" > Start Users\n")
	fmt.Printf("  - Proccess Info: %d/%d\n", totalUsers, totalProcessed)
	for totalUsers >= totalProcessed {
		var users []User
		atlasCol.Find(nil).Skip(totalProcessed).Limit(80).All(&users)

		for _, user := range users {
			err := localCol.Insert(&user)
			if err != nil {
				log.Fatal(err)
			}
		}

		totalProcessed = totalProcessed + 80
		fmt.Printf("  - Proccess Info: %d/%d\n", totalUsers, totalProcessed)

	}
	fmt.Print("  - End Pull For Tweets")
}