package regex

import (
	"fmt"
)

type DFA struct {
	Id          int
	states      StateSet
	transitions map[rune]*DFA
	final       bool
}

func NewDFA(id int, states StateSet) *DFA {
	dfa := &DFA{
		Id:          id,
		states:      states,
		transitions: make(map[rune]*DFA),
	}

	for _, state := range states {
		if state.Final {
			dfa.final = true
			break
		}
	}

	return dfa
}

func (d *DFA) IsDeadState() bool {
	if d.final {
		return false
	}
	if len(d.transitions) == 0 {
		return true
	}
	for _, dfa := range d.transitions {
		if d != dfa {
			return false
		}
	}
	return true
}

func (d *DFA) PrettyPrint() {
	fmt.Printf("--------STATE: %d , Final = %v--------\n", d.Id, d.final)
	fmt.Printf("StateSet: { ")
	for id, _ := range d.states {
		fmt.Printf("%d ", id)
	}
	fmt.Println("}")
	for r, neighbor := range d.transitions {
		fmt.Printf("Transition: %c => %d\n", r, neighbor.Id)
	}
	fmt.Println()
}
