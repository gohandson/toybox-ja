package parawalk

import (
	"context"
	"io/fs"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

type WalkFunc func(ctx context.Context, path string, info fs.FileInfo, err error) error

func Walk(ctx context.Context, root string, fn WalkFunc) error {
	// TODO: errgroup.WithContextでerrgroupを作成する

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		// すでにコンテキストがキャンセルされている
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// エラー処理が必要またはディレクトリの場合はそのまま処理
		if err != nil || info.IsDir() {
			return fn(ctx, path, info, err)
		}

		// ファイルの場合はゴールーチンで処理
		eg.Go(func() error {
			return fn(ctx, path, info, err)
		})

		return nil
	})

	if err != nil {
		return err
	}

	return eg.Wait()
}
