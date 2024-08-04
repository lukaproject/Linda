package models

import (
	"encoding/json"
	"io"

	"github.com/lukaproject/xerr"
)

func Serialize[T any](v T) []byte {
	return xerr.Must(json.Marshal(v))
}

func Deserialize[T any](b []byte, v T) {
	xerr.Must0(json.Unmarshal(b, v))
}

func ReadJSON[T any](reader io.Reader, v T) {
	Deserialize(xerr.Must(io.ReadAll(reader)), v)
}
