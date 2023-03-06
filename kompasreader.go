package kompasreader

import (
	"fmt"
	"os"

	"github.com/kukymbr/kompasreader/domain"
	"github.com/kukymbr/kompasreader/internal/fileinfo"
	"github.com/kukymbr/kompasreader/internal/metainfo"
	"github.com/kukymbr/kompasreader/internal/zipper"
)

// New creates new KompasReader for the filepath
func New(filepath string) (reader *KompasReader, err error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, domain.ErrInvalidFilePath
	}

	return &KompasReader{filepath: filepath}, nil
}

// KompasReader reads Kompas3d file
type KompasReader struct {
	filepath string
	zipper   *zipper.Zipper
}

// Read reads specified Kompas file to the domain.Document instance
func (k *KompasReader) Read() (doc *domain.Document, err error) {
	doc = &domain.Document{}

	k.zipper, err = zipper.NewZipper(k.filepath)
	if err != nil {
		return nil, fmt.Errorf("init zipper: %w", err)
	}

	info, err := k.zipper.ReadTextFile(domain.FilenameFileInfo)
	if err != nil {
		return nil, fmt.Errorf("read fileinfo: %w", err)
	}

	doc.FileInfo, err = fileinfo.Unmarshall(info)
	if err != nil {
		return nil, fmt.Errorf("unmarshall fileinfo: %w", err)
	}

	doc.MetaInfo = &domain.MetaInfo{}

	{
		meta, err := k.zipper.ReadTextFile(domain.FilenameMetaInfo)
		if err != nil {
			return nil, fmt.Errorf("read meta info: %w", err)
		}

		metaUnm := metainfo.NewUnmarshaller(meta)
		doc.MetaInfo.SpcStructSections, err = metaUnm.Unmarshall()
		if err != nil {
			return nil, fmt.Errorf("unmarshall meta info: %w", err)
		}
	}

	return doc, nil
}
