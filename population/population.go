package pop

import (
	"fmt"
	"genetic-evolution-sim/grid"
	"genetic-evolution-sim/neural_net"
)

type individual struct {
	alive     bool
	index     int
	loc       grid.Loc
	birthLoc  grid.Loc
	genome    *neural_net.Genome
	neuralNet *neural_net.NeuralNet
	age       uint
}

type Pop []individual

func InitPop(population int) *Pop {
	pop := make(Pop, population)
	return &pop
}

// Spawn a new individual and store it in the pop array
func (pop *Pop) InitIndividual(index int, loc grid.Loc, genome *neural_net.Genome) {
	(*pop)[index] = individual{
		alive:    true,
		index:    index,
		loc:      loc,
		birthLoc: loc,
		age:      0,
		genome:   genome,
	}
}

func (pop *Pop) Print() {
	for _, indiv := range *pop {
		fmt.Println("Alive:", indiv.alive, " Age:", indiv.age, "Loc:", indiv.loc, " ")
	}
}
