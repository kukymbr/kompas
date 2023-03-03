package zipper

import (
	"archive/zip"
	"errors"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
)

// NewZipper create Zipper instance for the specified zip archive file path
func NewZipper(zipFilePath string) (zipper *Zipper, err error) {
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return nil, err
	}

	return &Zipper{reader: reader}, nil
}

// Zipper is a struct to operate with Kompas' zip file
type Zipper struct {
	reader *zip.ReadCloser
}

func (z *Zipper) Close() error {
	return z.reader.Close()
}

// OpenFile opens file from the zip archive by its name
func (z *Zipper) OpenFile(name string) (reader io.Reader, err error) {
	if strings.Contains(name, "/") {
		return nil, errors.New("only root directory files requests in zip allowed")
	}

	return z.reader.Open(name)
}

// ReadTextFile reads content of the text file from the archive
func (z *Zipper) ReadTextFile(name string) (text string, err error) {
	file, err := z.OpenFile(name)
	if err != nil {
		return "", err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	text = strings.ToValidUTF8(string(b), "")

	text = strings.Map(func(rune rune) rune {
		if rune == utf8.RuneError {
			return -1
		}

		if unicode.IsPrint(rune) || unicode.IsSpace(rune) {
			return rune
		}

		return -1
	}, text)

	return text, nil
}
