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
		{"empty", args{""}, false},
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

func Test_isDate(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"date OK", args{"+2020-02-04T00:00:00Z/11"}, true},
		{"date KO", args{"+ 2000"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDate(tt.args.value); got != tt.want {
				t.Errorf("isDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDimension(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"w/o unit", args{"123"}, false},
		{"w/o U code", args{"123 cm"}, false},
		{"U code w/o value", args{"U174728"}, false},
		{"dim in cm", args{"123U174728"}, true},
		{"zero cm", args{"0U174728"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDimension(tt.args.value); got != tt.want {
				t.Errorf("isDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isQid(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"value", args{"Quorum"}, false},
		{"Qid", args{"Q42"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isQid(tt.args.value); got != tt.want {
				t.Errorf("isQid() = %v, want %v", got, tt.want)
			}
		})
	}
}
