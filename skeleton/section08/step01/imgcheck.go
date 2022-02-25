package imgcheck

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// ベースとなるエラー
var (
	// TODO: errors.New関数を使って"画像フォーマットが違います"というエラーを作成し、変数ErrFormatに代入する

)

// バリデーションルールを表す関数
type Rule func(img image.Image, format string) error

// 画像フォーマットをチェックするルール
func Format(format string) Rule {
	return func(_ image.Image, _format string) error {
		// TODO: フォーマットが一致しない場合はErrFormatを返す

		return nil
	}
}

// 画像のバリデーションを行う
func Validate(r io.Reader, rules ...Rule) error {
	// 画像を読み込む
	img, format, err := image.Decode(r)
	switch {
	case err == image.ErrFormat:
		// 画像として読み込めなかった
		return nil
	case err != nil:
		return err
	}

	for _, rule := range rules {
		// TODO: ruleの呼び出しでエラーが発生した場合はreturnする
		rule(img, format)
	}

	return nil
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
			return err
		}
		return nil
	}
	return filepath.Walk(root, walkfunc)
}
