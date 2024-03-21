package main

import (
	"log/slog"

	"main/flarelog"
)

func test1() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(flarelog.NewHandler(opts))

	logger.Info("Info Level Log")
	logger.Debug("This is a debug message")
	logger.Warn("This is a warn message")
	logger.Error("This is an error message")
}

func main() {
	test1()
}
