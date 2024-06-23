package comm

import "sync"

type ReleaseLocker struct {
	Mut *sync.Mutex
}

func NewRLocker(mut *sync.Mutex) *ReleaseLocker {
	return &ReleaseLocker{Mut: mut}
}

func (rl *ReleaseLocker) Run(f func()) {
	rl.Mut.Lock()
	defer rl.Mut.Unlock()
	f()
}
