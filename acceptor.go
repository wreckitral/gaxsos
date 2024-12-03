package gaxsos

import "fmt"

type acceptor struct {
    id int

    AcceptedProposalOrd int // proposal order number that this acceptor has accepted so far
    AcceptedProposalVal string // proposal value that associated with the accepted proposal order

    maxProposalOrd int// tracks the highest proposal order number received in the Prepare phase

    receives chan message // channel for receiving message from proposers
    proposers []chan message // channels of to sends responses back to proposers
}

func NewAcceptor(id int, receives chan message, proposers []chan message) *acceptor {
    a := acceptor{}

    a.id = id
    a.AcceptedProposalOrd = 0
    a.AcceptedProposalVal = ""
    a.maxProposalOrd = 0
    a.receives = receives
    a.proposers = proposers

    return &a
}

func (a *acceptor) Run() {
    fmt.Printf("Acceptor %v: started\n", a.id)

    for {
        msg := <-a.receives

        switch msg.status {
        // phase-1(Prepare-Promise)
        case Prepare:
            if msg.proposalOrd > a.maxProposalOrd { // if proposal order number is bigger than acceptor's biggest order number then acceptor will send promise
                a.maxProposalOrd = msg.proposalOrd
                a.proposers[msg.from] <- NewPromiseMessage(a.id, a.AcceptedProposalOrd, a.AcceptedProposalVal)
                fmt.Printf("Acceptor %v: sending Promise on order number %v\n", a.id, a.AcceptedProposalOrd)
            }

        // phase-1(Accept-Accepted)
        case Accept:
            if msg.proposalOrd >= a.maxProposalOrd { // if order number on msg is bigger or same then acceptor's max order number is msg's
                a.maxProposalOrd = msg.proposalOrd
                a.AcceptedProposalOrd = msg.proposalOrd
                a.AcceptedProposalVal = msg.proposalVal
            }

            a.proposers[msg.from] <- NewAcceptedMessage(a.id, a.maxProposalOrd)
            fmt.Printf("Acceptor %v: sending Accepted on order number %v\n", a.id, a.AcceptedProposalOrd)
        }
    }
}
