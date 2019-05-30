package internal

func parseProtoDefinition(result *ParseResult, definition *ProtobufDefinition) {
	optKey := encodeOptionKey(protoKey, "")
	options := parseOptions(result.Options[optKey])
	definition.Options = options
}
