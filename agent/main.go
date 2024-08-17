package main

import (
	"Linda/agent/internal/config"
	"Linda/agent/internal/handler"
)

func main() {
	h := handler.NewHandler(config.Instance())
	h.Run()
}
