package agents

import (
	"Linda/baselibs/codes/errno"
	"sync"
)

const (
	node_STATES_JOINING = "node_STATES_JOINING"
	node_STATES_FREEING = "node_STATES_FREEING"
	node_STATES_FREE    = "node_STATES_FREE"
	node_STATES_ON_WORK = "node_STATES_ON_WORK"
	emptyBagName        = ""
)

type nodeStates struct {
	BagName string
	State   string

	mut sync.Mutex
}

func (ns *nodeStates) Join(bagName string) (err error) {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	if ns.State == node_STATES_FREE {
		ns.BagName = bagName
		ns.State = node_STATES_JOINING
	} else if ns.State == node_STATES_JOINING && ns.BagName == bagName {
		// TODO
		// it is ok for rejoin with same bagName
	} else {
		return errno.ErrNodeBelongsToAnotherBag
	}
	return
}

func (ns *nodeStates) JoinFinished(bagName string) {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	if ns.BagName == bagName && ns.State == node_STATES_JOINING {
		ns.State = node_STATES_ON_WORK
	}
}

func (ns *nodeStates) Free() {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	if ns.State == node_STATES_ON_WORK {
		ns.BagName = emptyBagName
		ns.State = node_STATES_FREEING
	}
}

func (ns *nodeStates) FreeFinished() {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	if ns.BagName == emptyBagName && ns.State == node_STATES_FREEING {
		ns.State = node_STATES_FREE
	}
}

func (ns *nodeStates) GetBagName() string {
	if name, state := ns.GetBagNameWithState(); state == node_STATES_ON_WORK {
		return name
	}
	return emptyBagName
}

// GetBagNameWithState is a method return bagname and state
func (ns *nodeStates) GetBagNameWithState() (bagName string, state string) {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	return ns.BagName, ns.State
}

func (ns *nodeStates) IsOnGoingStates() bool {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	return ns.State == node_STATES_FREEING || ns.State == node_STATES_JOINING
}

func (ns *nodeStates) IsSteadyStates() bool {
	ns.mut.Lock()
	defer ns.mut.Unlock()
	return ns.State == node_STATES_ON_WORK || ns.State == node_STATES_FREE
}

func newNodeStates() *nodeStates {
	return &nodeStates{
		BagName: emptyBagName,
		State:   node_STATES_FREE,
	}
}
