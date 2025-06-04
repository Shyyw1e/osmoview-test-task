package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Shyyw1e/osmoview-test-task/internal/adapters/filewriter"
	"github.com/Shyyw1e/osmoview-test-task/internal/adapters/logger"
	"github.com/Shyyw1e/osmoview-test-task/internal/app"
	portHttp "github.com/Shyyw1e/osmoview-test-task/internal/ports/http"
)


func main() {
	l := logger.New(slog.LevelDebug)

	outputDir := "output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		l.Error("failed to create output directory", slog.String("error", err.Error()))
		os.Exit(1)
	}

	writer := filewriter.New(outputDir, 10)
	runner := app.NewRunner(writer, l)
	server := portHttp.NewServer(runner, l)

	http.HandleFunc("/start", server.StartHandler)
	l.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		l.Error("failed to listen server", slog.String("error", err.Error()))
	}
}