package proto_reader

func NewReader() ProtobufReader {

	// Each read creates a new instance of the parser:
	factory := func() ProtobufParser {
		return NewProtobufParser()
	}

	return NewProtobufReader(factory)
}
