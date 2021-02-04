package randomdotorg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	jsonrpc2 "github.com/AdamSLevy/jsonrpc2/v14"
	"github.com/cemdorst/randomdotorg/model"
)

//APIKey to access api.random.org
var APIKey string

//API base URL
const apiurl string = "https://api.random.org/json-rpc/2/invoke"

//Mininum sides on a dice
const minsides = "2"
const minresult = "1"

//ErrDiceSides states minimum Dice sides
var ErrDiceSides = errors.New("Dice must have more than" + minsides)

//ErrJSON related to JSON operations
var ErrJSON = errors.New("JSON Marshal/Unmarshal error")

//ErrHTTP related to HTTP operations
var ErrHTTP = errors.New("HTTP related error")

//RollDice will throw (n) dice with (s) sides
func RollDice(n string, s string) (result []int) {
	if s < minsides {
		fmt.Println(ErrDiceSides)
		return result
	}

	params := []string{APIKey, n, minresult, s}

	req, err := json.Marshal(jsonrpc2.Request{Method: "generateIntegers",
		Params: params,
		ID:     0,
	})

	if err != nil {
		fmt.Println(ErrJSON.Error())
	}

	httpRes, err := http.Post(apiurl, "application/json",
		bytes.NewReader(req))

	if err != nil {
		fmt.Printf(ErrHTTP.Error())
	}

	respBytes, err := ioutil.ReadAll(httpRes.Body)

	if err != nil {
		fmt.Printf(ErrHTTP.Error())
	}

	var diceRoll model.Dice

	err = json.Unmarshal(respBytes, &diceRoll)
	if err != nil {
		fmt.Printf(ErrJSON.Error())
	}

	result = diceRoll.Result.Random.Data
	return
}
