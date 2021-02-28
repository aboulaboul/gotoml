package gotoml

import (
	"math/rand"
)

//ActionMode mode 0=4 directions N(orth) E(st) S(outh) W(est) mode 1=mode 0+NE SE Sw Nw ...
type ActionMode int

//Cell of a grid world
type Cell struct {
	Weight  float64   //weight (reward) of State (cell)
	Actions []float64 //actions in x directions (N 0, E 1, S 2, O 3, NE 4, SE 5, SO 6, NO 7)
}

//TakeAction on a state based on epsilon and mode
//epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
//mode 0 NESW mode 1 0+NESESWNW
func (c *Cell) TakeAction(epsilG float64, mode ActionMode) (action int) {
	//set action space
	sl := 4
	if mode != 0 {
		sl = len(c.Actions)
	}

	//choose between exploration or exploitation
	if float64(rand.Intn(1000)/1000) < epsilG {
		action = rand.Intn(sl)
		return action
	}
	ref := float64(0)
	for i := 0; i < sl; i++ {
		if c.Actions[i] > ref {
			ref = c.Actions[i]
			action = i
		}
	}
	return action
}

//Grid world contains states and Qtables
type Grid [][]Cell
