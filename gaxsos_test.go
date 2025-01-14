package gaxsos

import "testing"

func TestBroadcast(t *testing.T) {
    acceptors := []chan message{make(chan message, 100), make(chan message, 100)}
    msg := message{proposalVal: "test_message"}
    broadcast(acceptors, msg)
    for _, acceptor := range acceptors {
        if <-acceptor != msg {
            t.Errorf("Received message was different to sent message!")
        }
    }
}

func TestSingleProposer(t *testing.T) {
    n, err := NewNetwork(1, 3, 2, []string{"test_value"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    go n.acceptors[0].Run()
    go n.acceptors[1].Run()
    go n.acceptors[2].Run()
    go n.proposers[0].Run()

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoProposersSameValue(t *testing.T) {
    n, err := NewNetwork(2, 7, 2, []string{"same_value", "same_value"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    go n.proposers[0].Run()
    go n.proposers[1].Run()

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoProposersDifferentValue(t *testing.T) {
    n, err := NewNetwork(2, 7, 2, []string{"value_one", "value_two"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    go n.proposers[0].Run()
    go n.proposers[1].Run()

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersDifferentValues(t *testing.T) {
    n, err := NewNetwork(5, 11, 2, []string{"value_1", "value_2", "value_3", "value_4", "value_5"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    for _, p := range n.proposers {
        go p.Run()
    }

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoAcceptors(t *testing.T) {
    n, err := NewNetwork(1, 7, 2, []string{"value_three"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    go n.acceptors[0].Run()
    go n.acceptors[1].Run()
    go n.proposers[0].Run()

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersManyAcceptorsSameValue(t *testing.T) {
    n, err := NewNetwork(5, 11, 2, []string{"same_value", "same_value", "same_value", "same_value", "same_value"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    for _, p := range n.proposers {
        go p.Run()
    }

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersManyAcceptorsDifferentValues(t *testing.T) {
    n, err := NewNetwork(5, 11, 2, []string{"value_one", "value_two", "value_three", "value_four", "value_five"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    for _, p := range n.proposers {
        go p.Run()
    }

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersManyAcceptorsSemiSameValues(t *testing.T) {
    n, err := NewNetwork(5, 11, 2, []string{"value_one", "value_two", "value_one", "value_two", "value_one"})
    if err != nil {
        t.Fatalf("Failed to create network: %v", err)
    }

    for _, a := range n.acceptors {
        go a.Run()
    }

    for _, p := range n.proposers {
        go p.Run()
    }

    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestFaultTolerance(t *testing.T) {
	testCases := []struct {
		name            string
		nProposers      int
		nAcceptors      int
		nLearners       int
		values          []string
		expectedError   string
		shouldFail      bool
	}{
		{
			name:           "Insufficient Acceptors - 3 Proposers",
			nProposers:     3,
			nAcceptors:     5,  // should be 7 for 3 proposers
			nLearners:      2,
			values:         []string{"value_one", "value_two", "value_three"},
			expectedError:  "at least 2f + 1 acceptor is required for Paxos consensus",
			shouldFail:     true,
		},
		{
			name:           "Sufficient Acceptors - 3 Proposers",
			nProposers:     3,
			nAcceptors:     7,  // 2f+1 = 7 for 3 proposers
			nLearners:      2,
			values:         []string{"value_one", "value_two", "value_three"},
			expectedError:  "",
			shouldFail:     false,
		},
		{
			name:           "Insufficient Acceptors - 5 Proposers",
			nProposers:     5,
			nAcceptors:     9,  // should be 11 for 5 proposers
			nLearners:      2,
			values:         []string{"value_one", "value_two", "value_three", "value_four", "value_five"},
			expectedError:  "at least 2f + 1 acceptor is required for Paxos consensus",
			shouldFail:     true,
		},
		{
			name:           "Sufficient Acceptors - 5 Proposers",
			nProposers:     5,
			nAcceptors:     11, // 2f+1 = 11 for 5 proposers
			nLearners:      2,
			values:         []string{"value_one", "value_two", "value_three", "value_four", "value_five"},
			expectedError:  "",
			shouldFail:     false,
		},
		{
			name:           "Zero Proposers",
			nProposers:     0,
			nAcceptors:     1,
			nLearners:      2,
			values:         []string{},
			expectedError:  "at least one proposer is required",
			shouldFail:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Attempt to create network
			_, err := NewNetwork(tc.nProposers, tc.nAcceptors, tc.nLearners, tc.values)

			// Check if error occurred as expected
			if tc.shouldFail {
				if err == nil {
					t.Errorf("Expected an error, but got none")
					return
				}

				// Check error message if specific error is expected
				if tc.expectedError != "" && err.Error() != tc.expectedError {
					t.Errorf("Unexpected error message. Got: %v, Want: %v",
						err.Error(), tc.expectedError)
				}
			} else {
				// Should not fail
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
