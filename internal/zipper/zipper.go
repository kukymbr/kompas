package zipper

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
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
func (z *Zipper) OpenFile(name string) (reader io.ReadCloser, err error) {
	if strings.Contains(name, "/") {
		return nil, errors.New("only root directory files requests in zip allowed")
	}

	return z.reader.Open(name)
}

// ReadTextFile reads content of the text file from the archive
func (z *Zipper) ReadTextFile(name string) (reader io.Reader, err error) {
	file, err := z.OpenFile(name)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Convert from UTF-16
	utf16enc := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	utf16bom := unicode.BOMOverride(utf16enc.NewDecoder())
	reader = transform.NewReader(bytes.NewReader(b), utf16bom)

	return reader, nil
}
