package zipper_test

import (
	"io"
	"strings"
	"testing"

	"github.com/kukymbr/kompasreader/domain"
	"github.com/kukymbr/kompasreader/internal/zipper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewZipper_WhenInvalidFile_ExpectError(t *testing.T) {
	paths := []string{
		"../../testdata/invalid_path.spw",
		"../../testdata/invalid_file.spw",
	}

	for _, path := range paths {
		_, err := zipper.NewZipper(path)

		assert.Error(t, err)
	}
}

func TestZipper_OpenFile_WhenValidFile_ExpectNoError(t *testing.T) {
	filepath := "../../testdata/example.spw"
	files := []string{
		domain.FilenameFileInfo,
		domain.FilenameMetaInfo,
	}

	zip, err := zipper.NewZipper(filepath)
	require.NoError(t, err)

	for _, file := range files {
		f, err := zip.OpenFile(file)
		assert.NoError(t, err)

		err = f.Close()
		assert.NoError(t, err)
	}
}

func TestZipper_OpenFile_WhenInvalidFile_ExpectError(t *testing.T) {
	filepath := "../../testdata/example.spw"
	files := []string{
		"UnknownFile.1",
		"UnknownFile.2",
		"path/to/file",
		"../",
		"/",
	}

	zip, err := zipper.NewZipper(filepath)
	require.NoError(t, err)

	for _, file := range files {
		_, err := zip.OpenFile(file)
		assert.Error(t, err)
	}
}

func TestZipper_ReadTextFile_WhenValidFile_ExpectNoError(t *testing.T) {
	zip, err := zipper.NewZipper("../../testdata/example.spw")
	require.NoError(t, err)

	filesPrefixes := map[string]string{
		domain.FilenameFileInfo: "[FileInfo]",
		domain.FilenameMetaInfo: `<?xml version="1.0" encoding="utf-16"?>`,
	}

	for file, prefix := range filesPrefixes {
		reader, err := zip.ReadTextFile(file)
		assert.NoError(t, err)

		content, err := io.ReadAll(reader)
		assert.NoError(t, err)
		assert.True(t, strings.HasPrefix(string(content), prefix))
	}
}

func TestZipper_ReadTextFile_WhenInvalidFile_ExpectError(t *testing.T) {
	zip, err := zipper.NewZipper("../../testdata/example.spw")
	require.NoError(t, err)

	text, err := zip.OpenFile("unknown_file")
	assert.Error(t, err)
	assert.Empty(t, text)
}
