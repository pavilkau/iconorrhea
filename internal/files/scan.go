package files

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Name    string
	File    []byte
	Size    int64
	ModTime time.Time
	// Seen    bool
	// SeenAt  time.Time
}

func Scan(path string) (files []File, err error) {
	rawFiles, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, rawFile := range rawFiles {
		fileInfo, _ := rawFile.Info()
		file, err := os.ReadFile(filepath.Join(path, fileInfo.Name()))
		if err != nil {
			log.Printf("failed to read file: %s", err)
			continue
		}

		newFile := File{fileInfo.Name(), file, fileInfo.Size(), fileInfo.ModTime()}
		files = append(files, newFile)
	}

	return files, nil
}
