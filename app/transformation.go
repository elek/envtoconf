package app

import "fmt"

type Transformation func(map[string]string) string

func ToXml(keyvalues map[string]string) string {
	result := "<configuration>\n"
	for key, value := range keyvalues {
		result += fmt.Sprintf("<property><name>%s</name><value>%s</value></property>\n", key, value)
	}
	result += "</configuration>"
	return result
}

func ToEnv(keyvalues map[string]string) string {
	result := ""
	for key, value := range keyvalues {
		result += fmt.Sprintf("%s=%s\n", key, value)
	}
	result += ""
	return result
}

func ToSh(keyvalues map[string]string) string {
	result := ""
	for key, value := range keyvalues {
		result += fmt.Sprintf("export %s=%s\n", key, value)
	}
	result += ""
	return result
}

func ToProperties(keyvalues map[string]string) string {
	result := ""
	for key, value := range keyvalues {
		result += fmt.Sprintf("%s: %s\n", key, value)
	}
	result += ""
	return result
}

var transformations = map[string]Transformation{
	"xml": ToXml,
	"env": ToEnv,
	"cfg": ToEnv,
	"sh": ToSh,
	"conf": ToSh,
	"properties": ToProperties,
}