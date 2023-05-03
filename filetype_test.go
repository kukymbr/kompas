package kompas_test

import (
	"github.com/kukymbr/kompas"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileType_WhenValid_ExpectNoError(t *testing.T) {
	validTypes := []int{5}

	for _, typeCode := range validTypes {
		filetype, err := kompas.NewFileType(typeCode)

		assert.NoError(t, err)
		assert.Equal(t, kompas.FileType(typeCode), filetype)
	}
}

func TestNewFileType_WhenInvalid_ExpectError(t *testing.T) {
	invalidTypes := []int{0, 1, 2}

	for _, typeCode := range invalidTypes {
		_, err := kompas.NewFileType(typeCode)

		assert.ErrorIs(t, err, kompas.ErrInvalidFileType)
	}
}
