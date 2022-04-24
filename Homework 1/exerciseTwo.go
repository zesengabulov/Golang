package main

import (
	"fmt"
)

func main(){
	var pebbleWeight float64 = 0.1  // may be it was supposed to use const pebbleWeight float64 = 0.1 instead since we didn't change its interrior value
	var rockWeight float64 = 1.2    // may be it was supposed to use const rockWeight float64 = 1.2  instead since we didn't change its interrior value
	var boulderWeight float64 = 502.4   // may be it was supposed to use const boulderWeight float64 = 502.4  instead since we didn't change its interrior value
	var totalWeight float64 = pebbleWeight + rockWeight + boulderWeight
	fmt.Println(totalWeight)
}