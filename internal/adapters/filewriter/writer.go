package filewriter

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/Shyyw1e/osmoview-test-task/internal/domain"
	"github.com/Shyyw1e/osmoview-test-task/internal/ports/storage"
	"gopkg.in/yaml.v3"

)

type FileWriter struct {
	dir string
	filecount int
	locks []*sync.Mutex
}

func New(dir string, filecount int) storage.Writer {
	locks := make([]*sync.Mutex, filecount)
	for i := 0; i < filecount; i++ {
		locks[i] = &sync.Mutex{}
	}
	return &FileWriter {
		dir: dir,
		filecount: filecount,
		locks: locks,
	}
}

func (fw *FileWriter)Write(data domain.Data, fileIndex int) error {
	if fileIndex < 0 || fileIndex >= fw.filecount {
		return fmt.Errorf("invalid file index: %d", fileIndex)
	}
	fw.locks[fileIndex].Lock()
	defer fw.locks[fileIndex].Unlock()

	path := filepath.Join(fw.dir, fmt.Sprintf("data_%d.yaml", fileIndex))

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)
	defer enc.Close()

	if err := enc.Encode(&data); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}

	return nil
}

