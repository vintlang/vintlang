package module

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var XMLFunctions = map[string]object.ModuleFunction{}

func init() {
	XMLFunctions["escape"] = xmlEscape
	XMLFunctions["unescape"] = xmlUnescape
	XMLFunctions["validate"] = xmlValidate
	XMLFunctions["extract"] = xmlExtractValue
}

func xmlEscape(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"xml", "escape",
			"1 argument: text (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`xml.escape("<tag>content</tag>") -> "&lt;tag&gt;content&lt;/tag&gt;"`,
		)
	}

	data := args[0]
	if data.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"xml", "escape",
			"string text",
			string(data.Type()),
			`xml.escape("<tag>content</tag>") -> "&lt;tag&gt;content&lt;/tag&gt;"`,
		)
	}

	input := data.(*object.String).Value
	var buf strings.Builder
	xml.EscapeText(&buf, []byte(input))
	
	return &object.String{Value: buf.String()}
}

func xmlUnescape(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"xml", "unescape",
			"1 argument: escaped text (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`xml.unescape("&lt;tag&gt;content&lt;/tag&gt;") -> "<tag>content</tag>"`,
		)
	}

	data := args[0]
	if data.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"xml", "unescape",
			"string escaped text",
			string(data.Type()),
			`xml.unescape("&lt;tag&gt;content&lt;/tag&gt;") -> "<tag>content</tag>"`,
		)
	}

	input := data.(*object.String).Value
	
	// Replace common XML entities
	replacements := map[string]string{
		"&lt;":   "<",
		"&gt;":   ">",
		"&amp;":  "&",
		"&quot;": "\"",
		"&apos;": "'",
	}
	
	result := input
	for entity, char := range replacements {
		result = strings.ReplaceAll(result, entity, char)
	}
	
	return &object.String{Value: result}
}

func xmlValidate(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"xml", "validate",
			"1 argument: XML string (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`xml.validate("<root><child>value</child></root>") -> true/false`,
		)
	}

	data := args[0]
	if data.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"xml", "validate",
			"string XML data",
			string(data.Type()),
			`xml.validate("<root><child>value</child></root>") -> true/false`,
		)
	}

	input := data.(*object.String).Value
	
	// Try to parse the XML
	var v interface{}
	err := xml.Unmarshal([]byte(input), &v)
	
	if err != nil {
		return &object.Boolean{Value: false}
	}
	
	return &object.Boolean{Value: true}
}

func xmlExtractValue(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"xml", "extract",
			"2 arguments: XML string (string), tag name (string)",
			fmt.Sprintf("%d arguments", len(args)),
			`xml.extract("<root><name>John</name></root>", "name") -> "John"`,
		)
	}

	xmlData := args[0]
	tagName := args[1]
	
	if xmlData.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"xml", "extract",
			"string XML data for first argument",
			string(xmlData.Type()),
			`xml.extract("<root><name>John</name></root>", "name") -> "John"`,
		)
	}
	
	if tagName.Type() != object.STRING_OBJ {
		return ErrorMessage(
			"xml", "extract",
			"string tag name for second argument",
			string(tagName.Type()),
			`xml.extract("<root><name>John</name></root>", "name") -> "John"`,
		)
	}

	xmlStr := xmlData.(*object.String).Value
	tag := tagName.(*object.String).Value
	
	// Simple XML value extraction using string parsing
	startTag := fmt.Sprintf("<%s>", tag)
	endTag := fmt.Sprintf("</%s>", tag)
	
	startIdx := strings.Index(xmlStr, startTag)
	if startIdx == -1 {
		return &object.String{Value: ""}
	}
	
	startIdx += len(startTag)
	endIdx := strings.Index(xmlStr[startIdx:], endTag)
	if endIdx == -1 {
		return &object.String{Value: ""}
	}
	
	value := xmlStr[startIdx : startIdx+endIdx]
	return &object.String{Value: value}
}