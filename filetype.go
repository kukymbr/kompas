package kompas

const (
	// FileTypeSpw is an .spw file type
	FileTypeSpw = FileType(5)
)

var typesAvailable = map[FileType]struct{}{
	FileTypeSpw: {},
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
	_, ok := typesAvailable[t]

	return ok
}
