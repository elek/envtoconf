package app

import (
	"strings"
	"errors"
)

type Configfile struct {
	Entries map[string]string
	File    string
	Format  string
}

// Parse a environment variable key according to the name convention.
func ParseKey(declaration string) (filename string, format string, key string, err error) {
	parts := strings.FieldsFunc(declaration, func(r rune) bool {
		return (r == '.' || r == '_')
	})
	if len(parts) == 0 {
		return "", "", "", errors.New("Line can't be parsed: " + declaration)
	}
	name := strings.ToLower(parts[0])
	extension := ""
	config_key := ""
	if len(parts) > 1 {
		extension = strings.ToLower(parts[1])
		if len(declaration) > len(name) + len(extension) + 2 {
			config_key = strings.TrimSpace(declaration[len(name) + len(extension) + 2:])
		}

		if strings.Contains(extension, "!") {
			splitted := strings.FieldsFunc(extension, func(r rune) bool {
				return r == '!'
			})
			extension = splitted[0]
			format = splitted[1]
			config_key = strings.TrimSpace(declaration[len(name) + len(extension) + len(format) + 3:])
		} else {
			format = extension
		}

	} else {
		format = extension
	}
	if _, ok := transformations[format]; extension != "" && ok {
		return name + "." + extension, format, config_key, nil
	} else {
		return "", "", "", errors.New("invalid key")
	}
}

func TransformToString(content map[string]string, format string) (string, error) {
	if val, ok := transformations[format]; ok {
		return val(content), nil
	} else {
		return "", errors.New("No such transformation " + format)
	}
}


//Parse all the available key value config
func ParseKeyValues(envs map[string]string) map[string]Configfile {
	var result = make(map[string]Configfile)
	for key, value := range envs {
		filename, format, key, err := ParseKey(key)
		if err == nil {
			if configfile, ok := result[filename]; ok {
				configfile.Entries[filename] = value
			} else {
				cfg := Configfile{File: filename, Format: format, Entries: map[string]string{key: value}}
				result[filename] = cfg
			}
		}
	}
	return result
}

