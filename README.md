# randomdotorg
Golang functions to roll random dice using random.org API

Usage:

<pre>
package main

import (
	"fmt"

	"github.com/cemdorst/randomdotorg"
)

func main() {

	randomdotorg.APIKey = "xxxx-xxxx-xxxx-xxxx-xxxx"

  //Roll 5 D10 dice
	d := randomdotorg.RollDice("5", "10")
	fmt.Println(d)

}
</pre>
