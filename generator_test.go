package main

import "testing"

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
