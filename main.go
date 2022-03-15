package main

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
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
	Fields         []Field          `yaml:"fields"`
	AbsoluteConfig []AbsoluteConfig `yaml:"absolute-configs"`
}

type AbsoluteConfig struct {
	Key         string                 `yaml:"config-key"`
	Environment map[string]interface{} `yaml:"environment"`
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
	Execute()
}

func RunGenerate(args []string, profile, output string) {
	var mlist []*map[string]interface{}

	for _, v := range args {
		m := NewMapWithFile(v)
		mlist = append(mlist, &m)
	}

	m3 := MergeMaps(mlist...)
	homeDir := getHomeDir()
	trainer, err := NewV1TrainerYaml(homeDir + "/.spg/config.yml")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	g := NewGenerator(trainer, profile)
	yaml := ToYaml(g.Generate(m3))
	ioutil.WriteFile(output, yaml, 0644)
}

func HandleConfig(action, path string) {
	if action == "set" {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(1)
		}
		var config interface{}
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			fmt.Printf("Error parsing file: %v. It may be not yaml file.", err)
			os.Exit(1)
		}

		homeDir := getHomeDir()

		err = os.MkdirAll(homeDir+"/.spg", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v", err)
			os.Exit(1)
		}

		configPath := homeDir + "/.spg/config.yml"
		err = ioutil.WriteFile(configPath, file, 0644)
		if err != nil {
			fmt.Printf("Error saving file: %v", err)
			os.Exit(1)
		}

	} else if action == "unset" {
		homeDir := getHomeDir()
		configPath := homeDir + "/.spg/config.yml"
		err := os.Remove(configPath)
		if err != nil {
			fmt.Printf("Error removing config: %v", err)
			os.Exit(1)
		}
	} else if action == "print" {
		homeDir := getHomeDir()
		configPath := homeDir + "/.spg/config.yml"
		m := NewMapWithFile(configPath)
		fmt.Println(Pretty(m))
	}
}

func getHomeDir() string {
	h, err := homedir.Dir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v", err)
		os.Exit(1)
	}
	return h
}
