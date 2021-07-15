package proto_reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultProtobufReader_Read(t *testing.T) {

	code := `
syntax = "proto3";
package services;

import "commons.proto";

option go_package = "integration";

enum RootEnum {
	ZERO = 0;
	ONE = 1;
	TWO = 2;
}

// The input of Add
message AddInput {
	// The first number to sum up
    int32 numberA = 1;
	// The second number to sum up
    int32 numberB = 2;

	enum MessageEnum {
		ZERO = 0;
		ONE = 1;
		TWO = 2;
	}
}

// The output of Add
message AddOutput {
	// The sum of the input
    int32 result = 1;
	map<string, int32> took = 2;
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

	rpc Destroy (commons.Empty) returns (commons.Empty);
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
	assertEnums(t, result)
}

func assertEnums(t *testing.T, result *ProtobufDefinition) {
	assert.NotEmpty(t, result.Enums)
	assert.EqualValues(t, 2, len(result.Enums))

	assert.EqualValues(t, "RootEnum", result.Enums["RootEnum"].Name)
	assert.EqualValues(t, "ZERO", result.Enums["RootEnum"].Items[0].Name)
	assert.EqualValues(t, 0, result.Enums["RootEnum"].Items[0].Index)
	assert.EqualValues(t, "ONE", result.Enums["RootEnum"].Items[1].Name)
	assert.EqualValues(t, 1, result.Enums["RootEnum"].Items[1].Index)
	assert.EqualValues(t, "TWO", result.Enums["RootEnum"].Items[2].Name)
	assert.EqualValues(t, 2, result.Enums["RootEnum"].Items[2].Index)

	assert.EqualValues(t, "MessageEnum", result.Enums["MessageEnum"].Name)
	assert.EqualValues(t, "ZERO", result.Enums["MessageEnum"].Items[0].Name)
	assert.EqualValues(t, 0, result.Enums["MessageEnum"].Items[0].Index)
	assert.EqualValues(t, "ONE", result.Enums["MessageEnum"].Items[1].Name)
	assert.EqualValues(t, 1, result.Enums["MessageEnum"].Items[1].Index)
	assert.EqualValues(t, "TWO", result.Enums["MessageEnum"].Items[2].Name)
	assert.EqualValues(t, 2, result.Enums["MessageEnum"].Items[2].Index)
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
