package randomdotorg

import (
	"bytes"
	"encoding/json"
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

//RollDice will throw (n) dice with (s) sides
func RollDice(n string, s string) []int {

	i := "1" //Dice sides go from i to s
	params := []string{APIKey, n, i, s}

	req, _ := json.Marshal(jsonrpc2.Request{Method: "generateIntegers",
		Params: params,
		ID:     0,
	})

	httpRes, err := http.Post(apiurl, "application/json",
		bytes.NewReader(req))

	if err != nil {
		fmt.Printf("Error:%v", err.Error())
	}

	respBytes, err := ioutil.ReadAll(httpRes.Body)

	if err != nil {
		fmt.Printf("Error:%v", err.Error())
	}

	var diceRoll model.Dice

	err = json.Unmarshal(respBytes, &diceRoll)
	if err != nil {
		fmt.Printf("Error:%v", err.Error())
	}

	return diceRoll.Result.Random.Data
}
