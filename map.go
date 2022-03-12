package main

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"reflect"
)

func NewMapWithFile(path string) map[string]interface{} {
	m := make(map[string]interface{})
	LoadFromFile(&m, path)
	return m
}

func NewMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func LoadFromFile(m *map[string]interface{}, path string) {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, m)
	if err != nil {
		panic(err)
	}

}

func ToYaml(m map[string]interface{}) []byte {
	var buff bytes.Buffer
	encoder := yaml.NewEncoder(&buff)
	encoder.SetIndent(2)
	err := encoder.Encode(m)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func MergeMaps(maps ...*map[string]interface{}) map[string]interface{} {
	result := NewMap()

	for _, m := range maps {
		for k, v := range *m {
			result[k] = determineValue(result[k], v)
		}
	}

	return result
}

func determineValue(oldValue, newValue interface{}) interface{} {

	if oldValue == nil && newValue != nil {
		return newValue
	}

	if oldValue != nil && newValue == nil {
		return oldValue
	}

	if oldValue == nil && newValue == nil {
		return nil
	}

	if isSliceOrArray(newValue) {
		return determineArrayValue(oldValue, newValue)
	}

	switch newValue.(type) {
	case string:
		return newValue
	case int:
		return newValue
	case float32:
		return newValue
	case float64:
		return newValue
	case int8:
		return newValue
	case int64:
		return newValue
	case bool:
		return newValue
	case map[string]interface{}:
		return determineMapValue(oldValue, newValue)
	default:
		return newValue
	}

}

func isSliceOrArray(value interface{}) bool {
	rt := reflect.TypeOf(value)
	return rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array
}

func toSlice(value interface{}) []interface{} {
	s := reflect.ValueOf(value)
	result := make([]interface{}, s.Len())
	if isSliceOrArray(value) {
		for i := 0; i < s.Len(); i++ {
			result[i] = s.Index(i).Interface()
		}
	}

	return result
}

func isMap(value interface{}) bool {
	rt := reflect.TypeOf(value)
	return rt.Kind() == reflect.Map
}

func toMap(value interface{}) map[string]interface{} {
	m := reflect.ValueOf(value)
	result := make(map[string]interface{})
	if isMap(value) {
		for _, k := range m.MapKeys() {
			result[k.String()] = m.MapIndex(k).Interface()
		}
	}

	return result
}

func determineArrayValue(oldValue, newValue interface{}) interface{} {
	//if isSliceOrArray(oldValue) {
	//	newArr := toSlice(newValue)
	//	result := toSlice(oldValue)
	//	oldIndex := len(result)
	//	for i, newV := range newArr {
	//		if i >= oldIndex {
	//			result = append(result, newV)
	//		} else {
	//			result[i] = determineValue(result[i], newV)
	//		}
	//	}
	//
	//	return result
	//}

	return toSlice(newValue)
}

func determineMapValue(oldValue, newValue interface{}) interface{} {

	if isMap(oldValue) {
		result := toMap(oldValue)
		newMap := toMap(newValue)
		for k, newV := range newMap {
			result[k] = determineValue(result[k], newV)
		}
		return result
	}

	return newValue

}

func PrettyPrint(m map[string]interface{}) string {
	b, err := yaml.Marshal(m)
	if err != nil {
		panic(err)
	}

	return string(b)
}
