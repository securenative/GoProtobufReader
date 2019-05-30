package internal

import "github.com/emicklei/proto"

func parseServices(result *ParseResult, definition *ProtobufDefinition) {
	methods := parseMethods(result, definition)
	definition.Services = make(map[string]*Service)

	for _, s := range result.Services {
		var service Service
		service.Name = s.Name
		service.Comment = trimComment(s.Comment)

		service.Methods = make(map[string]*Method)
		for _, elem := range s.Elements {
			linkMethods(elem, methods, service.Methods)
		}

		optKey := encodeOptionKey(serviceKey, service.Name)
		options := parseOptions(result.Options[optKey])
		service.Options = options
		definition.Services[service.Name] = &service
	}
}

func linkMethods(element proto.Visitee, input map[string]*Method, output map[string]*Method) {
	switch element.(type) {
	case *proto.RPC:
		casted := element.(*proto.RPC)
		output[casted.Name] = input[casted.Name]
	}
}
