package main

import (
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

func (field *Field) GetEnvironment(environment string) *Environment {
	env := field.Environment[environment]
	return &env
}

type Information struct {
	Fields []Field `yaml:"fields"`
}
type V1TrainerYaml struct {
	Version     string      `yaml:"version""`
	Information Information `yaml:"information"`
}

func NewV1TrainerYaml(file string) (*V1TrainerYaml, error) {
	y, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var trainer V1TrainerYaml
	err = yaml.Unmarshal(y, &trainer)
	if err != nil {
		return nil, err
	}
	return &trainer, nil
}

func main() {
	m1 := NewMapWithFile("application.yml")
	m2 := NewMapWithFile("application-test.yml")
	m3 := MergeMaps(&m1, &m2)
	trainer, err := NewV1TrainerYaml("trainer.yml")
	if err != nil {
		panic(err)
	}
	g := NewGenerator(trainer, "oc")
	yaml := ToYaml(g.Generate(m3))
	ioutil.WriteFile("result.yml", yaml, 0644)
}
