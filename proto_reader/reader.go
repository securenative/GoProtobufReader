package proto_reader

import (
	"fmt"
	"io/ioutil"
	"path"
)

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
	parseEnums(result, &out)
	parseMessages(result, &out)
	parseServices(result, &out)
	parseProtoDefinition(result, &out)

	return &out, nil
}

func (this *defaultProtobufReader) ReadFile(protoFile, importPath string) (*ProtobufDefinition, error) {
	return this.ReadFileCustom(protoFile, importPath, &LocalFileReader{})
}

func (this *defaultProtobufReader) ReadFileCustom(protoFile, importPath string, fileReader FileReader) (*ProtobufDefinition, error) {
	result, err := this.parseRecursively(protoFile, importPath, fileReader, nil, "")
	if err != nil {
		return nil, err
	}

	var out ProtobufDefinition
	parseEnums(result, &out)
	parseMessages(result, &out)
	parseServices(result, &out)
	parseProtoDefinition(result, &out)

	return &out, nil
}

func (this *defaultProtobufReader) parseRecursively(protoFile string, importPath string, reader FileReader, parent *ParseResult, rootPkg string) (*ParseResult, error) {
	content := reader.ReadAll(protoFile)

	if importPath == "." {
		importPath = path.Dir(protoFile)
	}

	parser := this.parserFactory()
	result, err := parser.Parse(content)
	if err != nil {
		return nil, err
	}

	if parent == nil {
		rootPkg = result.Package
	}

	for childPath := range result.Imports {
		filePath := path.Join(importPath, childPath)
		if _, err := this.parseRecursively(filePath, importPath, reader, result, rootPkg); err != nil {
			return nil, err
		}
	}

	if parent != nil {
		for enumName, enum := range result.Enums {
			key := formatKey(result.Package, enumName, rootPkg)
			enum.Name = key
			parent.Enums[enumName] = enum
		}

		for msgName, msg := range result.Messages {
			key := formatKey(result.Package, msgName, rootPkg)
			msg.Name = key
			parent.Messages[key] = msg
		}
		return result, nil
	} else {
		return result, nil
	}
}

func formatKey(pkg string, name string, rootPkg string) string {
	if rootPkg == pkg {
		return name
	} else {
		return fmt.Sprintf("%s.%s", pkg, name)
	}
}

type LocalFileReader struct {
}

func (this *LocalFileReader) ReadAll(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}
