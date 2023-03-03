package domain_test

import (
	"testing"

	"github.com/kukymbr/kompasreader/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewFileType_WhenValid_ExpectNoError(t *testing.T) {
	validTypes := []int{5}

	for _, typeCode := range validTypes {
		filetype, err := domain.NewFileType(typeCode)

		assert.NoError(t, err)
		assert.Equal(t, domain.FileType(typeCode), filetype)
	}
}

func TestNewFileType_WhenInvalid_ExpectError(t *testing.T) {
	invalidTypes := []int{0, 1, 2}

	for _, typeCode := range invalidTypes {
		_, err := domain.NewFileType(typeCode)

		assert.ErrorIs(t, err, domain.ErrInvalidFileType)
	}
}
