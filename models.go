package gotoml

import (
	"fmt"
)

// Error struct for handling errors as constants
type Error string

func (e Error) Error() string {
	return string(e)
}

// Grid world struct contains states and Qtables
type Grid struct {
	Nodes  [][]Noder
	Params Params
}

// Params type
// Height
// Width
// EpsilonG for epsilon greedy
// Alpha for learning rate (0.1)
// Gamma for discount rate (0.9)
type Params map[string]float64

// DisplayRawVal of grid
func (g *Grid) DisplayRawVal() {
	for h := range g.Nodes {
		for w := range g.Nodes[h] {
			if g.Nodes[h][w] != nil {
				fmt.Printf("|%v", g.Nodes[h][w].GetRawVal())
			} else {
				fmt.Printf("|%v", ``)
			}
		}
		fmt.Printf("|\n")
	}
}

// DisplayReward of grid
func (g *Grid) DisplayReward() {
	for h := range g.Nodes {
		for w := range g.Nodes[h] {
			if g.Nodes[h][w] != nil {
				fmt.Printf("|%05.2f", g.Nodes[h][w].GetReward())
			} else {
				fmt.Printf("|%v", `     `)
			}
		}
		fmt.Printf("|\n")
	}
}

// DisplayVisited of grid
func (g *Grid) DisplayVisited() {
	for h := range g.Nodes {
		for w := range g.Nodes[h] {
			if g.Nodes[h][w] != nil {
				fmt.Printf("|%03d", g.Nodes[h][w].GetVisited())
			} else {
				fmt.Printf("|%v", `   `)
			}
		}
		fmt.Printf("|\n")
	}
}

// DisplayStateQv of grid
func (g *Grid) DisplayStateQv() {
	for h := range g.Nodes {
		for w := range g.Nodes[h] {
			if g.Nodes[h][w] != nil {
				fmt.Printf("|%05.2f", g.Nodes[h][w].GetStateQv())
			} else {
				fmt.Printf("|%v", `     `)
			}
		}
		fmt.Printf("|\n")
	}
}

// GetNode from Grid
func (g *Grid) GetNode(l, c int) Noder {
	if !g.checkInGrid(l, c) {
		return nil
	}
	if g.Nodes[l][c] == nil {
		return nil
	}
	return g.Nodes[l][c]
}

// PutNode to Grid
func (g *Grid) PutNode(l, c int, node Noder) bool {
	if !g.checkInGrid(l, c) {
		return false
	}
	Height := int(g.Params["Height"])
	Width := int(g.Params["Width"])
	switch node.GetActionsQv().(type) {
	case *ActionsQv4:
		// init border ActionsQv
		if l == 0 {
			node.PutActionQv(0, -1)
		} else if l == Height-1 {
			node.PutActionQv(2, -1)
		}
		if c == 0 {
			node.PutActionQv(3, -1)
		} else if c == Width-1 {
			node.PutActionQv(1, -1)
		}
	case *ActionsQv5:
		// init border ActionsQv
		if l == 0 {
			node.PutActionQv(1, -1)
		} else if l == Height-1 {
			node.PutActionQv(3, -1)
		}
		if c == 0 {
			node.PutActionQv(4, -1)
		} else if c == Width-1 {
			node.PutActionQv(2, -1)
		}
	default:
		return false
	}
	g.Nodes[l][c] = node
	return true
}

// checkInGrid checks if line & colomn coordinates are in height & width
func (g *Grid) checkInGrid(l, c int) bool {
	Height := int(g.Params["Height"])
	Width := int(g.Params["Width"])
	if l < 0 || l >= Height {
		return false
	}
	if c < 0 || c >= Width {
		return false
	}
	return true
}

// Move from a Node to another and update Q values
func (g *Grid) Move(fromL, fromC int, aid int) (toL, toC int, ok bool) {
	node := g.GetNode(fromL, fromC)
	if node == nil {
		return toL, toC, false
	}
	switch node.GetActionsQv().(type) {
	case *ActionsQv4:
		switch aid {
		case 0:
			toL, toC = fromL-1, fromC
		case 1:
			toL, toC = fromL, fromC+1
		case 2:
			toL, toC = fromL+1, fromC
		case 3:
			toL, toC = fromL, fromC-1
		}
	case *ActionsQv5:
		switch aid {
		case 0:
			toL, toC = fromL, fromC
		case 1:
			toL, toC = fromL-1, fromC
		case 2:
			toL, toC = fromL, fromC+1
		case 3:
			toL, toC = fromL+1, fromC
		case 4:
			toL, toC = fromL, fromC-1
		}
	default:
		return toL, toC, false
	}
	node1 := g.GetNode(toL, toC)
	if node1 == nil {
		return toL, toC, false
	}
	//update Q values
	Alpha := g.Params["Alpha"]
	Gamma := g.Params["Gamma"]

	q := (1-Alpha)*node.GetActionQv(aid) + Alpha*(node1.GetReward()+Gamma*node1.GetStateQv())
	node.PutActionQv(aid, q)
	node.PutVisited(node.GetVisited() + 1)
	if q > node.GetStateQv() {
		node.PutStateQv(q)
	}
	ok = g.PutNode(fromL, fromC, node)
	if !ok {
		return toL, toC, false
	}
	return toL, toC, ok
}
