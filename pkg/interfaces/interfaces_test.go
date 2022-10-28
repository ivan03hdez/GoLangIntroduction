package interfaces

import (
	"testing"

	mock_interfaces "golang/introduction/tests/mock/pkg/interfaces"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRunExample_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPersona := mock_interfaces.NewMockpersona(ctrl)

	var expectedErr error = nil
	var expectedValue bool = true

	mockPersona.EXPECT().IsMinor().Times(1).Return(true, nil)

	val, err := RunExampleTesting(mockPersona)

	assert.Equal(t, expectedValue, val)
	assert.Equal(t, expectedErr, err)
}
