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
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {

}

//NewGrid creates a grid based on height and width
func NewGrid(height, width int, cell Cell) (g Grid, err error) {
	if height <= 0 || width <= 0 {
		return Grid{}, ErrGridWrongDimensions
	}
	g.Cells = make([][]Cell, height)
	for h := range g.Cells {
		g.Cells[h] = make([]Cell, width)
	}
	g.Height = height
	g.Width = width
	return g, nil
}
