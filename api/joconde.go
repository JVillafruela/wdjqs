package api

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// Artwork data read from Joconde API
type Artwork struct {
	AcquisitionDate string
	Author          string
	City            string
	Denomination    string
	Dimensions      string
	Domain          string
	Inventory       string
	Materials       string
	Museo           string
	Reference       string
	School          string
	Title           string
	Vintage         string
}

func callArtworkAPI(ref string) (string, error) {
	url := "https://data.culturecommunication.gouv.fr/api/v2/catalog/datasets/base-joconde-extrait/records?where=%22" + ref + "%22&rows=10&pretty=false"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	dataInBytes, err := ioutil.ReadAll(res.Body)
	return string(dataInBytes), err
}

func JSONtoArtwork(ref string, js string) (Artwork, error) {
	count := gjson.Get(js, "total_count").Int()
	if count == 0 {
		return Artwork{}, errors.New("No record found for reference " + ref)
	}
	if count > 1 {
		return Artwork{}, errors.New("Multiple records found for reference " + ref)
	}

	a := Artwork{}
	a.AcquisitionDate = gjson.Get(js, "records.0.record.fields.dacq").String()
	a.Author = gjson.Get(js, "records.0.record.fields.autr").String()
	a.City = gjson.Get(js, "records.0.record.fields.ville").String()
	a.Denomination = gjson.Get(js, "records.0.record.fields.deno.0").String()
	a.Dimensions = gjson.Get(js, "records.0.record.fields.dims").String()
	a.Domain = gjson.Get(js, "records.0.record.fields.domn.0").String()
	a.Inventory = gjson.Get(js, "records.0.record.fields.inv").String()
	a.Materials = gjson.Get(js, "records.0.record.fields.tech.0").String()
	a.Museo = gjson.Get(js, "records.0.record.fields.museo").String()
	a.Reference = gjson.Get(js, "records.0.record.fields.ref").String()
	a.School = gjson.Get(js, "records.0.record.fields.ecol").String()
	a.Title = gjson.Get(js, "records.0.record.fields.titr").String()
	a.Vintage = gjson.Get(js, "records.0.record.fields.mill").String()

	return a, nil
}

//GetArtwork : Get artwork data for reference ref eg. 09940004427
func GetArtwork(ref string) (Artwork, error) {
	js, err := callArtworkAPI(ref)
	if err != nil {
		return Artwork{}, err
	}
	return JSONtoArtwork(ref, js)
}
