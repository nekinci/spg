package main

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	type args struct {
		maps []*map[string]interface{}
	}

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Should use as is for int",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 5,
						"b": 6,
					},
				},
			},
			want: map[string]interface{}{
				"a": 5,
				"b": 6,
			},
		},
		{
			name: "Should use as is for string",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": "5",
						"b": "6",
					},
				},
			},
			want: map[string]interface{}{
				"a": "5",
				"b": "6",
			},
		},
		{
			name: "Should use as is for bool",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": true,
						"b": false,
					},
				},
			},
			want: map[string]interface{}{
				"a": true,
				"b": false,
			},
		},
		{
			name: "Should use as is for float",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 5.5,
						"b": 6.6,
					},
				},
			},
			want: map[string]interface{}{
				"a": 5.5,
				"b": 6.6,
			},
		},
		{
			name: "Should use as is for array",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": []int{1, 2, 3},
						"b": []int{4, 5, 6},
					},
				},
			},
			want: map[string]interface{}{
				"a": []int{1, 2, 3},
				"b": []int{4, 5, 6},
			},
		},
		{
			name: "Should use as is for map",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": map[string]interface{}{
							"b": 1,
							"c": 2,
						},
						"b": map[string]interface{}{
							"b": 3,
							"c": 4,
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": 1,
					"c": 2,
				},
				"b": map[string]interface{}{
					"b": 3,
					"c": 4,
				},
			},
		},
		{
			name: "Should second override first I",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 5,
						"b": 6,
					},
					&map[string]interface{}{
						"a": 5.5,
						"b": 6.6,
					},
				},
			},
			want: map[string]interface{}{
				"a": 5.5,
				"b": 6.6,
			},
		},
		{
			name: "Should second override first II",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 5.5,
						"b": 6.6,
					},
					&map[string]interface{}{
						"a": 5,
						"b": 6,
					},
				},
			},
			want: map[string]interface{}{
				"a": 5,
				"b": 6,
			},
		},
		{
			name: "Should second override first III",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 1,
						"b": 3.5,
						"c": []int{1, 2, 3},
					},
					&map[string]interface{}{
						"a": 2,
						"b": 3.5,
					},
				},
			},
			want: map[string]interface{}{
				"a": 2,
				"b": 3.5,
				"c": []int{1, 2, 3},
			},
		},
		{
			name: "Should second override first IV",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 2,
						"b": 3,
						"c": "5",
						"d": []int{1, 2, 3},
					},
					&map[string]interface{}{
						"d": []int{2, 3, 4},
					},
				},
			},
			want: map[string]interface{}{
				"a": 2,
				"b": 3,
				"c": "5",
				"d": []interface{}{2, 3, 4},
			},
		},
		{
			name: "Should second override first V",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": []int{1, 2, 3, 4, 5},
					},
					&map[string]interface{}{
						"a": []int{2, 2, 3, 4},
					},
				},
			},
			want: map[string]interface{}{
				"a": []interface{}{2, 2, 3, 4},
			},
		},
		{
			name: "Should second override first VI",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": []interface{}{
							1, 2, 3,
							map[string]interface{}{
								"a": "1",
								"b": "2",
								"c": []interface{}{1, 2, 3, 4, 5},
							},
						},
					},
					&map[string]interface{}{
						"a": []interface{}{
							2, 2, 3, map[string]interface{}{
								"a": 2,
								"c": []interface{}{3},
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": []interface{}{
					2, 2, 3, map[string]interface{}{
						"a": 2,
						"c": []interface{}{3},
					},
				},
			},
		},
		{
			name: "Should second override first VII",
			args: args{
				maps: []*map[string]interface{}{
					{
						"a": []interface{}{1, 2, 3, map[string]interface{}{"a": 2, "b": map[string]interface{}{"c": 3}}},
					},
					{
						"a": []interface{}{4, 5, 6, map[string]interface{}{"a": 1}},
					},
				},
			},
			want: map[string]interface{}{
				"a": []interface{}{4, 5, 6, map[string]interface{}{"a": 1}},
			},
		},
		{
			name: "Should second override first VIII",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 5,
						"b": []interface{}{
							map[string]interface{}{
								"a": map[string]interface{}{
									"a": 5,
									"b": 6,
								},
							},
						},
					},
					&map[string]interface{}{
						"b": []interface{}{
							map[string]interface{}{
								"a": map[string]interface{}{
									"b": "c",
								},
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": 5,
				"b": []interface{}{
					map[string]interface{}{
						"a": map[string]interface{}{
							"b": "c",
						},
					},
				},
			},
		},
		{
			name: "Should second override first IX",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": map[string]interface{}{"c": "d"},
					},
					&map[string]interface{}{
						"a": []interface{}{1, 2, 3},
					},
				},
			},
			want: map[string]interface{}{
				"a": []interface{}{1, 2, 3},
			},
		},
		{
			name: "Should third override first and two",
			args: args{
				maps: []*map[string]interface{}{
					&map[string]interface{}{
						"a": 2,
						"b": "3",
						"c": []interface{}{1, 2, "3"},
						"d": map[string]interface{}{
							"a": 4,
						},
					},
					&map[string]interface{}{
						"e": 55,
						"c": []interface{}{4},
						"d": map[string]interface{}{
							"c": 5,
							"d": 6,
						},
					},
					&map[string]interface{}{
						"f": "g",
						"d": map[string]interface{}{
							"a": 7,
							"e": 7,
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": 2,
				"b": "3",
				"c": []interface{}{4},
				"d": map[string]interface{}{
					"a": 7,
					"c": 5,
					"d": 6,
					"e": 7,
				},
				"f": "g",
				"e": 55,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrettyPrint(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should print",
			args: args{
				map[string]interface{}{
					"a": 3,
					"b": 4,
					"c": map[string]interface{}{
						"d": "e",
					},
				},
			},
			want: string("a: 3\nb: 4\nc:\n    d: e\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyPrint(tt.args.m); got != tt.want {
				t.Errorf("PrettyPrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
