package app

import (
	"context"
	"log/slog"
	"math/rand"
	"sync"

	"github.com/Shyyw1e/osmoview-test-task/internal/domain"
	"github.com/Shyyw1e/osmoview-test-task/internal/ports/storage"
)

type Config struct {
	Iterations int
	Threads int
	FileCount int
}

type Runner struct {
	writer storage.Writer
	logger slog.Logger
}

func NewRunner(writer storage.Writer, logger *slog.Logger) *Runner {
	return &Runner{
		writer: writer,
		logger: *logger,
	}	
} 

func (r *Runner) Run(ctx context.Context, cfg Config) {
	r.logger.Info("Starting generation",
		slog.Int("threads", cfg.Threads),
		slog.Int("iterations", cfg.Iterations),
		slog.Int("files", cfg.FileCount),
	)

	var wg sync.WaitGroup

	iterChan := make(chan int)

	for i := 0; i < cfg.Threads; i++ {
		wg.Add(1)
		go func (workerID int) {
			defer wg.Done()
			for id := range iterChan {
				select {
				case <-ctx.Done():
					r.logger.Warn("Context cancelled", slog.Int("worker", workerID))
					return
				default:
					data := domain.RandomData(id)
					fileIndex := rand.Intn(cfg.FileCount)
					err := r.writer.Write(data, fileIndex)
					if err != nil {
						r.logger.Error("Write failed", slog.Int("worker", workerID), slog.Int("file", fileIndex), slog.String("error", err.Error()))
					} else {
						r.logger.Debug("Data written",
							slog.Int("worker", workerID),
							slog.Int("iteration", id),
							slog.Int("file", fileIndex),
						)

					}

				}
			}
		}(i)
	}

	for i := 0; i < cfg.Iterations; i++ {
		iterChan <- i
	}
	close(iterChan)
	wg.Wait()
	r.logger.Info("Generation completed")	
}