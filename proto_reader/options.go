package proto_reader

import (
	"fmt"
	"github.com/emicklei/proto"
)

func parseOptions(options []*proto.Option) []Option {
	out := make([]Option, len(options))
	for idx, o := range options {
		out[idx] = parseLiteral(o.Name, &o.Constant)
	}
	return out
}

func parseLiteral(optionName string, literal *proto.Literal) Option {
	if literal.Array != nil {
		return parseArray(optionName, literal.Array)
	} else if literal.OrderedMap != nil {
		return parseMap(optionName, literal.OrderedMap)
	}
	return &ConstOption{
		Key:   optionName,
		Value: literal.Source,
	}
}

func parseArray(name string, literals []*proto.Literal) Option {
	out := ArrayOption{Key: name}
	for idx, l := range literals {
		out.Value = append(out.Value, parseLiteral(fmt.Sprint(idx), l))
	}
	return &out
}

func parseMap(name string, literals proto.LiteralMap) Option {
	out := MapOption{Key: name, Value: make(map[string]Option)}
	for _, l := range literals {
		out.Value[l.Name] = parseLiteral(l.Name, l.Literal)
	}
	return &out
}
