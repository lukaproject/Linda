package main

import (
	"Linda/agent/internal/handler"
)

func main() {
	h := handler.NewHandler(nil)
	h.Run()
}
