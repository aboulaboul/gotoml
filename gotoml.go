//gotoml : Golang package containing miscellaneous functions
//for Machine Learning (#deep-Learning, #PSO, #metaheuristic),
//Reinforcement Learning (#MDP #Q-learning #deep-Q-Networks)

package gotoml

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {

	cell := Cell{Weight: 1}

	for i := 0; i < 100; i++ {
		log.Println(cell.TakeAction(1, 0))
	}

}
