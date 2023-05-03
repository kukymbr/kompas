package kompas

import (
	"time"

	"github.com/hashicorp/go-version"
)

// FileInfo is a data from the Kompas' FileInfo ini file
type FileInfo struct {
	Author     string
	Comment    string
	FileType   FileType
	AppVersion version.Version
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
