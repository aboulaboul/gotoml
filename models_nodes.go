package gotoml

import (
	"fmt"
)

// Noder interface
type Noder interface {
	GetRawVal() string
	PutRawVal(s string)
	GetReward() float64
	PutReward(f float64)
	GetVisited() int
	PutVisited(v int)
	GetActionsQv() interface{}
	GetActionQv(aid int) float64
	PutActionQv(aid int, v float64)
	TakeAction(epsilG float64) (aid int)
	GetStateQv() float64
	PutStateQv(v float64)
	Print()
}

// Node type of a Grid world
type Node struct {
	RawVal    string   //raw value in 'real' world
	Reward    float64  //Reward of Noder should be >-100 & <100
	StateQv   float64  //Noder Q value
	Visited   int      //number of updates
	ActionsQv Actioner //actions Q value in x directions
}

// GetRawVal to get raw value from node
func (n *Node) GetRawVal() string {
	return n.RawVal
}

// PutRawVal to put raw value on node
func (n *Node) PutRawVal(s string) {
	n.RawVal = s
}

// GetReward to get Reward value  from node
func (n *Node) GetReward() float64 {
	return n.Reward
}

// PutReward to put Reward value on node
func (n *Node) PutReward(f float64) {
	n.Reward = f
}

// GetVisited to get number of visits from node
func (n *Node) GetVisited() int {
	return int(n.Visited)
}

// PutVisited to put number of visits on node
func (n *Node) PutVisited(v int) {
	n.Visited = v
}

// GetActionsQv to get Actions Quality Value list from node
func (n *Node) GetActionsQv() interface{} {
	return n.ActionsQv.GetActionsQv()
}

// GetActionQv to get specific Quality Value from action node
func (n *Node) GetActionQv(aid int) float64 {
	return n.ActionsQv.GetActionQv(aid)
}

// PutActionQv to put specific Quality Value on action node
func (n *Node) PutActionQv(aid int, v float64) {
	n.ActionsQv.PutActionQv(aid, v)
}

// GetStateQv to get State Quality Value from node
func (n *Node) GetStateQv() float64 {
	return n.StateQv
}

// PutStateQv to put State Quality Value on node
func (n *Node) PutStateQv(v float64) {
	n.StateQv = v
}

// Print for Printer interface
func (n *Node) Print() {
	fmt.Printf("%v", n)
}

// TakeAction from a node based on epsilon
// epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
func (n *Node) TakeAction(epsilG float64) (aid int) {
	return n.ActionsQv.TakeAction(epsilG)
}
