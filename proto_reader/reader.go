package proto_reader

type defaultProtobufReader struct {
	parserFactory ParserFactory
}

func NewProtobufReader(parserFactory ParserFactory) ProtobufReader {
	return &defaultProtobufReader{parserFactory: parserFactory}
}

func (this *defaultProtobufReader) Read(protoText string) (*ProtobufDefinition, error) {

	parser := this.parserFactory()
	result, err := parser.Parse(protoText)
	if err != nil {
		return nil, err
	}

	var out ProtobufDefinition
	parseMessages(result, &out)
	parseServices(result, &out)
	parseProtoDefinition(result, &out)

	return &out, nil
}
