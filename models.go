package gotoml

import (
	"fmt"
	"math/rand"
)

// Error for handling errors as constants
type Error string

func (e Error) Error() string {
	return string(e)
}

// Cell smallest entity of Grid world
type Cell interface {
	TakeAction(epsilG float64) (action int)
	GetActions() []float64
	GetAction(a int) float64
	PutAction(a int, v float64)
	GetRawVal() string
	GetWeight() float64
	PutRawVal(s string) bool
	PutWeight(f float64) bool
}

// Cell4Dir of a Grid world
type Cell4Dir struct {
	RawVal  string    //raw value in real world
	Weight  float64   //weight (reward) of State (cell)
	Actions []float64 //actions in x directions (N 0, E 1, S 2, O 3)
}

//TakeAction on a state based on epsilon and mode
//epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
func (c Cell4Dir) TakeAction(epsilG float64) (action int) {
	//set action space
	sl := len(c.Actions)

	// exploration or exploitation ?
	var explor bool
	if epsilG == 0 {
		explor = false
	} else if epsilG == 1 {
		explor = true
	} else {
		if float64(rand.Intn(1000)/1000) < epsilG {
			explor = true
		} else {
			explor = false
		}
	}

	// take action
	if explor {
		action = rand.Intn(sl)
		return action
	}
	valRef := c.Actions[0]
	for i := 0; i < sl; i++ {
		if c.Actions[i] > valRef {
			valRef = c.Actions[i]
			action = i
		}
	}
	return action
}

// GetActions to get actions list
func (c Cell4Dir) GetActions() []float64 {
	return c.Actions
}

// GetAction to get specific action value
func (c Cell4Dir) GetAction(a int) float64 {

	return c.Actions[a]

}

// PutAction to put specific action value
func (c Cell4Dir) PutAction(a int, v float64) {
	c.Actions[a] = v
}

// GetRawVal to get rax value
func (c Cell4Dir) GetRawVal() string {
	return c.RawVal
}

// GetWeight to get rax value
func (c Cell4Dir) GetWeight() float64 {
	return c.Weight
}

// PutRawVal to put rax value
func (c Cell4Dir) PutRawVal(s string) bool {
	c.RawVal = s
	return true

}

// PutWeight to put rax value
func (c Cell4Dir) PutWeight(f float64) bool {
	c.Weight = f
	return true
}

//Grid world contains states and Qtables
type Grid struct {
	Cells  [][]Cell
	Width  int
	Height int
}

// DisplayRawVal of grid
func (g Grid) DisplayRawVal() {
	for h := range g.Cells {
		for w := range g.Cells[h] {
			if g.Cells[h][w] != nil {
				fmt.Printf("%v", g.Cells[h][w].GetRawVal())
			} else {
				fmt.Printf("%v", ``)
			}

		}
		fmt.Printf("\n")
	}

}

// DisplayWeight of grid
func (g Grid) DisplayWeight() {
	for h := range g.Cells {
		for w := range g.Cells[h] {
			if g.Cells[h][w] != nil {
				fmt.Printf("%v", g.Cells[h][w].GetWeight())
			} else {
				fmt.Printf("%v", ``)
			}
		}
		fmt.Printf("\n")
	}

}

// GetCell from Grid
func (g Grid) GetCell(l, c int) (Cell, bool) {
	if !g.CheckInGrid(l, c) {
		return nil, false
	}
	if g.Cells[l][c] == nil {
		return nil, false
	}
	return g.Cells[l][c], true
}

// PutCell to Grid
func (g Grid) PutCell(l, c int, cell Cell) bool {
	if !g.CheckInGrid(l, c) {
		return false
	}
	switch cell.(type) {
	case Cell4Dir:
		g.Cells[l][c] = Cell4Dir{
			RawVal:  cell.GetRawVal(),
			Weight:  cell.GetWeight(),
			Actions: cell.GetActions(),
		}
		// init border actions
		if l == 0 {
			g.Cells[l][c].PutAction(0, -1)
		} else if l == g.Height-1 {
			g.Cells[l][c].PutAction(2, -1)
		}
		if c == 0 {
			g.Cells[l][c].PutAction(3, -1)
		} else if c == g.Width-1 {
			g.Cells[l][c].PutAction(1, -1)
		}
		return true
	default:
		return false
	}
}

// CheckInGrid checks if line & colomn coordinates  are in height & width
func (g Grid) CheckInGrid(l, c int) bool {
	if l < 0 || l >= g.Height {
		return false
	}
	if c < 0 || c >= g.Width {
		return false
	}
	return true
}
