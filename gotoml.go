//gotoml : Golang package containing miscellaneous functions
//for Machine Learning (#deep-Learning, #PSO, #metaheuristic),
//Reinforcement Learning (#MDP #Q-learning #deep-Q-Networks)

package gotoml

import (
	"math/rand"
	"time"
)

// Errxxxx constants lists all errors
const (
	ErrGridWrongDimensions = Error("Wrong dimensions for Grid creation [height or width <= 0]")
	ErrEpsilonOutOfRange   = Error("epsilon greedy must be >= 0 & <= 1")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

}

//NewGrid creates a grid based on height and width
func NewGrid(p Params, node Noder) (g Grid, err error) {
	height := int(p["Height"])
	width := int(p["Width"])
	if height <= 0 || width <= 0 {
		return Grid{}, ErrGridWrongDimensions
	}
	g.Nodes = make([][]Noder, height, height)
	for h := range g.Nodes {
		g.Nodes[h] = make([]Noder, width, width)
	}
	g.Params = p
	return g, nil
}
