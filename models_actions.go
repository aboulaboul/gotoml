package gotoml

import (
	"fmt"
	"log"
	"math/rand"
)

// Actioner interface
type Actioner interface {
	GetActionsQv() interface{}
	GetActionQv(aid int) float64
	PutActionQv(aid int, v float64)
	TakeAction(epsilG float64) (aid int)
	Print()
}

// ActionsQv4 4 directions
type ActionsQv4 [4]float64

// GetActionsQv to get Actions Quality Value list
func (a *ActionsQv4) GetActionsQv() interface{} {
	return a
}

// GetActionQv to get specific action Quality Value
func (a *ActionsQv4) GetActionQv(aid int) float64 {
	return a[aid]
}

// PutActionQv to put specific action Quality Value
func (a *ActionsQv4) PutActionQv(aid int, v float64) {
	a[aid] = v
}

//TakeAction on a state based on epsilon
//epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
func (a *ActionsQv4) TakeAction(epsilG float64) (aid int) {

	//set action space
	sl := len(a)

	// exploration or exploitation ?
	explor, err := chooseExplor(epsilG)
	if err != nil {
		log.Println(err)
	}

	// take random action
	if explor {
		aid = rand.Intn(sl)
		return aid
	}
	// take best action
	valRef := a[0]
	for i := 0; i < sl; i++ {
		if a[i] > valRef {
			valRef = a[i]
			aid = i
		}
	}
	return aid
}

// Print for Printer interface
func (a *ActionsQv4) Print() {
	fmt.Printf("%v", a)
}

// ActionsQv5 On Plot + 4 directions
type ActionsQv5 [5]float64

// GetActionsQv to get Actions Quality Value list
func (a *ActionsQv5) GetActionsQv() interface{} {
	return a
}

// GetActionQv to get specific action Quality Value
func (a *ActionsQv5) GetActionQv(aid int) float64 {
	return a[aid]
}

// PutActionQv to put specific action Quality Value
func (a *ActionsQv5) PutActionQv(aid int, v float64) {
	a[aid] = v
}

//TakeAction on a state based on epsilon
//epsilon between 0-1 1 for exploration max and 0 for exploitation max greedy)
func (a *ActionsQv5) TakeAction(epsilG float64) (aid int) {

	//set action space
	sl := len(a)

	// exploration or exploitation ?
	explor, err := chooseExplor(epsilG)
	if err != nil {
		log.Println(err)
	}

	// take random action
	if explor {
		aid = rand.Intn(sl)
		return aid
	}
	// take best action
	valRef := a[0]
	for i := 0; i < sl; i++ {
		if a[i] > valRef {
			valRef = a[i]
			aid = i
		}
	}
	return aid
}

// Print for Printer interface
func (a *ActionsQv5) Print() {
	fmt.Printf("%v", a)
}

func chooseExplor(epsilG float64) (explor bool, err error) {
	if epsilG < 0 || epsilG > 1 {
		return false, Error(ErrEpsilonOutOfRange)
	}
	if epsilG == 0 {
		explor = false
	} else if epsilG == 1 {
		explor = true
	} else {
		if float64(rand.Intn(100)/100) < epsilG {
			explor = true
		} else {
			explor = false
		}
	}
	return explor, nil
}
