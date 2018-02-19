package paxos

type acceptor struct {
	id int
	learners []int
	accept message
	promised promise
}

func newAcceptor(id int, learners ...int) *acceptor {
	return &acceptor{id: id, promised: message{}, learners: learners}
}

func (acceptor *acceptor) run() {
	// Handle prepare and propose msgs
}

