package kompasreader_test

import (
	"testing"

	"github.com/kukymbr/kompasreader"
	"github.com/kukymbr/kompasreader/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_ValidFile_ExpectNoError(t *testing.T) {
	filepath := "tests/example.spw"
	r, err := kompasreader.New(filepath)

	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestNew_InvalidFile_ExpectError(t *testing.T) {
	filepath := "tests/invalid_path.spw"
	r, err := kompasreader.New(filepath)

	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestKompasReader_Read_WhenValidFile_ExpectNoError(t *testing.T) {
	filepath := "tests/example.spw"
	reader, err := kompasreader.New(filepath)
	require.NoError(t, err)

	doc, err := reader.Read()

	require.NoError(t, err)
	require.NotNil(t, doc)

	assert.Equal(t, domain.FileTypeSpw, doc.FileInfo.FileType)
}
