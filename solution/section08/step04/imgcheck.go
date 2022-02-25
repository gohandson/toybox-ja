package imgcheck

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"go.uber.org/multierr"
)

// 画像のバリデーションを行う
func Validate(r io.Reader, rules ...Rule) error {
	// 画像を読み込む
	img, format, err := image.Decode(r)
	switch {
	case errors.Is(err, image.ErrFormat):
		// 画像として読み込めなかった
		return nil
	case err != nil:
		return err
	}

	var rerr error
	for _, rule := range rules {
		rerr = multierr.Append(rerr, rule.Validate(img, format))
	}

	return rerr
}

// ディレクトリ以下の画像ファイルのバリデーションを行う
func ValidateDir(root string, rules ...Rule) error {
	walkfunc := func(path string, info fs.FileInfo, err error) (rerr error) {

		// エラーが発生した
		if err != nil {
			return err
		}

		// ディレクトリ
		if info.IsDir() {
			return nil
		}

		// 変換前のファイルを開く
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		// 関数終了時にファイルを閉じる
		defer file.Close()

		// バリデーションをかける
		if err := Validate(file, rules...); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		return nil
	}
	return filepath.Walk(root, walkfunc)
}
