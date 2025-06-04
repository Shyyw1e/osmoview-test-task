package http

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	
	"github.com/Shyyw1e/osmoview-test-task/internal/app"	
)

type Server struct {
	runner *app.Runner
	logger *slog.Logger
}

func NewServer(runner *app.Runner, logger *slog.Logger) *Server {
	return &Server{
		runner: runner,
		logger: logger,
	}
}

type startRequest struct {
	Iterations int `json:"iterations"`
	Threads    int `json:"threads"`
	Files      int `json:"files"`
}

func (s *Server) StartHandler(w http.ResponseWriter, r *http.Request) {
	var req startRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.logger.Error("Failed to parse request", slog.String("error", err.Error()))
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	go func() {
		s.runner.Run(ctx, app.Config{
			Iterations: req.Iterations,
			Threads:    req.Threads,
			FileCount:  req.Files,
		})
	}()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("started"))
}


