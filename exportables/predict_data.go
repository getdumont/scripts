package exportables

import (
	"strconv"
	"strings"
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	. "github.com/getdumont/scripts/utilities"
)

var (
	DASS_QUESTIONS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
)

type arrayOfFloat []float64

type Dataset struct {
	Data []arrayOfFloat `json:"data"`
	Target []int `json:"target"`
	TargetNames []string `json:"target_names"`
	FeatureNames []string `json:"feature_names"`
}

func indexOf(slice []string, item string) int {
    for i, _ := range slice {
        if slice[i] == item {
            return i
        }
    }
    return -1
}

func flattenQuestions(questions []QuestionArray) []Question {
	return questions[0]
}

func PredictData() []byte {
	var _answers []AnswerResumed
	// head := []string{"sentiment", "magnitude", "verb", "pron", "propn"}
	head := []string{"sentiment"}
	dataset := Dataset{
		[]arrayOfFloat{},
		[]int{},
		DASS_QUESTIONS,
		[]string{},
	}

	// headAndIdx := map[string]int{
	// 	"VERB": 3,
	// 	"PRON": 4,
	// 	"PROPN": 5,
	// }

	answers, answerConnClose := ConnectAndGetCollection(LocalConfig, "answers")
	tweets, tweetConnClose := ConnectAndGetCollection(LocalConfig, "tweets")

	defer answerConnClose()
	defer tweetConnClose()

	answers.Pipe([]bson.M{{
		"$group": bson.M{
			"_id": "$to_tweet",
			"questions": bson.M{"$push": "$question"},
		},
	}}).All(&_answers)

	for _, a := range _answers {
		var t Tweet
		tweets.Find(bson.M{"_id": a.Id}).One(&t)
		data := []float64{}
		for range head {
			data = append(data, 0.0)
		}

		data[0] = t.CleanSentiment.Score
		// data[1] = t.CleanSentiment.Magnitude

		for _, word := range t.CleanTextTree {
			// idx := headAndIdx[word.Kind]
			// data[idx] = data[idx] + 1.0

			w := strings.ToLower(word.Value)
			headIdx := indexOf(head, w)

			if headIdx > -1 {
				data[headIdx] = data[headIdx] + 1.0
			} else {
				head = append(head, w)
				data = append(data, 1.0)
			}
		}

		for _, q := range flattenQuestions(a.Questions) {
			dataset.Data = append(dataset.Data, data)
			dataset.Target = append(dataset.Target, indexOf(DASS_QUESTIONS, strconv.Itoa(q.Index)))
		}
	}

	newData := []arrayOfFloat{}
	headLen := len(head)
	for _, d := range dataset.Data {
		missing := headLen - len(d)
		data := d
		for x := 0; x < missing; x++ {
			data = append(data, 0.0)
		}

		newData = append(newData, data)
	}

	dataset.FeatureNames = head
	dataset.Data = newData
	resp, _ := json.Marshal(dataset)

	return resp
}