package parawalk

import (
	"io/fs"
	"path/filepath"
	"sync"

	"go.uber.org/multierr"
)

type WalkFunc func(path string, info fs.FileInfo, err error) error

func Walk(root string, fn WalkFunc) error {
	var wg sync.WaitGroup
	errCh := make(chan error)
	rerr := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		// エラー処理が必要またはディレクトリの場合はそのまま処理
		if err != nil || info.IsDir() {
			return fn(path, info, err)
		}

		// ファイルの場合はゴールーチンで処理
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := fn(path, info, err)
			if err != nil {
				errCh <- err
			}
		}()

		return nil
	})

	go func() {
		for err := range errCh {
			rerr = multierr.Append(rerr, err)
		}
	}()

	wg.Wait()
	close(errCh)

	return rerr
}
