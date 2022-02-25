package imgcheck

import "errors"

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

func (err *ValidationError) Unwrap() error {
	return err.Err
}
