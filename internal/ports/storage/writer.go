package storage

import "github.com/Shyyw1e/osmoview-test-task/internal/domain"

type Writer interface {
	Write(data domain.Data, fileIndex int) error
}