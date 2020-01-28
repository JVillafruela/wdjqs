package wd

import (
	"testing"
)

func Test_quote(t *testing.T) {
	type args struct {
		str  string
		lang string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{"", ""}, "\"\""},
		{"empty+lang", args{"", "fr"}, "\"\""},
		{"no lang", args{"MG001", ""}, "\"MG001\""},
		{"lang fr", args{"MG001", "fr"}, "fr:\"MG001\""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quote(tt.args.str, tt.args.lang); got != tt.want {
				t.Errorf("quote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPropertyLang(t *testing.T) {
	type args struct {
		prop string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"emptty", args{""}, false},
		{"needed", args{"P1476"}, true},
		{"not needed", args{"P217"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPropertyLang(tt.args.prop); got != tt.want {
				t.Errorf("IsPropertyLang() = %v, want %v", got, tt.want)
			}
		})
	}
}
