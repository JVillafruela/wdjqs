package wd

import (
	"reflect"
	"testing"
)

func Test_getQid(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Qid OK", args{"http://www.wikidata.org/entity/Q203371"}, "Q203371"},
		{"Qid empty", args{"http://www.wikidata.org/entity/"}, ""},
		{"Empty uri", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQid(tt.args.uri); got != tt.want {
				t.Errorf("getQid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAuthorDecodeJSON(t *testing.T) {
	const js0 = `
	{
		"head": {
			"vars": [
				"item",
				"num"
			]
		},
		"results": {
			"bindings": []
		}
	}`
	const js1 = `
	{
		"head": {
			"vars": [
				"item",
				"num"
			]
		},
		"results": {
			"bindings": [
				{
					"item": {
						"type": "uri",
						"value": "http://www.wikidata.org/entity/Q203371"
					},
					"num": {
						"datatype": "http://www.w3.org/2001/XMLSchema#int",
						"type": "literal",
						"value": "0"
					}
				}
			]
		}
	} `

	const js2 = `
	{
		"head": {
			"vars": [
				"item",
				"num"
			]
		},
		"results": {
			"bindings": [
				{
					"item": {
						"type": "uri",
						"value": "http://www.wikidata.org/entity/Q1"
					},
					"num": {
						"datatype": "http://www.w3.org/2001/XMLSchema#int",
						"type": "literal",
						"value": "0"
					}
				},
				{
					"item": {
						"type": "uri",
						"value": "http://www.wikidata.org/entity/Q2"
					},
					"num": {
						"datatype": "http://www.w3.org/2001/XMLSchema#int",
						"type": "literal",
						"value": "1"
					}
				}
			]
		}
	} `

	const jsKO = `
	{
		"head": {
			"vars": [
				"item",
				"num"
	`

	type args struct {
		js string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Malformed json", args{jsKO}, []string{}, true},
		{"Author not found", args{js0}, []string{}, false},
		{"Author found", args{js1}, []string{"Q203371"}, false},
		{"Multiple authors found", args{js2}, []string{"Q1", "Q2"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findAuthorDecodeJSON(tt.args.js)
			if (err != nil) != tt.wantErr {
				t.Errorf("findAuthorDecodeJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAuthorDecodeJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
