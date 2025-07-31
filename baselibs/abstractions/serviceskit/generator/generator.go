package generator

import (
	"strings"

	"github.com/google/uuid"
)

type Generator interface {
	NodeId() string
}

type generatorImpl struct{}

func (gi *generatorImpl) NodeId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "") + ":nodeId"
}
