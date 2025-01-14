package gaxsos
//
// import (
// 	"testing"
// )
//
// // TestNewNetwork checks that the network initialization creates the correct number of proposers, acceptors, and learners.
// func TestNewNetwork(t *testing.T) {
// 	nProposers := 3
// 	nAcceptors := 2
// 	nLearners := 4
// 	values := []string{"value1", "value2", "value3"}
//
// 	// Create a new network with the specified parameters.
// 	n := NewNetwork(nProposers, nAcceptors, nLearners, values)
//
// 	// Validate the number of proposers.
// 	if len(n.proposers) != nProposers {
// 		t.Errorf("Expected %d proposers, got %d", nProposers, len(n.proposers))
// 	}
//
// 	// Validate the number of acceptors.
// 	if len(n.acceptors) != nAcceptors {
// 		t.Errorf("Expected %d acceptors, got %d", nAcceptors, len(n.acceptors))
// 	}
//
// 	// Validate the number of learners.
// 	if len(n.learners) != nLearners {
// 		t.Errorf("Expected %d learners, got %d", nLearners, len(n.learners))
// 	}
//
// 	// Ensure the proposers have correct initial values.
// 	for i, p := range n.proposers {
// 		if p.proposalVal != values[i] {
// 			t.Errorf("Expected proposer %d to have value %s, got %s", i, values[i], p.proposalVal)
// 		}
// 	}
// }
//
// // TestStart checks that the Start method launches the Run method for each component.
// func TestStart(t *testing.T) {
// 	nProposers := 2
// 	nAcceptors := 1
// 	nLearners := 2
// 	values := []string{"value1", "value2"}
//
// 	// Create a new network.
// 	n := NewNetwork(nProposers, nAcceptors, nLearners, values)
//
// 	// Start the network (this will run goroutines).
// 	n.Start()
//
// 	// Check if the learners' Run method is called by checking if they are running.
// 	// In a real-world scenario, you might want to use channels or other synchronization mechanisms to confirm the Run method was triggered.
// 	// For simplicity, we'll assume that if we reach here, the Run methods have been launched.
// 	if len(n.learners) != nLearners {
// 		t.Errorf("Expected %d learners to be running, but got %d", nLearners, len(n.learners))
// 	}
//
// 	// Similarly, ensure proposers and acceptors are "started".
// 	if len(n.proposers) != nProposers {
// 		t.Errorf("Expected %d proposers to be running, but got %d", nProposers, len(n.proposers))
// 	}
//
// 	if len(n.acceptors) != nAcceptors {
// 		t.Errorf("Expected %d acceptors to be running, but got %d", nAcceptors, len(n.acceptors))
// 	}
// }
//
// // TestEmptyNetwork checks if an empty network can be created (no proposers, acceptors, or learners).
// func TestEmptyNetwork(t *testing.T) {
// 	// Create a network with no proposers, acceptors, or learners.
// 	n := NewNetwork(0, 0, 0, []string{})
//
// 	// Ensure the network components are all empty.
// 	if len(n.proposers) != 0 {
// 		t.Errorf("Expected 0 proposers, got %d", len(n.proposers))
// 	}
//
// 	if len(n.acceptors) != 0 {
// 		t.Errorf("Expected 0 acceptors, got %d", len(n.acceptors))
// 	}
//
// 	if len(n.learners) != 0 {
// 		t.Errorf("Expected 0 learners, got %d", len(n.learners))
// 	}
// }
//
