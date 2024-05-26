package models

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lukaproject/xerr"
)

func Serialize[T any](v T) []byte {
	return xerr.Must(json.Marshal(v))
}

func Deserialize[T any](b []byte, v T) {
	xerr.Must0(json.Unmarshal(b, v))
}

func ReadJSON[T any](r *http.Request, v T) {
	Deserialize(xerr.Must(io.ReadAll(r.Body)), v)
}
