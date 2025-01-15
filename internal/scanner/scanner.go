package scanner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type FileInfo struct {
	FileName     string
	ByteSize     int64
	DateOfChange time.Time
	Extension    string
}

func ScanDir(directory string) ([]FileInfo, error) {
	var files []FileInfo
	var mu sync.Mutex
	var wg sync.WaitGroup

	numWorkers := 6
	filePaths := make(chan string, 100)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range filePaths {
				info, err := os.Stat(path)
				if err != nil {
					continue
				}
				if info.IsDir() {
					continue
				}
				file := FileInfo{
					FileName:     info.Name(),
					ByteSize:     info.Size(),
					DateOfChange: info.ModTime(),
					Extension:    filepath.Ext(info.Name()),
				}
				mu.Lock()
				files = append(files, file)
				mu.Unlock()
			}
		}()
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths <- path
		}
		return nil
	})
	close(filePaths)

	wg.Wait()

	if err != nil {
		return nil, err
	}

	return files, nil
}

func FindFilesWithExtension(directory, extension string) ([]string, error) {
	cmd := exec.Command("find", directory, "-type", "f", "-name", "*"+extension)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	files := strings.Split(strings.TrimSpace(out.String()), "\n")
	return files, nil
}

func FindFilesWithKeyword(directory, keyword string) ([]string, error) {
	cmdFind := exec.Command("find", directory, "-type", "f")
	var outFind bytes.Buffer
	cmdFind.Stdout = &outFind

	if err := cmdFind.Run(); err != nil {
		return nil, fmt.Errorf("error executing find command: %v", err)
	}

	cmdGrep := exec.Command("grep", "-l", keyword)
	cmdGrep.Stdin = &outFind

	var outGrep bytes.Buffer
	cmdGrep.Stdout = &outGrep

	if err := cmdGrep.Run(); err != nil {
		return nil, fmt.Errorf("error executing grep command: %v", err)
	}

	files := strings.Split(strings.TrimSpace(outGrep.String()), "\n")

	return files, nil
}
