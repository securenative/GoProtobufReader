package proto_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultProtobufReader_Read(t *testing.T) {

	code := `
syntax = "proto3";
package services;

option go_package = "integration";

// The input of Add
message AddInput {
	// The first number to sum up
    int32 numberA = 1;
	// The second number to sum up
    int32 numberB = 2;
}

// The output of Add
message AddOutput {
	// The sum of the input
    int32 result = 1;
}

// Service that adds two integers
service Adder {
	// Will take 2 numbers and sum them
    rpc Add (AddInput) returns (AddOutput) {
		option (google.api.http) = {
			post: "/api/v1/email"
			body: "*"
		};
	};
}
`

	reader := NewProtobufReader(func() ProtobufParser {
		return NewProtobufParser()
	})

	result, err := reader.Read(code)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, 2, len(result.Messages))
	assertMessages(t, result)
}

func assertMessages(t *testing.T, result *ProtobufDefinition) {
	assert.Equal(t, result.Messages["AddInput"].Name, "AddInput")
	assert.Equal(t, result.Messages["AddInput"].Comment, "The input of Add")
	assert.Equal(t, result.Messages["AddInput"].Fields["numberA"].Name, "numberA")
	assert.Equal(t, result.Messages["AddInput"].Fields["numberA"].Comment, "The first number to sum up")
	assert.False(t, result.Messages["AddInput"].Fields["numberA"].IsMap)
	assert.False(t, result.Messages["AddInput"].Fields["numberA"].IsRepeated)
	assert.Equal(t, result.Messages["AddInput"].Fields["numberA"].Type, "int32")
	assert.Nil(t, result.Messages["AddInput"].Fields["numberA"].SubType)
	assert.Equal(t, result.Messages["AddInput"].Fields["numberB"].Name, "numberB")
	assert.Equal(t, result.Messages["AddInput"].Fields["numberB"].Comment, "The second number to sum up")
	assert.False(t, result.Messages["AddInput"].Fields["numberB"].IsMap)
	assert.False(t, result.Messages["AddInput"].Fields["numberB"].IsRepeated)
	assert.Equal(t, result.Messages["AddInput"].Fields["numberB"].Type, "int32")
	assert.Nil(t, result.Messages["AddInput"].Fields["numberB"].SubType)
	assert.Equal(t, result.Messages["AddOutput"].Fields["result"].Name, "result")
	assert.Equal(t, result.Messages["AddOutput"].Fields["result"].Comment, "The sum of the input")
	assert.False(t, result.Messages["AddOutput"].Fields["result"].IsMap)
	assert.False(t, result.Messages["AddOutput"].Fields["result"].IsRepeated)
	assert.Equal(t, result.Messages["AddOutput"].Fields["result"].Type, "int32")
	assert.Nil(t, result.Messages["AddOutput"].Fields["result"].SubType)
}
