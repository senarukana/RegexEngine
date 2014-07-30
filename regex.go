package regex

import (
	"errors"
	"fmt"
	"reflect"
)

var debug = true

type Regex struct {
	DFATable    []*DFA
	NFAPoint    *NFA
	runes       []rune
	nextStateID int
}

func NewRegex(source string) (*Regex, error) {
	// TODO: remove the annoying global StateID
	StateID = -1
	var nfaSet NFASet
	lex := NewLex(source)
	parser := NewParser(lex)
	handlers, err := parser.Parse()

	if err != nil {
		return nil, errors.New("Invalid regex expression")
	}

	for _, handler := range handlers {
		nfaSet = handler.Accept(nfaSet)
	}

	if len(nfaSet) != 1 {
		return nil, errors.New("Invalid regex expression")
	}
	regex := &Regex{
		runes:    []rune(source),
		NFAPoint: nfaSet[0],
	}
	regex.DFATable = regex.nfa2dfa(nfaSet[0].Start)
	if debug {
		regex.printNFA()
		regex.printDFATable()
	}
	return regex, nil
}

func (r *Regex) Match(s string) bool {
	curDFA := r.DFATable[0]
	for _, r := range s {
		if nextDFA, ok := curDFA.transitions['?']; ok {
			curDFA = nextDFA
		} else if nextDFA, ok := curDFA.transitions[r]; ok {
			curDFA = nextDFA
		} else {
			curDFA = nil
			break
		}
	}
	if curDFA == nil || !curDFA.final {
		return false
	} else {
		return true
	}
}

func (regex *Regex) epsilonStates(startStates StateSet) StateSet {
	var (
		statesQueue []*State
		state       *State
	)
	states := NewStateSet()
	for _, s := range startStates {
		statesQueue = append(statesQueue, s)
		states.Insert(s)
	}

	for len(statesQueue) > 0 {
		state, statesQueue = statesQueue[0], statesQueue[1:]

		for _, epsilonState := range state.EpsilonSet {
			// fmt.Println(epsilonState)
			if !states.Exists(epsilonState) {
				statesQueue = append(statesQueue, epsilonState)
				states.Insert(epsilonState)
			}
		}
	}
	return states
}

func (regex *Regex) transition(states StateSet, r rune) StateSet {
	res := NewStateSet()
	for _, state := range states {
		if nextState, ok := state.Transition[r]; ok {
			if !res.Exists(nextState) {
				res.Insert(nextState)
			}
		}
	}
	return res
}

func (regex *Regex) printNFA() {
	fmt.Println("----------DFA-TABLE----------")
	regex.NFAPoint.PrettyPrint()
}

func (regex *Regex) printDFATable() {
	fmt.Println("----------DFA-TABLE----------")
	for _, dfa := range regex.DFATable {
		dfa.PrettyPrint()
	}
}

func (regex *Regex) nfa2dfa(startState *State) (dfaTable []*DFA) {
	var (
		dfaQueue []*DFA
		curDFA   *DFA
		states   StateSet
	)
	startSet := NewStateSet()
	startSet.Insert(startState)
	dfaStartSet := regex.epsilonStates(startSet)
	startDFA := NewDFA(regex.nextStateID, dfaStartSet)
	regex.nextStateID += 1

	dfaQueue = append(dfaQueue, startDFA)
	dfaTable = append(dfaTable, startDFA)
	for len(dfaQueue) > 0 {
		curDFA, dfaQueue = dfaQueue[0], dfaQueue[1:]
		for _, r := range regex.runes {
			if transitionStates := regex.transition(curDFA.states, r); len(transitionStates) != 0 {
				states = regex.epsilonStates(transitionStates)
				// check if transitionStates is in dfaTable
				found := false
				for _, dfa := range dfaTable {
					if reflect.DeepEqual(dfa.states, states) {
						found = true
						curDFA.transitions[r] = dfa
						break
					}
				}
				if !found {
					nextDFA := NewDFA(regex.nextStateID, states)
					regex.nextStateID += 1
					curDFA.transitions[r] = nextDFA

					dfaTable = append(dfaTable, nextDFA)
					dfaQueue = append(dfaQueue, nextDFA)
				}
			}
		}
	}
	return dfaTable
}
