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
    n := NewNetwork(1, 1, 2, []string{"test_value"})
    go n.acceptors[0].Run()
    go n.proposers[0].Run()
    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoProposersSameValue(t *testing.T) {
    n := NewNetwork(2, 1, 2, []string{"same_value", "same_value"})
    go n.acceptors[0].Run()
    go n.proposers[0].Run()
    go n.proposers[1].Run()
    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoProposersDifferentValue(t *testing.T) {
    n := NewNetwork(2, 1, 2, []string{"value_one", "value_two"})
    go n.acceptors[0].Run()
    go n.proposers[0].Run()
    go n.proposers[1].Run()
    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersDifferentValues(t *testing.T) {
    n := NewNetwork(5, 1, 2, []string{"value_1", "value_2", "value_3", "value_4", "value_5"})
    go n.acceptors[0].Run()
    for _, p := range n.proposers {
        go p.Run()
    }
    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestTwoAcceptors(t *testing.T) {
    n := NewNetwork(1, 2, 2, []string{"value_three"})
    go n.acceptors[0].Run()
    go n.acceptors[1].Run()
    go n.proposers[0].Run()
    if n.learners[0].Run() != n.learners[1].Run() {
        t.Errorf("Did not receive the same value!")
    }
}

func TestManyProposersManyAcceptorsSameValue(t *testing.T) {
    n := NewNetwork(5, 5, 2, []string{"same_value", "same_value", "same_value", "same_value", "same_value"})
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
    n := NewNetwork(5, 5, 2, []string{"value_one", "value_two", "value_three", "value_four", "value_five"})
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
    n := NewNetwork(5, 5, 2, []string{"value_one", "value_two", "value_one", "value_two", "value_one"})
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
