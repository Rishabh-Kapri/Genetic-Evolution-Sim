package main

import (
	"fmt"
	"genetic-evolution-sim/simulator"
	"math/rand"
	"time"
)

// Thing needed
// genome sequence handler, parse and returns action
// feed forward function
// neural network
// sensors

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(0x7fff) // 16 bit
	// @TODO
	// take config file from command line, if not provided use conf.json
	confFilePath := "conf.json"
	sim.Run(confFilePath)
}
