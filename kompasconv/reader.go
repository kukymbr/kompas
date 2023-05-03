package kompasconv

import (
	"fmt"
	"github.com/kukymbr/kompas"
	"github.com/kukymbr/kompas/internal/fileinfo"
	"github.com/kukymbr/kompas/internal/metainfo"
	"github.com/kukymbr/kompas/internal/zipper"
	"os"
)

// NewReader creates new Reader for the filepath
func NewReader(filepath string) (reader *Reader, err error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, kompas.ErrInvalidFilePath
	}

	return &Reader{filepath: filepath}, nil
}

// Reader reads Kompas file
type Reader struct {
	filepath string
	zipper   *zipper.Zipper
}

// Read reads specified Kompas file to the domain.Document instance
func (k *Reader) Read() (doc *kompas.Document, err error) {
	doc = &kompas.Document{}

	k.zipper, err = zipper.NewZipper(k.filepath)
	if err != nil {
		return nil, fmt.Errorf("init zipper: %w", err)
	}

	info, err := k.zipper.ReadTextFile(filenameFileInfo)
	if err != nil {
		return nil, fmt.Errorf("read fileinfo: %w", err)
	}

	doc.FileInfo, err = fileinfo.Unmarshal(info)
	if err != nil {
		return nil, fmt.Errorf("unmarshal fileinfo: %w", err)
	}

	doc.MetaInfo = &kompas.MetaInfo{}

	{
		meta, err := k.zipper.ReadTextFile(filenameMetaInfo)
		if err != nil {
			return nil, fmt.Errorf("read meta info: %w", err)
		}

		metaUnm := metainfo.NewUnmarshaler(meta)
		doc.MetaInfo.SpcStructSections, err = metaUnm.Unmarshal()
		if err != nil {
			return nil, fmt.Errorf("unmarshal meta info: %w", err)
		}
	}

	return doc, nil
}
