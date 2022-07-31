package filecache

import (
	"fmt"
	"os"
	"sync"
)

var fileCache = "file.cache"

type FileCache struct {
	mu       sync.Mutex
	rmu      sync.RWMutex
	cache    map[string]string
	MaxItems int
	MaxSize  int64
}

func (f *FileCache) CreateFileCache() (*os.File, error) {
	file, err := os.Create(fileCache)
	if err != nil {
		return nil, fmt.Errorf("Unable to create fileCache: %v", err)
	}

	return file, nil
}

func (f *FileCache) lock() {
	f.mu.Lock()
}

func (f *FileCache) unlock() {
	f.mu.Unlock()
}

func (f *FileCache) rlock() {
	f.rmu.RLock()
}

func (f *FileCache) runlock() {
	f.rmu.RUnlock()
}

func (f *FileCache) Write(key, value string) error {
	f.lock()
	defer f.unlock()

	if key == "" && value == "" {
		return fmt.Errorf("Invalid key/value data")
	}

	f.cache[key] = value

	return nil
}

func (f *FileCache) Read(key string) (string, error) {
	var err error

	if key == nil {
		return "", fmt.Errorf("Empty key found")
	}

	f.rlock()
	v, err := f.cache[key]
	if err != nil {
		return "", fmt.Errorf("Key does not exist: %v", err)
	}

	f.runlock()
	return v, nil
}
