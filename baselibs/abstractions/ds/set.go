package ds

import (
	"github.com/ecodeclub/ekit/mapx"
)

type empty struct{}
type Set[T comparable] map[T]empty

func (s *Set[T]) ListByChan(ch chan T) {
	for k := range *s {
		ch <- k
	}
	close(ch)
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Insert(v T) {
	(*s)[v] = empty{}
}

func (s *Set[T]) Remove(v T) {
	delete((*s), v)
}

func (s *Set[T]) Exist(v T) bool {
	_, ok := (*s)[v]
	return ok
}

func (s *Set[T]) ToArray() []T {
	return mapx.Keys(*s)
}
