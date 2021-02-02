package model

//Dice represents the tye of dice being played
//JSON has nested maps like: map[string]map[string]map[string]json.RawMessage
type Dice struct {
	Result struct {
		Random struct {
			Data []int `json:"data"`
		} `json:"random"`
	} `json:"result"`
}
