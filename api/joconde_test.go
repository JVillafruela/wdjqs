package api

import (
	"reflect"
	"strings"
	"testing"
)

const js1 = `
{
	"total_count": 1,
	"links": [
		{
			"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait/records?start=0&rows=10&timezone=UTC&where=%2209940004427%22",
			"rel": "self"
		},
		{
			"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait/records?start=0&rows=10&timezone=UTC&where=%2209940004427%22",
			"rel": "first"
		},
		{
			"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait/records?start=0&rows=10&timezone=UTC&where=%2209940004427%22",
			"rel": "last"
		}
	],
	"records": [
		{
			"links": [
				{
					"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait/records/c3d52224de1363c373597cdf3a45e8069d3e6747",
					"rel": "self"
				},
				{
					"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait",
					"rel": "dataset"
				},
				{
					"href": "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets",
					"rel": "datasets"
				}
			],
			"record": {
				"id": "c3d52224de1363c373597cdf3a45e8069d3e6747",
				"timestamp": "2018-11-15T16:37:23.286Z",
				"size": 561,
				"fields": {
					"ville": "Grenoble",
					"paut": "Vic-sur-Seille, 1593 ; Lunéville, 1652",
					"milu": null,
					"lieux": null,
					"ddpt": null,
					"inv": "MG 408",
					"loca": [
						"Grenoble",
						"musée de Grenoble"
					],
					"mill": "1628 entre ; 1630 et",
					"decv": null,
					"epoq": null,
					"peri": [
						"2e quart 17e siècle"
					],
					"label": "Musée de France#au sens de la loi n°2002-5 du 4 janvier 2002",
					"geohi": null,
					"museo": "M0994",
					"peru": null,
					"titr": "Saint Jérôme pénitent#Dit aussi Saint Jérôme à l'auréole",
					"deno": [
						"tableau"
					],
					"ref": "09940004427",
					"stat": "Grenoble ; musée de Grenoble",
					"domn": [
						"peinture"
					],
					"dmaj": null,
					"dmis": "2007-09-06",
					"repr": null,
					"util": null,
					"dims": "H. 157, L. 100",
					"autr": "LA TOUR Georges de",
					"dacq": "1800",
					"srep": null,
					"ecol": "France",
					"adpt": null,
					"depo": null,
					"geolocalisation_ville": {
						"lat": 45.18627,
						"lon": 5.725358
					},
					"tech": [
						"peinture à l'huile (toile)"
					],
					"peoc": null,
					"appl": null,
					"onom": null
				}
			}
		}
	]
}
`

func TestJSONtoArtwork(t *testing.T) {
	js0 := strings.ReplaceAll(js1, "\"total_count\": 1", "\"total_count\": 0")
	js2 := strings.ReplaceAll(js1, "\"total_count\": 1", "\"total_count\": 2")

	type args struct {
		ref string
		js  string
	}
	tests := []struct {
		name    string
		args    args
		want    Artwork
		wantErr bool
	}{
		{
			"09940004427 (Saint Jérôme à l'auréole)",
			args{"09940004427", js1},
			Artwork{
				AcquisitionDate: "1800",
				Author:          "LA TOUR Georges de",
				City:            "Grenoble",
				Denomination:    "tableau",
				Dimensions:      "H. 157, L. 100",
				Domain:          "peinture",
				Inventory:       "MG 408",
				Materials:       "peinture à l'huile (toile)",
				Museo:           "M0994",
				Reference:       "09940004427",
				School:          "France",
				Title:           "Saint Jérôme pénitent#Dit aussi Saint Jérôme à l'auréole",
				Vintage:         "1628 entre ; 1630 et",
			},
			false,
		},
		{
			"99999999999 not found",
			args{"99999999999", js0},
			Artwork{},
			true,
		},
		{
			"99999999999 mutiple records",
			args{"99999999999", js2},
			Artwork{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONtoArtwork(tt.args.ref, tt.args.js)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONtoArtwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONtoArtwork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWordUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Simple name uppercase", args{"DELAUNAY"}, true},
		{"Hyphenated name uppercase", args{"CASIMIR-PERIER"}, true},
		{"Hyphenated name uppercase national", args{"CASIMIR-PÉRIER"}, true},
		{"Multiple name uppercase", args{"DE LA TOUR"}, true},
		{"Simple name mixed case", args{"Delaunay"}, false},
		{"Simple name lower case", args{"delaunay"}, false},
		{"Hyphenated name lowercase", args{"Casimir-Perier"}, false},
		{"Hyphenated name lowercase", args{"Casimir-Périer"}, false},
		{"Multiple name mixedcase", args{"De la TOUR"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWordUpper(tt.args.s); got != tt.want {
				t.Errorf("isWordUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Simple name uppercase", args{"DELAUNAY"}, "DELAUNAY"},
		{"lastname uppercase firstname mixed", args{"DELAUNAY Robert"}, "Robert DELAUNAY"},
		{"firstname mixed lastname uppercase", args{"Robert DELAUNAY"}, "Robert DELAUNAY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseName(tt.args.name); got != tt.want {
				t.Errorf("ReverseName() = %v, want %v", got, tt.want)
			}
		})
	}
}
