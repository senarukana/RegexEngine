package regex

import (
// "fmt"
)

type Handler interface {
	Accept(nfaSet NFASet) NFASet
}

type CharHandler struct {
	*Token
}

type ReplHandler struct {
	*Token
}

type ConsHandler struct {
	*Token
}

type QMarkHandler struct {
	*Token
}

type AltHandler struct {
	*Token
}

func NewCharHandler(tok *Token) *CharHandler {
	return &CharHandler{
		Token: tok,
	}
}

func NewQMarkHandler(tok *Token) *QMarkHandler {
	return &QMarkHandler{
		Token: tok,
	}
}

func NewConsHandler(tok *Token) *ConsHandler {
	return &ConsHandler{
		Token: tok,
	}
}

func NewReplHandler(tok *Token) *ReplHandler {
	return &ReplHandler{
		Token: tok,
	}
}

func NewAltHandler(tok *Token) *AltHandler {
	return &AltHandler{
		Token: tok,
	}
}

func (h *CharHandler) Accept(n NFASet) NFASet {
	// fmt.Println("CHAR")
	s0 := NewState()
	s1 := NewState()
	s0.Final = false
	s0.Transition[h.Val] = s1
	nfa := NewNFA(s0, s1)
	return n.Push(nfa)
}

func (h *QMarkHandler) Accept(n NFASet) NFASet {
	// fmt.Println("QMARK")
	s0 := NewState()
	s1 := NewState()
	s0.Transition[h.Val] = s1
	s0.EpsilonSet.Insert(s1)
	nfa := NewNFA(s0, s1)
	return n.Push(nfa)
}

func (h *ReplHandler) Accept(n NFASet) NFASet {
	// fmt.Println("REPL")
	nfa, n := n.Pop()
	s0 := NewState()
	s1 := NewState()
	if h.Token.Type == STAR {
		s0.EpsilonSet.Insert(s1)
	}
	s0.EpsilonSet.Insert(nfa.Start)
	nfa.End.EpsilonSet.Insert(nfa.Start, s1)
	s0.Final = false
	newNfa := NewNFA(s0, s1)
	return n.Push(newNfa)
}

func (h *ConsHandler) Accept(n NFASet) NFASet {
	// fmt.Println("CONS")
	n1, n := n.Pop()
	n0, n := n.Pop()
	n0.End.EpsilonSet.Insert(n1.Start)
	n0.End.Final = false
	nfa := NewNFA(n0.Start, n1.End)
	return n.Push(nfa)
}

func (h *AltHandler) Accept(n NFASet) NFASet {
	n0, n := n.Pop()
	n1, n := n.Pop()
	s0 := NewState()
	s1 := NewState()
	s0.EpsilonSet.Insert(n0.Start, n1.Start)
	n0.End.Final = false
	n0.End.EpsilonSet.Insert(s1)
	n1.End.Final = false
	n1.End.EpsilonSet.Insert(s1)
	nfa := NewNFA(s0, s1)
	return n.Push(nfa)
}
