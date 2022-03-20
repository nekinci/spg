package main

import (
	"reflect"
	"testing"
)

func TestUrl_Hostname(t *testing.T) {
	type fields struct {
		Url string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test 1",
			fields: fields{
				Url: "https://www.google.com",
			},
			want: "www.google.com",
		},
		{
			name: "Test 2",
			fields: fields{
				Url: "https://www.google.com/",
			},
			want: "www.google.com",
		},
		{
			name: "Test 3",
			fields: fields{
				Url: "https://www.google.com/search?q=test",
			},
			want: "www.google.com",
		},
		{
			name: "Test 4",
			fields: fields{
				Url: "https://www.google.com/search?q=test&oq=test",
			},
			want: "www.google.com",
		},
		{
			name: "Test 5",
			fields: fields{
				Url: "https://www.google.com/search?q=test&oq=test&aqs=chrome..69i57j0l5.1209j0j7&sourceid=chrome&ie=UTF-8",
			},
			want: "www.google.com",
		},
		{
			name: "Test 6",
			fields: fields{
				Url: "https://linkedin.com/in/joseph-m-kim-a-b8b9b9b9",
			},
			want: "linkedin.com",
		},
		{
			name: "With http",
			fields: fields{
				Url: "http://www.google.com",
			},
			want: "www.google.com",
		},
		{
			name: "With http and port",
			fields: fields{
				Url: "http://www.google.com:80",
			},
			want: "www.google.com:80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Url{
				Url: tt.fields.Url,
			}
			if got := u.Hostname(); got != tt.want {
				t.Errorf("Hostname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUrl_Path(t *testing.T) {
	type fields struct {
		Url string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test 1",
			fields: fields{
				Url: "https://www.google.com",
			},
			want: "",
		},
		{
			name: "Test 2",
			fields: fields{
				Url: "https://www.google.com/",
			},
			want: "/",
		},
		{
			name: "Test 3",
			fields: fields{
				Url: "https://www.google.com/search?q=test",
			},
			want: "/search?q=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := Url{
				Url: tt.fields.Url,
			}
			if got := u.Path(); got != tt.want {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_generateString(t *testing.T) {

	trainer := V1TrainerYaml{
		Version: "v1",
		Information: Information{
			Fields: []Field{
				{
					Keys: []string{"test-service", "test-srv", "service-test", "testService", "testServiceUrl"},
					Type: "url",
					Environment: map[string]Environment{
						"test": {
							Value:  "test-service-test.cloud.com",
							Scheme: "http",
						},
						"oc": {
							Value:  "test-service-cloud:8080",
							Scheme: "http",
						},
						"prod": {
							Value:  "test-service.com",
							Scheme: "https",
						},
						"preprod": {
							Value:  "test-service-prp.cloud.com",
							Scheme: "https",
						},
					},
				},
			},
		},
	}
	type fields struct {
		Trainer     *V1TrainerYaml
		environment string
	}
	type args struct {
		k string
		v string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "",
			fields: fields{
				Trainer:     &trainer,
				environment: "test",
			},
			args: args{
				k: "test-service",
				v: "http://test-service-cloud:8080",
			},
			want: "http://test-service-test.cloud.com",
		},
		{
			name: "",
			fields: fields{
				Trainer:     &trainer,
				environment: "test",
			},
			args: args{
				k: "test-service",
				v: "http://test-service-cloud:8080/rest/partyManagement",
			},
			want: "http://test-service-test.cloud.com/rest/partyManagement",
		},
		{
			name: "",
			fields: fields{
				Trainer:     &trainer,
				environment: "test",
			},
			args: args{
				k: "test-service",
				v: "http://test-service-test.cloud.com/rest/partyManagement",
			},
			want: "http://test-service-test.cloud.com/rest/partyManagement",
		},
		{
			name: "",
			fields: fields{
				Trainer:     &trainer,
				environment: "prod",
			},
			args: args{
				k: "test-service",
				v: "http://test-service-test.cloud.com/rest/partyManagement",
			},
			want: "https://test-service.com/rest/partyManagement",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				Trainer:     tt.fields.Trainer,
				environment: tt.fields.environment,
			}
			if got := g.generateString(tt.args.k, tt.args.v); got != tt.want {
				t.Errorf("generateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_GenerateForAbsoluteConfig(t *testing.T) {
	cond := "a.b.d == 10"
	cond1 := "a.b.d == 5"
	type fields struct {
		Trainer     *V1TrainerYaml
		environment string
		currentMap  map[string]interface{}
	}
	type args struct {
		key string
		m   map[string]interface{}
	}

	tr := V1TrainerYaml{
		Version: "v1",
		Information: Information{
			Fields: []Field{},
			AbsoluteConfig: []AbsoluteConfig{
				{
					Key: "crm-data-service.environment",
					Environment: map[string]interface{}{
						"oc":   "DEV",
						"test": "TEST",
						"prp":  "CLOUD_PRP",
						"prod": "CLOUD_PROD",
					},
				},
				{
					Key: "keycloak.security[0].authRoles",
					Environment: map[string]interface{}{
						"oc":   "A",
						"test": "B",
						"prp":  "C",
						"prod": "D",
					},
				},
				{
					Key: "a[0][1]",
					Environment: map[string]interface{}{
						"oc":   "B",
						"test": "B",
						"prp":  "B",
						"prod": "B",
					},
				},
				{
					Key: "a[1][]",
					Environment: map[string]interface{}{
						"oc":   "B",
						"test": "B",
						"prp":  "B",
						"prod": "B",
					},
				},
				{
					Key: "b[][]",
					Environment: map[string]interface{}{
						"oc":   "1",
						"test": "1",
						"prp":  "1",
						"prod": "1",
					},
				},
				{
					Key: "*.bb.cc",
					Environment: map[string]interface{}{
						"oc":   "<new-value>",
						"test": "<new-value>",
						"prp":  "<new-value>",
						"prod": "<new-value>",
					},
				},
				{
					Key: "1.*.5",
					Environment: map[string]interface{}{
						"oc":   "zzz",
						"test": "zzz",
						"prp":  "zzz",
						"prod": "zzz",
					},
				},
			},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "Should Change Absoulute Value",
			fields: fields{
				Trainer:     &tr,
				environment: "prp",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"crm-data-service": map[string]interface{}{
						"environment": "DEV",
					},
				},
			},
			want: map[string]interface{}{
				"crm-data-service": map[string]interface{}{
					"environment": "CLOUD_PRP",
				},
			},
		},
		{
			name: "Should change absolute values II",
			fields: fields{
				Trainer:     &tr,
				environment: "prp",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"keycloak": map[string]interface{}{
						"security": []interface{}{
							map[string]interface{}{
								"authRoles": "Z",
							},
							map[string]interface{}{
								"securityCollections": "C",
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"keycloak": map[string]interface{}{
					"security": []interface{}{
						map[string]interface{}{
							"authRoles": "C",
						},
						map[string]interface{}{
							"securityCollections": "C",
						},
					},
				},
			},
		},
		{
			name: "Should not change when 'if condition' is not true",
			fields: fields{
				Trainer: &V1TrainerYaml{
					Version: "v1",
					Information: Information{
						Fields: nil,
						AbsoluteConfig: []AbsoluteConfig{
							{
								Key: "a.b.c",
								Environment: map[string]interface{}{
									"oc": "1", "test": "2", "prp": "3", "prod": "4",
								},
								Condition: &cond,
							},
						},
					},
				},
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
			},
		},
		{
			name: "Should change when 'if condition' is true",
			fields: fields{
				Trainer: &V1TrainerYaml{
					Version: "v1",
					Information: Information{
						Fields: nil,
						AbsoluteConfig: []AbsoluteConfig{
							{
								Key: "a.b.c",
								Environment: map[string]interface{}{
									"oc": "1", "test": "2", "prp": "3", "prod": "4",
								},
								Condition: &cond1,
							},
						},
					},
				},
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{"b": map[string]interface{}{"c": "2", "d": "5"}},
			},
		},
		{
			name: "Should change when 'if condition' is true with wildcard",
			fields: fields{
				Trainer: &V1TrainerYaml{
					Version: "v1",
					Information: Information{
						Fields: nil,
						AbsoluteConfig: []AbsoluteConfig{
							{
								Key: "*.b.c",
								Environment: map[string]interface{}{
									"oc": "1", "test": "2", "prp": "3", "prod": "4",
								},
								Condition: &cond1,
							},
						},
					},
				},
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"a": map[string]interface{}{"b": map[string]interface{}{"c": 3, "d": "5"}},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{"b": map[string]interface{}{"c": "2", "d": "5"}},
			},
		},
		{
			name: "Should change absolute array values",
			fields: fields{
				Trainer:     &tr,
				environment: "prp",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"a": []interface{}{
						[]interface{}{
							"A",
							"b",
						},
						[]interface{}{
							"B",
							"C",
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": []interface{}{
					[]interface{}{
						"A",
						"B",
					},
					[]interface{}{
						"B",
						"B",
					},
				},
			},
		},
		{
			name: "All array fields should be changed",
			fields: fields{
				Trainer:     &tr,
				environment: "prp",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"b": []interface{}{
						[]interface{}{
							"A",
							"A",
						},
						[]interface{}{
							"A",
							"A",
						},
					},
				},
			},
			want: map[string]interface{}{
				"b": []interface{}{
					[]interface{}{
						"1",
						"1",
					},
					[]interface{}{
						"1",
						"1",
					},
				},
			},
		},
		{
			name: "Should change * values",
			fields: fields{
				Trainer:     &tr,
				environment: "test",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"aa": map[string]interface{}{
						"bb": map[string]interface{}{
							"cc": "1",
						},
					},
					"dd": map[string]interface{}{
						"bb": map[string]interface{}{
							"cc": "2",
						},
					},
				},
			},
			want: map[string]interface{}{
				"aa": map[string]interface{}{
					"bb": map[string]interface{}{
						"cc": "<new-value>",
					},
				},
				"dd": map[string]interface{}{
					"bb": map[string]interface{}{
						"cc": "<new-value>",
					},
				},
			},
		},
		{
			name: "Should change * values",
			fields: fields{
				Trainer:     &tr,
				environment: "test",
			},
			args: args{
				key: "",
				m: map[string]interface{}{
					"1": map[string]interface{}{
						"2": map[string]interface{}{
							"3": map[string]interface{}{
								"4": map[string]interface{}{
									"5": 123,
								},
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"1": map[string]interface{}{
					"2": map[string]interface{}{
						"3": map[string]interface{}{
							"4": map[string]interface{}{
								"5": "zzz",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				Trainer:     tt.fields.Trainer,
				environment: tt.fields.environment,
				currentMap:  tt.fields.currentMap,
			}
			if got := g.GenerateForAbsoluteConfig(tt.args.key, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateForAbsoluteConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_getConditionResult(t *testing.T) {
	cond := "a.c == 3"
	var tr = &V1TrainerYaml{
		Version: "v1",
		Information: Information{
			Fields: nil,
			AbsoluteConfig: []AbsoluteConfig{
				{
					Key: "a.b",
					Environment: map[string]interface{}{
						"oc":   "100",
						"test": "100",
						"prp":  "200",
						"prod": "200",
					},
					Condition: &cond,
				},
			},
		},
	}
	type fields struct {
		Trainer     *V1TrainerYaml
		environment string
		currentMap  map[string]interface{}
	}
	type args struct {
		cond string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should return true",
			fields: fields{
				Trainer:     tr,
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{
						"b": "2",
						"c": "3",
					},
				},
			},
			args: args{
				cond: cond,
			},
			want: true,
		},
		{
			name: "Should return false",
			fields: fields{
				Trainer:     tr,
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{
						"b": "22",
						"c": "11",
					},
				},
			},
			args: args{
				cond: "a == 2",
			},
			want: false,
		},
		{
			name: "Should return false",
			fields: fields{
				Trainer:     tr,
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{
						"b": "22",
						"c": "11",
					},
				},
			},
			args: args{
				cond: "a.b == abc",
			},
			want: false,
		},
		{
			name: "Should return true",
			fields: fields{
				Trainer:     tr,
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{
						"b": "22",
						"c": "11",
					},
				},
			},
			args: args{
				cond: "a.c == 11",
			},
			want: true,
		},
		{
			name: "Should return true",
			fields: fields{
				Trainer:     tr,
				environment: "test",
				currentMap: map[string]interface{}{
					"a": map[string]interface{}{
						"b": "22",
						"c": "11",
					},
				},
			},
			args: args{
				cond: "a.c == a.c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				Trainer:     tt.fields.Trainer,
				environment: tt.fields.environment,
				currentMap:  tt.fields.currentMap,
			}

			if got := g.getConditionResult(tt.args.cond); got != tt.want {
				t.Errorf("getConditionResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
