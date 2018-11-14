package exportables

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

func AnsweredSample() []byte {
	var _answers []AnswerResumed
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