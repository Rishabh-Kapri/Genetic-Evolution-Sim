package sim

import (
	"fmt"
	"genetic-evolution-sim/config"
	"genetic-evolution-sim/grid"
	"genetic-evolution-sim/neural_net"
	pop "genetic-evolution-sim/population"
)

// import pop "genetic-evolution-sim/population"

func Run(confFilePath string) {
	conf := config.LoadConfig(confFilePath)
	grid := grid.CreateGrid(conf.SizeX, conf.SizeY)
	pop := pop.InitPop(conf.Population)
	generation := 0
	fmt.Println(generation)
	pop.Print()
	initGeneration0(conf, pop, grid)
	pop.Print()
}

func initGeneration0(conf config.Config, pop *pop.Pop, grid *grid.Grid) {
	// @TODO
	// reset grid
	for i := 0; i < conf.Population; i++ {
		genome := neural_net.MakeRandomGenome(conf.MaxGenomeLength, conf.MinGenomeLength)
		pop.InitIndividual(i, grid.FindRandomEmptyLoc(), genome)
	}
}
