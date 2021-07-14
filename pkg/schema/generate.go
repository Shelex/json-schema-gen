package schema

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var nullValueSchema = map[string]interface{}{
	"type": "null",
}

type Schema struct {
	App   string
	Title string
	Data  []byte
}

func NewSchema(data map[string]interface{}, path string) (Schema, error) {
	app, filename := splitFilePath(path)
	s := Schema{App: app, Title: filename}
	if err := s.Generate(data); err != nil {
		return s, err
	}
	return s, nil
}

func splitFilePath(file string) (string, string) {
	absolutePath, file := filepath.Split(file)
	ext := filepath.Ext(file)
	fileName := file[0 : len(file)-len(ext)]
	folders := strings.Split(absolutePath, string(os.PathSeparator))
	var app string
	for i, folder := range folders {
		// find root app folder inside cypress/fixtures/
		if folder == "cypress" && folders[i+1] == "fixtures" {
			app = folders[i+2]
		}
	}
	return app, fileName
}

// Generate - wrap schema in js keywords and add imports if needed
func (s *Schema) Generate(response map[string]interface{}) error {
	if len(response) == 0 {
		return errors.New("json is empty")
	}
	schema := parseJSON(response)
	schema["title"] = s.Title
	jsonBytes, err := json.MarshalIndent(schema, "", "\t")
	if err != nil {
		return err
	}
	jsonString := string(jsonBytes)
	s.Data = []byte(jsonString)
	return nil
}

// parseJSON - convert parsed json struct to json schema struct
func parseJSON(input interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	switch t := input.(type) {
	// golang always return numbers in interface as float, so integer check handled here
	case float64:
		res["type"] = numberHandler(input)
		return res
	case string:
		s := input.(string)
		res["type"] = "string"
		format := formatChecker(s)
		if format != s {
			res["format"] = format
		}
		return res
	case bool:
		res["type"] = "boolean"
		return res
	// handle json array
	case []interface{}:
		res["type"] = "array"
		if arrayValue, ok := input.([]interface{}); ok {
			res["items"] = arrayHandler(arrayValue)
		}
		return res
	// handle json object
	case map[string]interface{}:
		res["type"] = "object"
		if objectValue, ok := input.(map[string]interface{}); ok {
			res["properties"] = objectHandler(objectValue)
			mapProperties := reflect.ValueOf(res["properties"])
			if mapProperties.Kind() == reflect.Map {
				mapKeys := mapProperties.MapKeys()
				keys := make([]string, len(mapKeys))
				for index, key := range mapKeys {
					keys[index] = key.String()
				}
				res["required"] = keys
			}
		}
		return res
	default:
		res["type"] = t
		return res
	}
}

func arrayHandler(array []interface{}) interface{} {
	if len(array) == 0 {
		return nullValueSchema
	}
	var item interface{}
	var max int
	for _, v := range array {
		temporarySchema := parseJSON(v)
		marshalled, err := json.Marshal(temporarySchema)
		if err != nil {
			return nil
		}
		size := len(marshalled)
		if size > max {
			max = size
			item = temporarySchema
		}
	}
	return item
}

func objectHandler(object map[string]interface{}) map[string]interface{} {
	props := map[string]interface{}{}
	for key, value := range object {
		props[key] = parseJSON(value)
	}
	if len(props) == 0 {
		return nullValueSchema
	}
	return props
}

func numberHandler(input interface{}) string {
	if math.Mod(input.(float64), 1.0) == 0 {
		return "integer"
	}
	return "number"
}
