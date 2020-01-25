package wd

import "testing"

func TestFindDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"arg empty", args{""}, "", false},
		{"Domain OK", args{"peinture"}, "Q3305213", false},
		{"Domain unknown", args{"unknown domain"}, "", true},
		{"Domain w/o qid", args{"estampe;ethnologie"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindDomain(tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
