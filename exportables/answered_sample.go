package exportables

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

type questionArray []Question

type answerResumed struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Questions []questionArray `bson:"questions" json:"questions"`
}

func AnsweredSample() []byte {
	var _answers []answerResumed
	answers, answerConnClose := ConnectAndGetCollection(LocalConfig, "answers")

	defer answerConnClose()
	answers.Pipe([]bson.M{{
		"$group": bson.M{
			"_id": "$to_tweet",
			"questions": bson.M{"$push": "$question"},
		},
	}}).All(&_answers)

	resp, _ := json.Marshal(_answers)
	return resp
}