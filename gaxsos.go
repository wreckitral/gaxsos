package gaxsos

// network structure holds all the components of a Paxos system.
type network struct {
	proposers []*proposer // a slice of proposer components in the network.
	acceptors []*acceptor // a slice of acceptor components in the network.
	learners  []*learner  // a slice of learner components in the network.
}

func NewNetwork(nProposers, nAcceptors, nLearners int, vs []string) *network {
	// create communication channels for proposers, acceptors, and learners.
	cProposers := makeChannels(nProposers)
	cAcceptors := makeChannels(nAcceptors)
	cLearners := makeChannels(nLearners)

	// initialize the network struct.
	n := new(network)
	n.proposers = make([]*proposer, nProposers) // allocate slice for proposers.
	n.acceptors = make([]*acceptor, nAcceptors) // allocate slice for acceptors.
	n.learners = make([]*learner, nLearners)   // allocate slice for learners.

	// create proposers, each with an initial value from the vs slice.
	for i := range n.proposers {
		n.proposers[i] = NewProposer(i, vs[i], cProposers[i], cAcceptors, cLearners)
	}

	// create acceptors, with each acceptor having its own communication channel.
	for i := range n.acceptors {
		n.acceptors[i] = NewAcceptor(i, cAcceptors[i], cProposers)
	}

	// create learners, each with its own communication channel.
	for i := range n.learners {
		n.learners[i] = NewLearner(i, cLearners[i])
	}

	return n // return the initialized network.
}

// makeChannels creates a slice of channels for communication, each with a buffer size of 1024.
func makeChannels(n int) []chan message {
	chans := make([]chan message, n) // create a slice of channels with size n.

	for i := range chans {
		chans[i] = make(chan message, 1024) // initialize each channel with a buffer size of 1024.
	}
	return chans // return the slice of channels.
}

// start starts the Paxos algorithm by launching goroutines for each component.
func (n *network) Start() {
	// start each learner in a separate goroutine.
	for _, l := range n.learners {
		go l.Run() // start the learner's Run method concurrently.
	}

	// start each acceptor in a separate goroutine.
	for _, a := range n.acceptors {
		go a.Run() // start the acceptor's Run method concurrently.
	}

	// start each proposer in a separate goroutine.
	for _, p := range n.proposers {
		go p.Run() // start the proposer's Run method concurrently.
	}
}

