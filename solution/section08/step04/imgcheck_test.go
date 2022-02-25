package imgcheck_test

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	imgcheck "github.com/gohandson/toybox-ja/solution/section08/step04"
	"go.uber.org/multierr"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		file  string
		rules []imgcheck.Rule

		wantErrs []error
	}{
		{"format ok", "72x72.png", []imgcheck.Rule{imgcheck.Format("png")}, nil},
		{"format ng", "72x72.png", []imgcheck.Rule{imgcheck.Format("jpeg")}, []error{imgcheck.ErrFormat}},
		{"format-pattern ok", "72x72.png", []imgcheck.Rule{formatPattern(t, ".+g")}, nil},
		{"format-pattern ng", "72x72.png", []imgcheck.Rule{formatPattern(t, "jpe?g")}, []error{imgcheck.ErrFormat}},
		{"height ok", "72x72.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, nil},
		{"width ok", "72x72.png", []imgcheck.Rule{imgcheck.MaxWidth(72)}, nil},
		{"height ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, []error{imgcheck.ErrTooLarge}},
		{"width ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxWidth(72)}, []error{imgcheck.ErrTooLarge}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(filepath.Join("testdata", tt.file))
			if err != nil {
				t.Fatal("予期しないエラー:", err)
			}
			t.Cleanup(func() { f.Close() })

			err = imgcheck.Validate(f, tt.rules...)
			errs := multierr.Errors(err)
			if len(tt.wantErrs) != len(errs) {
				t.Fatalf("予期したエラーの数が異なります:want %d got %d", len(tt.wantErrs), len(errs))
			}

			for i := range errs {
				var verr *imgcheck.ValidationError
				if !errors.As(errs[i], &verr) || verr.Rule != tt.rules[i] {
					t.Errorf("期待したルールのエラーではありません: i = %d", i)
				}

				if !errors.Is(errs[i], tt.wantErrs[i]) {
					t.Errorf("予期したエラーと異なります: want %v got %v", tt.wantErrs[i], errs[i])
				}
			}
		})
	}
}

func formatPattern(t *testing.T, pattern string) imgcheck.Rule {
	t.Helper()
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatal("予期しないエラー:", err)
	}
	return imgcheck.FormatPattern(re)
}
