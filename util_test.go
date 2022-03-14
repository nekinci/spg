package main

import "testing"

func Test_isMatchesForArray(t *testing.T) {
	type args struct {
		base  string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				base:  "a",
				value: "a",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				base:  "a",
				value: "b",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				base:  "a.b",
				value: "a.b",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				base:  "a.b",
				value: "a.c",
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				base:  "a.b[1]",
				value: "a.b[0]",
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				base:  "a.b[1]",
				value: "a.b[1]",
			},
			want: true,
		},
		{
			name: "test7",
			args: args{
				base:  "a.b[]",
				value: "a.b[2]",
			},
			want: true,
		},
		{
			name: "test8",
			args: args{
				base:  "a.b[1].c.d[].e",
				value: "a.b[1].c.d[3].e",
			},
			want: true,
		},
		{
			name: "test9",
			args: args{
				base:  "a.b[1].c.d[].e",
				value: "a.b[1].c.d[3].f",
			},
			want: false,
		},
		{
			name: "test10",
			args: args{
				base:  "a.b[1].c.d[].e.[4].a[].b",
				value: "a.b[1].c.d[3].e.[4].a[3].b",
			},
			want: true,
		},
		{
			name: "test11",
			args: args{
				base:  "a[0]",
				value: "a[0]",
			},
			want: true,
		},
		{
			name: "test12",
			args: args{
				base:  "a[0][1]",
				value: "a[0][1]",
			},
			want: true,
		},
		{
			name: "test13",
			args: args{
				base:  "a[3][]",
				value: "a[1][2]",
			},
			want: false,
		},
		{
			name: "test14",
			args: args{
				base:  "a[1][]",
				value: "a[1][2]",
			},
			want: true,
		},
		{
			name: "test15",
			args: args{
				base:  "a[][][]",
				value: "a[1][2][3]",
			},
			want: true,
		},
		{
			name: "test16",
			args: args{
				base:  "a.b[].c[].d[].e[2]",
				value: "a.b[1].c[2].d[3].e[2]",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchesForArray(tt.args.base, tt.args.value); got != tt.want {
				t.Errorf("isMatchesForArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWildCardMatches(t *testing.T) {
	type args struct {
		base  string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Only wildcard should return false",
			args: args{
				base:  "*",
				value: "abc",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				base:  "*",
				value: "*",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWildCardMatches(tt.args.base, tt.args.value); got != tt.want {
				t.Errorf("isWildCardMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
