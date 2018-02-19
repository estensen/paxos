package paxos

type proposer struct {
	id int
	seq int
	proposalNum int
	proposalVal string
	acceptors map[int]promise
}

func newProposer(id, int, value string, acceptors ...int) *proposer {
	proposer := &proposer{id: id, seq: 0, proposalVal: value, acceptors: make(map[int]promise)}
	for _, acceptor := range acceptors {
		proposer.acceptors[acceptor] = message{}
	}
	return proposer

func (proposer *proposer) propose() []message {
	for !proposer.majorityReached() {
		proposer.prepare()
	}

	proposer.propose()
}

func (proposer *proposer) prepare() []message {
	// Send msgs to acceptors
}

func (proposer *proposer) propose() []message {
	// Send msgs to acceptors
}

func (proposer *proposer) majorityReached() bool {
	// Calculate if majority
}
