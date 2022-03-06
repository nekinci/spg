package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Environment struct {
	Scheme string `yaml:"scheme"`
	Value  string `yaml:"value"`
}

type Field struct {
	Keys        []string               `yaml:"keys""`
	Type        string                 `yaml:"type"`
	Environment map[string]Environment `yaml:"environment"`
}

type Information struct {
	Fields []Field `yaml:"fields"`
}
type V1TrainerYaml struct {
	Version     string      `yaml:"version""`
	Information Information `yaml:"information"`
}

func main() {
	file, err := ioutil.ReadFile("application.yml")
	if err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	m1 := make(map[string]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		panic(err)
	}

	file1, err := ioutil.ReadFile("application-test.yml")
	err = yaml.Unmarshal(file1, &m1)
	if err != nil {
		panic(err)
	}
	m3 := mergeTwoMap(m, m1)
	_ = m3

	trainer := V1TrainerYaml{}
	file4, err := ioutil.ReadFile("trainer.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file4, &trainer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", trainer)
}

func processYaml(m3 map[string]interface{}, environment string, trainerYaml V1TrainerYaml) {

}

func mergeTwoMap(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	m3 := make(map[string]interface{})

	for k, v := range m1 {
		m3[k] = v
	}

	for k, v := range m2 {
		copyMap(&m3, k, v)
	}

	prettyPrintMap(m3, -1)
	return m3
}

func copyMap(m3 *map[string]interface{}, k string, v interface{}) {
	switch v.(type) {
	case string:
		(*m3)[k] = v
	case int:
		(*m3)[k] = v
	case bool:
		(*m3)[k] = v
	case map[string]interface{}:
		for k1, v1 := range v.(map[string]interface{}) {
			a := v.(map[string]interface{})
			copyMap(&a, k1, v1)
		}
		(*m3)[k] = v
	}
}

// TODO implement arrays...
func prettyPrintMap(m map[string]interface{}, deep int) {
	for k, v := range m {
		switch v.(type) {
		case string:
			printKeyValue(k, v, deep+1)
		case int:
			printKeyValue(k, v, deep+1)
		case bool:
			printKeyValue(k, v, deep+1)
		case map[string]interface{}:
			printKey(k, deep+1)
			prettyPrintMap(v.(map[string]interface{}), deep+1)
		}
	}
}

func printKey(k string, deep int) {
	for i := 0; i < deep; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%s:\n", k)
}

func printKeyValue(k string, v interface{}, deep int) {
	for i := 0; i < deep; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%s: %v\n", k, v)
}
