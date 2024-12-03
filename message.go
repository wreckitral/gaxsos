package gaxsos

type message struct {
    status      messageStatus // status of the message
    from        int // id of the sender
    proposalOrd int // the order of proposal, higher = newer
    proposalVal  string // the proposal value
}

type messageStatus int

const (
    Prepare messageStatus = iota // = 0; iota = easily assigned message status to each int that is assigned to this following status
    Promise // = 1
    Accept // = 2
    Accepted // = 3
    Chosen // = 4
)

// phase-1(Prepare-Promise), proposer asks acceptors to prepare for a proposal with number of proposalOrd, eg. saying hai
func NewPrepareMessage(from, proposalOrd int) message {
    return message{status: Prepare, from: from, proposalOrd: proposalOrd}
}

// used by acceptors to promise not to accept proposals with number lower than proposalOrd
func NewPromiseMessage(from, proposalOrd int, proposalVal string) message {
    return message{status: Promise, from: from, proposalOrd: proposalOrd, proposalVal: proposalVal}
}

// phase-2(Accept-Accepted), send by proposer to ask acceptors to accept a proposal
func NewAcceptMessage(from, proposalOrd int, proposalVal string) message {
    return message{status: Accept, from: from, proposalOrd: proposalOrd, proposalVal: proposalVal}
}

// send by acceptors to say that they have accepted a proposal
func NewAcceptedMessage(from, proposalOrd int) message {
    return message{status: Accepted, from: from, proposalOrd: proposalOrd}
}

// proposer sent to learners (all node)  to announce the final chosen value after consensus
func NewChosenMessage(from int, proposalVal string) message {
    return message{status: Chosen, from: from, proposalVal: proposalVal}
}
