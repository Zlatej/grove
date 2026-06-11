package main

import (
	"fmt"
	"log/slog"

	"github.com/zlatej/grove/internal/config"
)

func main() {
	_, err := config.Load()
	if err != nil {
		slog.Error("loading config", "error", err.Error())
	}
	fmt.Println("67")
}
