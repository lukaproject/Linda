package xctx

import "sync"

type ReleaseLocker struct {
	Mut sync.Locker
}

func NewLocker(mut sync.Locker) *ReleaseLocker {
	return &ReleaseLocker{Mut: mut}
}

func (rl *ReleaseLocker) Run(f func()) {
	rl.Mut.Lock()
	defer rl.Mut.Unlock()
	f()
}
