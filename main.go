package main

import (
	"os"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/getdumont/scripts/exportables"
	"github.com/getdumont/scripts/select_sample"
	"github.com/getdumont/scripts/dataset_to_local"
)

var (
	app = kingpin.New("dumont_scripts", "CLI for database edition of dumont project")
	processing_version = app.Flag("processing_version", "The processing version that will be used to get sample").Short('p').Default("0").Int16()
	no_date = app.Flag("no_date", "The processing version that will be used to get sample").Short('n').Default("true").Bool()

	pull = app.Command("pull", "Bring cloud database to local")
	sample = app.Command("sample", "Mount a sample collection")

	export = app.Command("export", "Mount an exportable collection")
	exportKind = export.Arg("kind", "What export do you want").Enum(exportables.Kinds...)
	exportPath = export.Flag("output", "Output path").Short('o').String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case pull.FullCommand():
			fmt.Print("--- Start Local Pull ---")
			dataset_to_local.TransferUsers()
			dataset_to_local.TransferTweets()
		case sample.FullCommand():
			select_sample.Run(*processing_version)
		case export.FullCommand():
			exportables.Run(*exportKind, *exportPath, *no_date)
	}
}