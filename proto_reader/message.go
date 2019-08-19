package proto_reader

import (
	"github.com/emicklei/proto"
)

func parseMessages(result *ParseResult, definition *ProtobufDefinition) {
	definition.Messages = make(map[string]*Message)
	for _, m := range result.Messages {
		var message Message
		message.Name = m.Name
		message.Comment = trimComment(m.Comment)
		message.Fields = parseFields(result, m.Elements)
		optKey := encodeOptionKey(messageKey, message.Name)
		message.Options = parseOptions(result.Options[optKey])

		definition.Messages[message.Name] = &message
	}
}

func parseFields(result *ParseResult, fields []proto.Visitee) map[string]*Field {
	out := make(map[string]*Field)
	for _, f := range fields {
		field := parseField(f)
		if field == nil {
			continue
		}
		key := encodeOptionKey(fieldKey, field.Name)
		options := parseOptions(result.Options[key])
		field.Options = options
		out[field.Name] = field
	}
	return out
}

func parseField(field proto.Visitee) *Field {
	switch field.(type) {
	case *proto.NormalField:
		casted := field.(*proto.NormalField)
		return &Field{
			Name:       casted.Name,
			Type:       casted.Type,
			IsRepeated: casted.Repeated,
			IsMap:      false,
			SubType:    nil,
			Comment:    trimComment(casted.Comment),
		}
	case *proto.MapField:
		casted := field.(*proto.MapField)
		return &Field{
			Name:       casted.Name,
			Type:       casted.KeyType,
			IsRepeated: false,
			IsMap:      true,
			SubType:    &casted.Type,
			Comment:    trimComment(casted.Comment),
		}
	case *proto.OneOfField:
		panic("OneOf fields are not supported")
	}

	return nil
}
