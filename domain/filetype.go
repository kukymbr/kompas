package domain

// Kompas file types
const (
	FileTypeSpw = FileType(5)
	FileTypeCdw = FileType(0)
)

var typesEnabled = map[FileType]bool{
	FileTypeSpw: true,
	FileTypeCdw: false,
}

// NewFileType creates new FileType if typeCode is valid
func NewFileType(typeCode int) (filetype FileType, err error) {
	filetype = FileType(typeCode)
	if !filetype.IsValid() {
		return 0, ErrInvalidFileType
	}

	return filetype, nil
}

// FileType is a Kompas 3d file type
type FileType int

// IsValid checks if file type is valid and enabled
func (t FileType) IsValid() bool {
	enabled, ok := typesEnabled[t]

	return ok && enabled
}
