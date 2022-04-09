package neural_net

import (
	"math/rand"
)

type Gene struct {
	sourceType   int
	sourceNum    int // 7 bit
	sinkType     int
	sinkNum      int     // 7 bit
	weight       int     // 16 bit
	weightScaled float32 // -4.0 to 4.0
}

type Genome []Gene

type Neuron struct {
	output float32
	// driven bool
}

type NeuralNet struct {
	connections []Gene
	neurons     []Neuron
}

func MakeRandomGene() *Gene {
	lower := -32678 // 16 bit lower
	upper := 32767  // 16 bit upper
	rng := upper - lower
	weight := rand.Intn(rng) + lower
	gene := Gene{
		sourceType:   rand.Int() & 1,
		sourceNum:    rand.Intn(0x7fff),
		sinkType:     rand.Int() & 1,
		sinkNum:      rand.Intn(0x7fff),
		weight:       weight,
		weightScaled: float32(weight) / 8192.0,
	}
	return &gene
}

func MakeRandomGenome(maxGenomeLen int, minGenomeLen int) *Genome {
	rng := maxGenomeLen - minGenomeLen + 1
	length := rand.Intn(rng) + minGenomeLen
	genome := make(Genome, 0)
	for length > 0 {
		genome = append(genome, *MakeRandomGene())
		length--
	}
	return &genome
}
