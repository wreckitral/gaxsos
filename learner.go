package gaxsos

import "fmt"

type learner struct {
    id int // identifier for learner
    receives chan message // channel to listen to chosen value
}

func NewLearner(id int, receives chan message) *learner {
    l := learner{}

    l.id = id
    l.receives = receives

    return &l
}

func (l *learner) Run() string {
    val := "zero" // default value

    for val == "zero" {
        msg := <- l.receives

        switch msg.status {
        case Chosen:
            val = msg.proposalVal

        default:
        }
    }

    fmt.Printf("Learner %v: Chosen %v\n", l.id, val)

    return val
}

