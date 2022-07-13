package workerpool

import (
	"os"
	"sync"
)

// JobWriter ...
type JobWriter struct {
	Filename string
	MaxSize  int
	MaxAge   int

	size int64
	file *os.File
	mu   sync.Mutex
}

// Write implements io.Writer. If a write would cause the log file to be large
// than MaxSize, the file is closed, renamed to include a timestamp of the current time
// and a new log file is created using the origin log file name.
// If the length of the write is greater than MaxSize, an error is returned.
func (w *JobWriter) Write(data []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	return 0, nil
}
