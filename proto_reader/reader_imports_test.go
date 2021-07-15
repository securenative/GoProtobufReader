package proto_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultProtobufReader_ReadFile(t *testing.T) {
	reader := NewProtobufReader(func() ProtobufParser {
		return NewProtobufParser()
	})

	result, err := reader.ReadFile("./test_data/users.proto", "./test_data")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Services["Users"].Methods["Logout"].Output)
	assert.NotNil(t, result.Messages["Empty"])
}