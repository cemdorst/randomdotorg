package randomdotorg

import (
	"os"
	"strings"
	"testing"
)

func Init() {
	APIKey = os.Getenv("RANDOMDOTORG_APIKEY")
}

func TestRollDice(t *testing.T) {
	Init()
	result := RollDice("4", "6")
	if len(result) != 4 {
		t.Errorf("RollDice result length should be 4, but returned: %d", len(result))
	}
	for dice := range result {
		if dice > 6 {
			t.Errorf("RollDice should be <= 6, but was: %d", result)
		}
	}
}

func TestRollDiceEmpty(t *testing.T) {
	Init()
	result := RollDice("2", "6")
	if len(result) == 0 {
		t.Errorf("RollDice result length should be 1, but returned: %d", len(result))
	}

}
