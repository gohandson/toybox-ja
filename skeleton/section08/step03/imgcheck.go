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

// ベースとなるエラー
var (
	ErrFormat   = errors.New("画像フォーマットが違います")
	ErrTooLarge = errors.New("画像が大きすぎます")
)

// バリデーションエラー
type ValidationError struct {
	Rule Rule
	Err  error
}

func (err *ValidationError) Error() string {
	return err.Err.Error()
}

// TODO: Unwrapメソッド実装。Errフィールドの値を返す


// バリデーションルールを表すインタフェース
type Rule interface {
	Validate(img image.Image, format string) error
}

type formatRule struct {
	format string
}

func (r *formatRule) Validate(img image.Image, format string) error {
	if r.format != format {
		return &ValidationError{
			Rule: r,
			Err:  fmt.Errorf("期待%s 実際%s: %w", r.format, format, ErrFormat),
		}
	}
	return nil
}

// 画像フォーマットをチェックするルール
func Format(format string) (rule Rule) {
	return &formatRule{format: format}
}

type maxSizeRule struct {
	height *int
	width  *int
}

func (r *maxSizeRule) Validate(img image.Image, _ string) error {
	bounds := img.Bounds()
	var err error

	if r.height != nil && bounds.Dy() > *r.height {
		err = multierr.Append(err, &ValidationError{
			Rule: r,
			Err:  fmt.Errorf("期待する高さ%d 実際%d: %w", *r.height, bounds.Dy(), ErrTooLarge),
		})
	}

	if r.width != nil && bounds.Dx() > *r.width {
		err = multierr.Append(err, &ValidationError{
			Rule: r,
			// TODO: エラーをラップしてErrフィールドに設定する。高さの実装を参考にする。

		})

	}

	return err
}

// 高さをチェックするルール
func MaxHeight(h int) Rule {
	return &maxSizeRule{height: &h}
}

// 幅をチェックするルール
func MaxWidth(w int) Rule {
	return &maxSizeRule{width: &w}
}

// 画像のバリデーションを行う
func Validate(r io.Reader, rules ...Rule) error {
	// 画像を読み込む
	img, format, err := image.Decode(r)
	switch {
	case /* errがimage.ErrFormatかどうかerrors.Is関数で判定する */:
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
			// TODO: ファイルパスをつけてエラーをラップして返す

		}
		return nil
	}
	return filepath.Walk(root, walkfunc)
}
