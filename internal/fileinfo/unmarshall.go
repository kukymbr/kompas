package fileinfo

import (
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/kukymbr/kompasreader/domain"
	"gopkg.in/ini.v1"
)

// Unmarshall reads ini data from reader and converts it to domain.FileInfo instance
func Unmarshall(configRaw string) (info *domain.FileInfo, err error) {
	inidoc, err := ini.Load(strings.NewReader(configRaw))
	if err != nil {
		return nil, err
	}

	data := &fileInfoRAW{}
	err = inidoc.Section("FileInfo").MapTo(&data)
	if err != nil {
		return nil, err
	}

	info, err = data.createFileInfo()
	if err != nil {
		return nil, err
	}

	return info, nil
}

type fileInfoRAW struct {
	FileType   int    `ini:"FileType"`
	Author     string `ini:"Author"`
	Comment    string `ini:"Comment"`
	AppVersion string `ini:"AppVersion"`
	CreateData string `ini:"CreateData"`
	ModifyData string `ini:"ModifyData"`
}

func (data *fileInfoRAW) createFileInfo() (info *domain.FileInfo, err error) {
	info = &domain.FileInfo{
		Author:  data.Author,
		Comment: data.Comment,
	}

	filetype, err := domain.NewFileType(data.FileType)
	if err != nil {
		return nil, err
	}
	info.FileType = filetype

	data.AppVersion = strings.ToLower(strings.ReplaceAll(data.AppVersion, "KOMPAS_", ""))
	appVersion, err := version.NewVersion(data.AppVersion)
	if err != nil {
		return nil, err
	}
	info.AppVersion = *appVersion.Core()

	date, err := time.Parse("02.01.2006 15:04:05", data.CreateData)
	if err == nil {
		info.CreatedAt = date
	}

	date, err = time.Parse("02.01.2006 15:04:05", data.ModifyData)
	if err == nil {
		info.UpdatedAt = date
	}

	return info, nil
}
