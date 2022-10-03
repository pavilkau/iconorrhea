package files

import (
	"os"
	"time"
)

type File struct {
	Name    string
	Size    int64
	ModTime time.Time
}

func Scan(path string) (files []File, err error) {
	rawFiles, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, rawFile := range rawFiles {
		info, _ := rawFile.Info()
		newFile := File{info.Name(), info.Size(), info.ModTime()}
		files = append(files, newFile)
	}

	return files, nil
}
