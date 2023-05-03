package fileinfo_test

import (
	"github.com/kukymbr/kompas"
	"strings"
	"testing"

	"github.com/kukymbr/kompas/internal/fileinfo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshal_WhenValidData_ExpectNoError(t *testing.T) {
	cfg := strings.NewReader(`
[FileInfo]
AppName=КОМПАС-3D V16.1
AppVersion=KOMPAS_16.1
BuildNum=1971_revK_124444_revM_101135
AppPlatform=x64
MathFileVersion=0x10001002
AppFileVersion=0x10001006
FileType=5
CreateAppVersion=0x10001006
CreateData=05.10.2022 14:05:09
ModifyData=21.10.2022 09:52:05
Author=Author Name
OrgName=
Comment=Test Comment
AutoSave=false
`)

	info, err := fileinfo.Unmarshal(cfg)

	require.NoError(t, err)
	assert.Equal(t, kompas.FileTypeSpw, info.FileType)
	assert.Equal(t, "16.1.0", info.AppVersion.String())
	assert.Equal(t, "Author Name", info.Author)
	assert.Equal(t, "Test Comment", info.Comment)
	assert.Equal(t, "20221005_140509", info.CreatedAt.Format("20060102_150405"))
	assert.Equal(t, "20221021_095205", info.UpdatedAt.Format("20060102_150405"))
}

func TestUnmarshal_WhenInvalidData_ExpectError(t *testing.T) {
	invalid := []string{
		`AppVersion=KOMPAS_16.1
FileType=5
CreateData=05.10.2022 14:05:09
ModifyData=21.10.2022 09:52:05
Author=Author Name
Comment=Test Comment`,
		`[FileInfo]
AppVersion=
FileType=5
CreateData=05.10.2022 14:05:09
ModifyData=21.10.2022 09:52:05
Author=Author Name
Comment=Test Comment`,
		`[FileInfo]
AppVersion=KOMPAS_16.1
FileType=1000
CreateData=05.10.2022 14:05:09
ModifyData=21.10.2022 09:52:05
Author=Author Name
Comment=Test Comment`,
	}

	for _, content := range invalid {
		cfg := strings.NewReader(content)
		info, err := fileinfo.Unmarshal(cfg)

		assert.Error(t, err)
		assert.Nil(t, info)
	}
}
