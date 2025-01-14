package gaxsos

import "fmt"

type proposer struct {
    id          int // the identifier for each proposer

    proposalOrd  int // ordering number for the proposal
    proposalVal  string // value that intended to be proposed

    acceptors   []chan message // channels for sending message to acceptors
    receives    chan message // channel for receiving messages from acceptors
    learners    []chan message // channels for sending the choosen value to learners
}

func NewProposer(id int, val string, receives chan message, acceptors, learners []chan message) *proposer {
    p := proposer{}

    p.id = id
    p.proposalOrd = 0
    p.proposalVal = val
    p.receives = receives
    p.acceptors = acceptors
    p.learners = learners

    return &p
}

func (p *proposer) Run() {
    fmt.Printf("Proposer %v: started\n", p.id)

    decide := false

    for !decide {
        // phase-1(Prepare-Promise)
        p.prepare()

        responded := make(map[int]bool) // to track which acceptors have responded

        max := p.proposalOrd

        // wait for a majority of Promise responses
        for len(responded) < len(p.acceptors) / 2 + 1 { // Majority = more than half of acceptors
            msg := <- p.receives

            switch msg.status {
            case Promise:
                responded[msg.from] = true // mark the acceptor as having responded

                if msg.proposalOrd > max {
                    p.proposalVal = msg.proposalVal
                    max = msg.proposalOrd
                }

            default:
            }
        }

        // phase-2(Accept-Accepted)
        p.accept()

        responded = make(map[int]bool)

        max = p.proposalOrd

        // wait for a majority of Accepted responses
        for len(responded) < len(p.acceptors) / 2 + 1 { // Majority = more than half of acceptors
            msg := <-p.receives

            switch msg.status {
            case Accepted:
                responded[msg.from] = true // mark the acceptor as having responded

                if msg.proposalOrd > max {
                    max = msg.proposalOrd
                }

            default:
            }
        }

        // check if the current proposal number matches the maximum observed.
        if p.proposalOrd == max { // if no higher proposal was observed, consensus is reached
            break;
        }

        p.proposalOrd = max
    }

    // notify learners the chosen value
    p.chosen()
}

func (p *proposer) prepare() {
    p.proposalOrd++

    msg := NewPrepareMessage(p.id, p.proposalOrd)
    fmt.Printf("Proposer %v: sending Prepare with order number %v\n", p.id, p.proposalOrd)
    broadcast(p.acceptors, msg)
}

func (p *proposer) accept() {
    msg := NewAcceptMessage(p.id, p.proposalOrd, p.proposalVal)

    fmt.Printf("Proposer %v: sending Accept with order number %v\n", p.id, p.proposalOrd)
    broadcast(p.acceptors, msg)
}

func (p *proposer) chosen() {
    msg := NewChosenMessage(p.id, p.proposalVal)

    fmt.Printf("Proposer %v: sending Chosen with val %v\n", p.id, p.proposalVal)
    broadcast(p.learners, msg)
}

func broadcast(peers []chan message, msg message) {
    for _, peer := range peers {
        peer <- msg
    }
}
