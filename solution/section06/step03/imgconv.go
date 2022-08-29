package imgconv

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/tiff"
)

// Formatは画像形式を表す
type Format int

const (
	Unknown Format = iota
	PNG
	JPEG
	TIFF
)

// Stringは画像形式に対応する文字列を返す
func (f Format) String() string {
	switch f {
	case PNG:
		return "png"
	case JPEG:
		return "jpeg"
	case TIFF:
		return "tiff"
	}
	return "unknown"
}

// Setは文字列形式から対応する画像形式を設定する
func (f *Format) Set(s string) error {
	switch s {
	case "png":
		*f = PNG
	case "jpg", "jpeg":
		*f = JPEG
	case "tiff":
		*f = TIFF
	}
	return image.ErrFormat
}

// Extは画像形式に対応する拡張子を取得する
func (f Format) Ext() string {
	switch f {
	case PNG:
		return ".png"
	case JPEG:
		return ".jpeg"
	case TIFF:
		return ".tiff"
	}
	return ""
}

// FormatFromPathは指定したパスの拡張子から画像形式を取得する
func FormatFromPath(path string) Format {
	ext := filepath.Ext(path)
	switch strings.ToLower(ext) {
	case ".png":
		return PNG
	case ".jpeg", ".jpg":
		return JPEG
	case ".tiff":
		return TIFF
	}
	return Unknown
}

// ReplaceExtは拡張子を指定した形式のものに書き換える
func ReplaceExt(path string, f Format) string {
	if f == Unknown {
		return path
	}

	ext := filepath.Ext(path)
	i := len(path) - len(ext)
	return path[:i] + f.Ext()
}

// Encodeは画像をio.Readerからデータを読み込みimage.Image型に変換する
// 変換する画像形式はfで指定する
func Encode(w io.Writer, img image.Image, f Format) error {
	switch f {
	case PNG:
		return png.Encode(w, img)
	case JPEG:
		return jpeg.Encode(w, img, nil)
	case TIFF:
		return tiff.Encode(w, img, nil)
	}
	return image.ErrFormat
}

// Decodeはio.Readerからデータを読み込みimage.Image型に変換する
// 第2戻り値で画像の形式を返す
func Decode(r io.Reader) (image.Image, Format, error) {
	img, format, err := image.Decode(r)
	if err != nil {
		return nil, Unknown, err
	}

	switch format {
	case "png":
		return img, PNG, nil
	case "jpeg":
		return img, JPEG, nil
	case "tiff":
		return img, TIFF, nil
	}

	return nil, Unknown, image.ErrFormat
}

// ConvertAllは指定したディレクトリ以下の画像ファイルの変換を行う
// 変換前の形式をfromで変換後の形式をtoで指定する
func ConvertAll(root string, to, from Format) error {
	walkfunc := func(path string, info fs.FileInfo, err error) (rerr error) {

		// エラーが発生した
		if err != nil {
			return err
		}

		// ディレクトリ
		if info.IsDir() {
			return nil
		}

		// フォーマットが一致しない
		extFormat := FormatFromPath(path)
		if extFormat != from {
			return nil
		}

		// 変換前のファイルを開く
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		// 関数終了時にファイルを閉じる
		defer srcFile.Close()

		// 画像を読み込む
		img, format, err := Decode(srcFile)
		if err != nil {
			return err
		}

		// 拡張子だけ合ってて読み込んだ画像の形式が対象ではない
		if format != from {
			return nil
		}

		// 変換後に画像を保存するファイルを作成
		dstPath := ReplaceExt(path, to)
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}

		// 関数終了時にファイルを閉じる
		defer func() {
			// 閉じる際にエラーが発生した場合にはその値を返す
			// ただしすでにエラーが返されている場合(rerr != nil)は何もしない
			if err := dstFile.Close(); rerr == nil && err != nil {
				// 名前付き戻り値のrerrに代入することでwalkfuncの戻り値にする
				rerr = err
			}
		}()

		// フォーマットを指定して画像を保存する
		if err := Encode(dstFile, img, to); err != nil {
			return err
		}

		return nil
	}
	return filepath.Walk(root, walkfunc)
}
