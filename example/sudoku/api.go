package sudoku

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GET https://sudoku.com/api/getLevel/easy
// {
//  "answer": "success",
//  "message": "Level exist",
//  "desc": [
//    "489501020750000810000020594008090075500008000001003000160374082000005736003062450",
//    "489531627752649813316827594238496175547218369691753248165374982924185736873962451",
//    9,
//    3,
//    false
//  ]
// }
type Level struct {
	Answer  string        `json:"answer"`
	Message string        `json:"message"`
	Desc    []interface{} `json:"desc"`
}

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
	Expert Difficulty = "expert"
)

func GetLevel(d Difficulty) (*Level, error) {
	resp, err := http.Get("https://sudoku.com/api/getLevel/" + string(d))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var l Level
	err = json.Unmarshal(bytes, &l)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func GetPair() (*Value, *Value, error) {
	level, err := GetLevel(Expert)
	if err != nil {
		return nil, nil, err
	}

	var a, b Value
	i := 0
	for x := range a {
		for y := range a[x] {
			a[x][y] = int(level.Desc[0].(string)[i] - '0')
			i++
		}
	}

	i = 0
	for x := range b {
		for y := range b[x] {
			b[x][y] = int(level.Desc[1].(string)[i] - '0')
			i++
		}
	}

	return &a, &b, nil
}
