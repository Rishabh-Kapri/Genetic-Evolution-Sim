package sensors

type Sensors int
type Actions int

// Neuron Sources (Sensors)
// Neuron Sinks (Actions)
const (
	DIST_X         Sensors = iota // distance from left side
	DIST_Y                        // distance from bottom
	NEAREST_DIST_X                // individual x distance from nearset side
	NEAREST_DIST_Y                // individual y distance from nearset side
	NEAREST_DIST                  // individual distance from nearset side
	POPULATION                    // population density in the neighbourhood
	POPULATION_FWD                // population density in forward-reverse direction
	POPULATION_LR                 // population density in left-right direction
	AGE                           // individual's age
)

const (
	MOVE_X       Actions = iota // move in x direction
	MOVE_Y                      // move in y direction
	MOVE_FORWARD                // continue last direction
	MOVE_RANDOM                 // move everyone randomly
	MOVE_LEFT                   // move everyone to the left
	MOVE_RIGHT                  // move everyone to the right
	MOVE_UP                     // move everyone up
	MOVE_DOWN                   // move everyone down
	MOVE_REVERSE                // move everyone reverse
)
