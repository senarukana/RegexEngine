package regex

import (
	"fmt"
)

var (
	StateID = -1
)

type State struct {
	StateId    int
	EpsilonSet StateSet
	Transition map[rune]*State
	Final      bool
}

func NewState() *State {
	StateID += 1
	return &State{
		StateId:    StateID,
		Transition: make(map[rune]*State),
		Final:      true,
		EpsilonSet: NewStateSet(),
	}
}

func (s *State) String() string {
	return fmt.Sprintf("--------STATE: %d , Final = %v--------", s.StateId, s.Final)
}

type StateSet map[int]*State

func NewStateSet() StateSet {
	return make(map[int]*State)
}

func (s StateSet) Insert(states ...*State) {
	for _, state := range states {
		s[state.StateId] = state
	}
}

func (s StateSet) Exists(state *State) bool {
	if _, ok := s[state.StateId]; ok {
		return true
	} else {
		return false
	}
}
