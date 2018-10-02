package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/getdumont/scripts/select_sample"
	"github.com/getdumont/scripts/dataset_to_local"
)

var (
	app = kingpin.New("dumont_scripts", "CLI for database edition of dumont project")

	pull = app.Command("pull", "Bring cloud database to local")
	sample = app.Command("sample", "Mount a sample collection")
)
func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case pull.FullCommand():
			dataset_to_local.TransferUsers()
			dataset_to_local.TransferTweets()
		case sample.FullCommand():
			select_sample.Run()
	}
}