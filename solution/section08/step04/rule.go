package imgcheck

import (
	"fmt"
	"image"
	"regexp"

	"go.uber.org/multierr"
)

// バリデーションルールを表すインタフェース
type Rule interface {
	Validate(img image.Image, format string) error
}

// 画像フォーマットをチェックするルール
func Format(format string) (rule Rule) {
	return &formatRule{format: format}
}

// 画像フォーマットを正規表現でチェックするルール
func FormatPattern(pattern *regexp.Regexp) (rule Rule) {
	return &formatRule{pattern: pattern}
}

// 高さをチェックするルール
func MaxHeight(h int) Rule {
	return &maxSizeRule{height: &h}
}

// 幅をチェックするルール
func MaxWidth(w int) Rule {
	return &maxSizeRule{width: &w}
}

type formatRule struct {
	format  string
	pattern *regexp.Regexp
}

func (r *formatRule) Validate(img image.Image, format string) error {
	if r.pattern != nil {
		if !r.pattern.MatchString(format) {
			return &ValidationError{
				Rule: r,
				Err:  fmt.Errorf("期待%s 実際%s: %w", r.pattern, format, ErrFormat),
			}
		}
		return nil
	}

	if r.format != format {
		return &ValidationError{
			Rule: r,
			Err:  fmt.Errorf("期待%s 実際%s: %w", r.format, format, ErrFormat),
		}
	}
	return nil
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
			Err:  fmt.Errorf("期待する幅%d 実際%d: %w", *r.width, bounds.Dx(), ErrTooLarge),
		})

	}

	return err
}
