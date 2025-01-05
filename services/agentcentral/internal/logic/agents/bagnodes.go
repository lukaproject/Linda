package agents

import (
	"Linda/baselibs/abstractions/ds"
	"Linda/protocol/models"
	"errors"
	"sync"
)

const (
	defaultChannelSize = 20
)

type BagNodes struct {
	rwMut sync.RWMutex
	mp    map[string]ds.Set[string]
}

func (bn *BagNodes) Initial() {
	bn.mp = make(map[string]ds.Set[string])
}

func (bn *BagNodes) ListByBagName(bagName string) (ch chan string, err error) {
	bn.rwMut.RLock()
	defer bn.rwMut.RUnlock()

	if nodes, ok := bn.mp[bagName]; !ok {
		err = errors.New(models.BAG_NOT_EXIST)
		return
	} else {
		ch = make(chan string, defaultChannelSize)
		go nodes.ListByChan(ch)
		return ch, nil
	}
}

func (bn *BagNodes) Add(bagName, nodeId string) (err error) {
	bn.rwMut.Lock()
	defer bn.rwMut.Unlock()

	nodes, ok := bn.mp[bagName]
	if !ok {
		bn.mp[bagName] = nodes
	}
	nodes.Insert(nodeId)
	return
}
