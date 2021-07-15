package proto_reader

import (
	"fmt"
	"github.com/emicklei/proto"
)

type Option interface {
	Name() string
}

type ConstOption struct {
	Key   string
	Value string
}

type ArrayOption struct {
	Key   string
	Value []Option
}

type MapOption struct {
	Key   string
	Value map[string]Option
}

type Field struct {
	Name       string
	Type       string
	IsRepeated bool
	IsMap      bool
	SubType    *string
	Comment    string
	Options    []Option
}

func (this *Field) TypeString() string {
	if this.IsRepeated {
		return fmt.Sprintf("repeated %s", this.Type)
	} else if this.IsMap {
		return fmt.Sprintf("map<%s, %s>", this.Type, *this.SubType)
	} else {
		return this.Type
	}
}

type EnumEntry struct {
	Name  string
	Index int
}

type Enum struct {
	Name  string
	Items []EnumEntry
}

type Message struct {
	Name    string
	Fields  map[string]*Field
	Comment string
	Options []Option
}

type Method struct {
	Name            string
	Input           *Message
	Output          *Message
	StreamingInput  bool
	StreamingOutput bool
	Comment         string
	Options         []Option
}

type Service struct {
	Name    string
	Methods map[string]*Method
	Comment string
	Options []Option
}

type ProtobufDefinition struct {
	Messages map[string]*Message
	Services map[string]*Service
	Enums    map[string]*Enum
	Options  []Option
}

type ParseResult struct {
	Messages map[string]*proto.Message
	Services map[string]*proto.Service
	Methods  map[string]*proto.RPC
	Options  map[string][]*proto.Option
	Enums    map[string]*proto.Enum
	Imports  map[string]*proto.Import
}

type ProtobufParser interface {
	Parse(protoText string) (*ParseResult, error)
}

type ParserFactory = func() ProtobufParser

type ProtobufReader interface {
	Read(protoText string) (*ProtobufDefinition, error)
	ReadFile(protoFile, importPath string) (*ProtobufDefinition, error)
	ReadFileCustom(protoFile, importPath string, fileReader FileReader) (*ProtobufDefinition, error)
}

type FileReader interface {
	ReadAll(filePath string) string
}

func (this *ConstOption) Name() string {
	return this.Key
}

func (this *ArrayOption) Name() string {
	return this.Key
}

func (this *MapOption) Name() string {
	return this.Key
}
