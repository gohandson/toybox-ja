package parawalk

import (
	"io/fs"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

type WalkFunc func(path string, info fs.FileInfo, err error) error

func Walk(root string, fn WalkFunc) error {
	var eg errgroup.Group
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		// エラー処理が必要またはディレクトリの場合はそのまま処理
		if err != nil || info.IsDir() {
			return fn(path, info, err)
		}

		// TODO: 関数fnをeg.Goで呼び出す
		// ファイルの場合はゴールーチンで処理
		

		return nil
	})

	if err != nil {
		return err
	}

	// TODO: errgroupの結果を待ってエラーがあったら戻り値にする

}
