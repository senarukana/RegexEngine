package regex

import (
	"fmt"
)

type NFA struct {
	Start *State
	End   *State
}

func NewNFA(start, end *State) *NFA {
	return &NFA{
		Start: start,
		End:   end,
	}
}

func (n *NFA) PrettyPrint() {
	var (
		stateQueue []*State
		curState   *State
		stateSet   = NewStateSet()
	)
	stateQueue = append(stateQueue, n.Start)
	stateSet.Insert(n.Start)
	for len(stateQueue) > 0 {
		curState, stateQueue = stateQueue[0], stateQueue[1:]
		fmt.Println(curState)

		fmt.Printf("EpsilonSet:{ ")
		for _, state := range curState.EpsilonSet {
			fmt.Printf("%d ", state.StateId)
			if !stateSet.Exists(state) {
				stateQueue = append(stateQueue, state)
				stateSet.Insert(state)
			}
		}
		fmt.Println("}")
		for r, state := range curState.Transition {
			fmt.Printf("Transition: %c => %d\n", r, state.StateId)
			if !stateSet.Exists(state) {
				stateQueue = append(stateQueue, state)
				stateSet.Insert(state)
			}
		}
		fmt.Println()
	}
}

type NFASet []*NFA

func (n NFASet) Pop() (*NFA, NFASet) {
	return n[len(n)-1], n[:len(n)-1]
}

func (n NFASet) Push(nfa *NFA) NFASet {
	return append(n, nfa)
}
