package main

import (
	"fmt"
	"os"
	"strings"
)

type Generator struct {
	Trainer     *V1TrainerYaml
	environment string
	currentMap  map[string]interface{}
}

type Url struct {
	Url string
}

func NewUrl(url interface{}) Url {
	switch url.(type) {
	case string:
		return Url{
			Url: url.(string),
		}
	case Environment:
		return Url{
			Url: fmt.Sprintf("%s://%s", url.(Environment).Scheme, url.(Environment).Value),
		}
	case *Environment:
		return Url{
			Url: fmt.Sprintf("%s://%s", url.(*Environment).Scheme, url.(*Environment).Value),
		}

	default:
		panic("unsupported url type")
	}
}

func (u Url) String() string {
	return u.Url
}

func (u Url) WithoutScheme() string {
	if u.Url == "" {
		return ""
	}

	httpContains := strings.Contains(u.Url, "http://")
	if httpContains {
		return strings.Replace(u.Url, "http://", "", 1)
	}

	httpsContains := strings.Contains(u.Url, "https://")
	if httpsContains {
		return strings.Replace(u.Url, "https://", "", 1)
	}

	return u.Url
}

func (u Url) Scheme() string {
	if u.Url == "" {
		return ""
	}

	httpContains := strings.Contains(u.Url, "http://")
	if httpContains {
		return "http"
	}

	httpsContains := strings.Contains(u.Url, "https://")
	if httpsContains {
		return "https"
	}

	// TODO: think about this, maybe we return empty string
	return "http"
}

func (u Url) Hostname() string {
	if u.Url == "" {
		return ""
	}

	withoutScheme := u.WithoutScheme()
	withoutScheme = strings.Split(withoutScheme, "/")[0]
	return withoutScheme
}

func (u Url) Path() string {
	if u.Url == "" {
		return ""
	}

	withoutScheme := u.WithoutScheme()
	index := strings.Index(withoutScheme, "/")
	if index == -1 {
		return ""
	}

	return withoutScheme[index:]
}

func NewGenerator(trainer *V1TrainerYaml, environment string) *Generator {
	checkEnvironment(trainer, environment)
	return &Generator{
		Trainer:     trainer,
		environment: environment,
	}
}

func (g *Generator) Generate(m map[string]interface{}) map[string]interface{} {
	g.currentMap = m
	m = g.GenerateForFields("", m)
	m = g.GenerateForAbsoluteConfig("", m)
	return m
}

func (g *Generator) GenerateForAbsoluteConfig(key string, m map[string]interface{}) map[string]interface{} {

	for k, v := range m {
		kk := getKey(key, k)
		switch v.(type) {
		case map[string]interface{}:
			m[k] = g.GenerateForAbsoluteConfig(kk, v.(map[string]interface{}))
		case []interface{}:
			m[k] = g.GenerateForAbsoluteConfigForArray(kk, v.([]interface{}))
		case interface{}:
			m[k] = g.decideConfigValue(kk, v)
		default:
			m[k] = v
		}
	}

	return m
}

func (g *Generator) GenerateForAbsoluteConfigValue(k string, v interface{}) {

}

func (g *Generator) GenerateForAbsoluteConfigForArray(k string, arr []interface{}) []interface{} {

	for index, v := range arr {
		key := fmt.Sprintf("%s[%d]", k, index)
		switch v.(type) {
		case map[string]interface{}:
			arr[index] = g.GenerateForAbsoluteConfig(key, (v).(map[string]interface{}))
		case []interface{}:
			arr[index] = g.GenerateForAbsoluteConfigForArray(key, v.([]interface{}))
		case interface{}:
			arr[index] = g.decideConfigValue(key, v)

		}
	}
	return arr
}

func getKey(key string, k string) string {
	if key == "" {
		return k
	}
	return key + "." + k
}

func (g *Generator) decideConfigValue(k string, v interface{}) interface{} {
	config := g.getConfig(k)
	if config == nil {
		return v
	}

	if config.Condition == nil {
		vii := config.Environment[g.environment]
		return vii
	} else {
		cond := *config.Condition
		if g.getConditionResult(cond) {
			vii := config.Environment[g.environment]
			return vii
		}
		return v
	}

}

func (g *Generator) getConditionResult(cond string) bool {

	// TODO: refactor it, add new operators but not for now
	if strings.Contains(cond, "==") {
		cond = strings.Replace(cond, " ", "", -1)
		split := strings.Split(cond, "==")
		if len(split) != 2 {
			fmt.Printf("invalid condition: %s", cond)
			os.Exit(1)
		}

		return g.getMapValueByKey(split[0]) == split[1] || split[0] == g.getMapValueByKey(split[1])
	} else {
		fmt.Printf("unsupported condition %s", cond)
		os.Exit(1)
	}

	return false
}

func (g *Generator) getMapValueByKey(key string) interface{} {

	if key == "" {
		return nil
	}

	keys := strings.Split(key, ".")
	m := g.currentMap
	for i, k := range keys {
		a := m[k]
		if a == nil {
			return nil
		}

		if v, ok := a.(map[string]interface{}); ok {
			if i == len(keys)-1 {
				return v
			}
			m = v
		} else {
			if i == len(keys)-1 {
				return a
			} else {
				return nil
			}
		}
	}

	return m
}

func (g *Generator) getConfig(k interface{}) *AbsoluteConfig {

	if k == "" {
		return nil
	}

	for _, config := range g.Trainer.Information.AbsoluteConfig {
		if config.Key == k {
			return &config
		} else if isMatchesForArray(config.Key, k.(string)) {
			return &config
		} else if isWildCardMatches(config.Key, k.(string)) {
			return &config
		}
	}

	return nil
}

func (g *Generator) GenerateForFields(key string, m map[string]interface{}) map[string]interface{} {

	for k, v := range m {
		kk := getKey(key, k)
		switch v.(type) {
		case string:
			m[k] = g.generateString(kk, v.(string))
		case map[string]interface{}:
			m[k] = g.GenerateForFields(kk, v.(map[string]interface{}))
		case []interface{}:
			m[k] = v
		case int:
			m[k] = v
		case bool:
			m[k] = v
		default:
			m[k] = v
		}
	}

	return m
}

func (g *Generator) generateString(k, v string) string {

	if v == "" {
		return v
	}

	if !strings.Contains(v, "http://") && !strings.Contains(v, "https://") {
		text, ok := g.getText(v)
		if !ok {
			return v
		}
		return text
	}

	environmentUrl := g.getEnvironmentUrl(k, v)
	currentUrl := NewUrl(v)
	return fmt.Sprintf("%s://%s%s", environmentUrl.Scheme(), environmentUrl.Hostname(), currentUrl.Path())
}

func (g *Generator) getText(text string) (string, bool) {

	if text == "" {
		return "", false
	}

	field := g.findField(text)

	if field == nil {
		return "", false
	}

	if field.Type != "text" {
		return "", false
	}

	environment := field.Environment[g.environment]

	return environment.Value, true
}

func (g *Generator) getEnvironmentUrl(key, currentUrl string) Url {
	field := g.findField(currentUrl)

	if field != nil {
		environment := field.GetEnvironment(g.environment)
		if environment == nil {
			return NewUrl(currentUrl)
		}

		return NewUrl(environment)
	}

	fieldByKey := g.findFieldByKey(key)
	if fieldByKey == nil {
		return NewUrl(currentUrl)
	}

	environment := fieldByKey.GetEnvironment(g.environment)
	if environment == nil {
		return NewUrl(currentUrl)
	}

	return NewUrl(environment)

}

func (g *Generator) findField(val string) *Field {

	for _, field := range g.Trainer.Information.Fields {
		for _, environment := range field.Environment {
			if environment.Value != "" && strings.Contains(val, environment.Value) {
				return &field
			}
		}
	}

	return nil
}

func (g *Generator) findFieldByKey(key string) *Field {

	for _, field := range g.Trainer.Information.Fields {
		for _, k := range field.Keys {
			kLength := len(k)
			keyLength := len(key)
			if kLength > keyLength {
				if strings.Contains(k, key) {
					return &field
				}
			} else {
				if strings.Contains(key, k) {
					return &field
				}
			}
		}
	}

	return nil
}

func checkEnvironment(trainer *V1TrainerYaml, environment string) {
	for _, field := range trainer.Information.Fields {
		isDefined := false
		for k, _ := range field.Environment {
			if k == environment {
				isDefined = true
			}
		}

		if !isDefined {
			fmt.Printf("environment %s is not defined in %v\n", environment, field.Keys)
			os.Exit(1)
		}
	}
}
