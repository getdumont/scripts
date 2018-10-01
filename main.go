package main

import "github.com/getdumont/scripts/dataset_to_local"

func main() {
	dataset_to_local.TransferUsers()
	dataset_to_local.TransferTweets()
}