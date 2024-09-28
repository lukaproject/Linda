package agents

import (
	"strings"

	"github.com/google/uuid"
)

func GenNodeId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "") + ":nodeId"
}
