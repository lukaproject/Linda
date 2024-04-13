package models

import (
	"encoding/json"

	"github.com/lukaproject/xerr"
)

func Serialize[T any](v T) []byte {
	return xerr.Must(json.Marshal(v))
}

func Deserialize[T any](b []byte, v T) {
	xerr.Must0(json.Unmarshal(b, v))
}
