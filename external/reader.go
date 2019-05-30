package protobuf

import (
	"github.com/securenative/GoProtobufReader/internal"
)

func NewReader() internal.ProtobufReader {

	// Each read creates a new instance of the parser:
	factory := func() internal.ProtobufParser {
		return internal.NewProtobufParser()
	}

	return internal.NewProtobufReader(factory)
}
