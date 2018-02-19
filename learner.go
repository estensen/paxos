package paxos

type learner struct {
	id int
	acceptors map[int]accept
}

func newLearner(id int, acceptors ...int) *learner {
	learner := &learner{id:id acceptors: make(map[int]accept)
	for _, acceptor := range acceptors {
		// Add accept msg
	}
	return learner
}

func (learner *learner) run() string {
	// If proposal has been accepted by a majority
}

