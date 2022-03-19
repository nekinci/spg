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
			name: "Should return true",
			args: args{
				base:  "1.*.5",
				value: "1.2.3.4.5",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.2.3.4.6",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.*.5",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.2.3.5.4",
			},
			want: false,
		},
		{
			name: "Should return true",
			args: args{
				base:  "1.*.5",
				value: "1.5",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.5.3.4",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.5.5.4",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "2.3.4.5",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "5",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.2.3.4.5.6",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5",
				value: "1.5.6",
			},
			want: false,
		},
		{
			name: "Should return true",
			args: args{
				base:  "1.*.5.6.7.*.10",
				value: "1.2.3.4.5.6.7.8.9.10",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "1.*.5.6.7.*.10",
				value: "1.2.3.4.5.6.7.8.9",
			},
			want: false,
		},
		{
			name: "Should return true",
			args: args{
				base:  "1.*.5.6.7.*.10",
				value: "1.3.5.6.7.8.9.10",
			},
			want: true,
		},
		{
			name: "Should return true",
			args: args{
				base:  "*.10",
				value: "10",
			},
			want: true,
		},
		{
			name: "Should return true",
			args: args{
				base:  "*.10",
				value: "12.3.45.2221.2215.25.2444.555.677.54235.235235.35.253.352.52.52352.5235.235.35.235.235.235.235.235.3326236236.236236.236..10",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "*.10",
				value: "12.3.45.2221.2215.25.2444.555.677.54235.235235.35.253.352.52.52352.5235.235.35.235.235.235.235.235.3326236236.236236.236",
			},
			want: false,
		},
		{
			name: "Should return true",
			args: args{
				base:  "1.*.10",
				value: "1.2.3.4[5].6.7.8.9.10",
			},
			want: true,
		},
		{
			name: "Should return true",
			args: args{
				base:  "*.keycloak.niyazi",
				value: "keycloak.niyazi",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "*.keycloak.niyazi",
				value: "bulk-operations.keycloak.niyazi",
			},
			want: true,
		},
		{
			name: "Should return false",
			args: args{
				base:  "*.keycloak.niyazi",
				value: "keycloak.cors",
			},
			want: false,
		},
		{
			name: "Should return false",
			args: args{
				base:  "*.keycloak.niyazi.a.b.c.d",
				value: "keycloak",
			},
			want: false,
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
