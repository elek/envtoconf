package app

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
	"gopkg.in/yaml.v2"
)

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

func renderYaml(yamlTree interface{}) ([]byte, error) {
	if result, err := yaml.Marshal(yamlTree); err != nil {
		return make([]byte, 0), err
	} else {
		return result, nil
	}
}

func ToYaml(keyvalues map[string]string) string {
	var yaml_struct  interface{}
	var parent_node interface{}
	yaml_struct = make(map[string]interface{})
	for key, value := range keyvalues {
		parts := strings.Split(key, ".")
		prev_part := ""
		node := yaml_struct
		parent_node = make(map[string]interface{})
		for _, part := range parts {
			//part is a digit
			if partInt, err := strconv.Atoi(part); err == nil {
				//actual node pointer points to a map
				if reflect.ValueOf(node).Kind() == reflect.Map {
					maked := make([]interface{}, 0)
					parent_node.(map[string]interface{})[prev_part] = maked
					node = maked
				}

				//it's an index, but there is not enough element
				for len(node.([]interface{})) < partInt {
					node = append(node.([]interface{}), make(map[string]interface{}))
					parent_node.(map[string]interface{})[prev_part] = node
				}

				//number
				parent_node = node
				node = node.([]interface{})[partInt - 1]
			} else {
				//part should be a string
				if _, ok := node.(map[string]interface{})[part]; !ok {
					node.(map[string]interface{})[part] = make(map[string]interface{})

				}
				parent_node = node
				node = node.(map[string]interface{})[part]
			}
			prev_part = part

		}
		if partInt, err := strconv.Atoi(parts[len(parts) - 1]); err == nil {
			parent_node.([]interface{})[partInt - 1] = value
			//if reflect.ValueOf(node).Kind() == reflect.Map {
			//	maked :=  make([]interface{}, 0)
			//	parent_node.(map[string]interface{})[prev_part] = maked
			//	node = maked
			//}
			//node = append(node.([]interface{}), value)
			//parent_node.(map[string]interface{})[prev_part] = node

		} else {
			parent_node.(map[string]interface{})[parts[len(parts) - 1]] = value
		}

	}
	if result, error := renderYaml(yaml_struct); error == nil {
		return string(result)
	} else {
		return "transformation.error: " + error.Error()
	}
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
	"yaml": ToYaml,
	"yml": ToYaml,
	"properties": ToProperties,
}
