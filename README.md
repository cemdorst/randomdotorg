# randomdotorg
Golang functions to roll random dice using random.org API

Usage:

<pre>
package main

import (
	"fmt"

	"github.com/cemdorst/randomdotorg"
)

//RANDOMDOTORG_APIKEY env var must be set:
//export RANDOMDOTORG_APIKEY="xxxx-xxxx-xxxx-xxxx-xxxx"

func main() {

  //Roll 5 D10 dice
	d := randomdotorg.RollDice("5", "10")
	fmt.Println(d)

}
</pre>
