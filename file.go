package eyhash

import (
	"os"
	"time"
)

type Info struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func FileInfo(path string) (*Info, error) {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &Info{
		Name:    file.Name(),
		Size:    info.Size(),
		ModTime: info.ModTime(),
	}, nil
}
