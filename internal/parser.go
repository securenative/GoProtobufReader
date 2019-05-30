package internal

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/emicklei/proto"
	"strings"
)

const fieldKey = "field"
const protoKey = "proto"
const enumKey = "enum"
const messageKey = "message"
const methodKey = "method"
const serviceKey = "service"

type defaultProtobufParser struct {
	consumed bool
	messages map[string]*proto.Message
	services map[string]*proto.Service
	methods  map[string]*proto.RPC
	options  map[string][]*proto.Option
	enums    map[string]*proto.Enum
}

func NewProtobufParser() ProtobufParser {
	return &defaultProtobufParser{
		consumed: false,
		messages: make(map[string]*proto.Message),
		services: make(map[string]*proto.Service),
		methods:  make(map[string]*proto.RPC),
		options:  make(map[string][]*proto.Option),
		enums:    make(map[string]*proto.Enum),
	}
}

func (this *defaultProtobufParser) addOption(parentType string, parentName string, option *proto.Option) {
	key := encodeOptionKey(parentType, parentName)
	optArr, found := this.options[key]
	if !found {
		optArr = make([]*proto.Option, 0)
	}

	optArr = append(optArr, option)

	this.options[key] = optArr
}

func (this *defaultProtobufParser) Parse(protoText string) (*ParseResult, error) {

	if this.consumed {
		return nil,
			errors.New("reader should be consumed only once, for each read you'll need to create a new reader")
	}

	this.consumed = true
	buffer := bytes.NewBuffer([]byte(protoText))
	parser := proto.NewParser(buffer)
	parseTree, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	proto.Walk(parseTree,
		proto.WithMessage(this.onMessage),
		proto.WithRPC(this.onMethod),
		proto.WithService(this.onService),
		proto.WithEnum(this.onEnum),
		proto.WithOption(this.onOption),
	)

	return &ParseResult{
		Methods:  this.methods,
		Messages: this.messages,
		Enums:    this.enums,
		Options:  this.options,
		Services: this.services,
	}, nil
}

func (this *defaultProtobufParser) onMessage(message *proto.Message) {
	this.messages[message.Name] = message
}

func (this *defaultProtobufParser) onMethod(method *proto.RPC) {
	this.methods[method.Name] = method
}

func (this *defaultProtobufParser) onService(service *proto.Service) {
	this.services[service.Name] = service
}

func (this *defaultProtobufParser) onOption(option *proto.Option) {
	var name string
	var kind string

	switch option.Parent.(type) {
	case *proto.Enum:
		name = option.Parent.(*proto.Enum).Name
		kind = enumKey
	case *proto.Message:
		name = option.Parent.(*proto.Message).Name
		kind = messageKey
	case *proto.RPC:
		name = option.Parent.(*proto.RPC).Name
		kind = methodKey
	case *proto.Service:
		name = option.Parent.(*proto.Service).Name
		kind = serviceKey
	case *proto.MapField:
		name = option.Parent.(*proto.MapField).Name
		kind = fieldKey
	case *proto.NormalField:
		name = option.Parent.(*proto.NormalField).Name
		kind = fieldKey
	case *proto.OneOfField:
		name = option.Parent.(*proto.OneOfField).Name
		kind = fieldKey
	case *proto.EnumField:
		name = option.Parent.(*proto.EnumField).Name
		kind = fieldKey
	case *proto.Proto:
		name = ""
		kind = protoKey
	}

	this.addOption(kind, name, option)
}

func (this *defaultProtobufParser) onEnum(enum *proto.Enum) {
	this.enums[enum.Name] = enum
}

func encodeOptionKey(kind string, name string) string {
	return fmt.Sprintf("%s-%s", kind, name)
}

func decodeOptionKey(key string) (kind string, name string, err error) {
	arr := strings.Split(key, "-")
	if arr == nil || len(arr) != 2 {
		return "", "", errors.New("cannot split the key")
	}

	return arr[0], arr[1], nil
}
