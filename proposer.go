package gaxsos

type proposer struct {
    id          int // the identifier for each proposer

    proposeOrd  int // ordering number for the proposal
    proposeVal  string // value that intended to be proposed

    acceptors   []chan message // channels for sending message to acceptors
    receives    chan message // channel for receiving messages from acceptors
    learners    []chan message // channels for sending the choosen value to learners
}
