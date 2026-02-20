// Параллельное сжатие файлов с использованием группы ожидания
package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"compress/gzip"
)

func main() {
	var wg sync.WaitGroup

	for _, file := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			compess(filename)
			wg.Done()
		}(file)
	}

	wg.Wait()
	fmt.Printf("compressed %d files\n", len(os.Args[1:]))
}

func compess(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}
