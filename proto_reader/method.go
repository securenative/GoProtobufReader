package proto_reader

import "github.com/emicklei/proto"

func parseMethods(result *ParseResult, definition *ProtobufDefinition) map[string]*Method {
	out := make(map[string]*Method)
	for _, v := range result.Methods {
		method := parseMethod(result, v, definition)
		out[method.Name] = method
	}
	return out
}

func parseMethod(result *ParseResult, method *proto.RPC, definition *ProtobufDefinition) *Method {
	var out Method
	out.Name = method.Name
	out.Comment = trimComment(method.Comment)
	out.Input = definition.Messages[method.RequestType]
	out.StreamingInput = method.StreamsRequest
	out.Output = definition.Messages[method.ReturnsType]
	out.StreamingOutput = method.StreamsReturns

	optKey := encodeOptionKey(methodKey, out.Name)
	out.Options = parseOptions(result.Options[optKey])

	return &out
}
