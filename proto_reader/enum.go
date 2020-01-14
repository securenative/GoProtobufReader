package proto_reader

import "github.com/emicklei/proto"

func parseEnums(result *ParseResult, definition *ProtobufDefinition) {
	definition.Enums = make(map[string]*Enum)

	for k, v := range result.Enums {
		var items []EnumEntry
		for _, item := range v.Elements {
			field, ok := item.(*proto.EnumField)
			if ok {
				items = append(items, EnumEntry{
					Name:  field.Name,
					Index: field.Integer,
				})
			}
		}

		definition.Enums[k] = &Enum{
			Name:  v.Name,
			Items: items,
		}
	}
}
