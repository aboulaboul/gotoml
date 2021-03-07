package gotoml

import (
	"fmt"
	"sync"
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
	GetActionQv(aid uint) float64
	PutActionQv(aid uint, v float64)
	TakeAction(epsilG float64) (aid int)
	GetStateQv() float64
	PutStateQv(v float64)
	Print()
}

// Node type of a Grid world
type Node struct {
	Mutx      *sync.RWMutex
	RawVal    string   //raw value in 'real' world
	Reward    float64  //Reward of Noder should be >-100 & <100
	StateQv   float64  //Noder Q value
	Visited   int      //number of updates
	ActionsQv Actioner //actions Q value in x directions
}

// GetRawVal to get raw value from node
func (c *Node) GetRawVal() string {
	return c.RawVal
}

// PutRawVal to put raw value on node
func (c *Node) PutRawVal(s string) {
	c.RawVal = s
}

// GetReward to get Reward value  from node
func (c *Node) GetReward() float64 {
	return c.Reward
}

// PutReward to put Reward value on node
func (c *Node) PutReward(f float64) {
	c.Reward = f
}

// GetVisited to get number of visits from node
func (c *Node) GetVisited() int {
	return int(c.Visited)
}

// PutVisited to put number of visits on node
func (c Node) PutVisited(v int) {
	c.Visited = v
}

// GetActionsQv to get Actions Quality Value list from node
func (c *Node) GetActionsQv() interface{} {
	return c.ActionsQv.GetActionsQv()
}

// GetActionQv to get specific Quality Value from action node
func (c *Node) GetActionQv(aid uint) float64 {
	return c.ActionsQv.GetActionQv(aid)
}

// PutActionQv to put specific Quality Value on action node
func (c *Node) PutActionQv(aid uint, v float64) {
	c.ActionsQv.PutActionQv(aid, v)
}

// GetStateQv to get State Quality Value from node
func (c *Node) GetStateQv() float64 {
	return c.StateQv
}

// PutStateQv to put State Quality Value on node
func (c *Node) PutStateQv(v float64) {
	c.StateQv = v
}

// Print for Printer interface
func (c *Node) Print() {
	fmt.Printf("%v", c)
}

// TakeAction from a node based on epsilon
// epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
func (c *Node) TakeAction(epsilG float64) (aid int) {
	return c.ActionsQv.TakeAction(epsilG)
}
