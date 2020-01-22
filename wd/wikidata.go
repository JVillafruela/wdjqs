package wd

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/JVillafruela/wdjqs/api"
	"github.com/tidwall/gjson"
)

// FindAuthor : Find Author entity by name
func FindAuthor(name string) (string, error) {

	const endpoint = "https://query.wikidata.org/bigdata/namespace/wdq/sparql?query=%s&format=json"

	const sparql = `
	SELECT * WHERE {
		SERVICE wikibase:mwapi {
			bd:serviceParam wikibase:api "EntitySearch" .
			bd:serviceParam wikibase:endpoint "www.wikidata.org" .
			bd:serviceParam mwapi:search "%s" .
			bd:serviceParam mwapi:language "fr" .
			?item wikibase:apiOutputItem mwapi:item .
			?num wikibase:apiOrdinal true .
		}
		?item (wdt:P279|wdt:P31) wd:Q5
	} ORDER BY ASC(?num) LIMIT 5 `

	query := fmt.Sprintf(sparql, name)
	url := fmt.Sprintf(endpoint, url.PathEscape(query))
	js, err := api.CallAPI(url)
	if err != nil {
		return "", err
	}
	qids, err := getQidsFromJSON(js)
	if err != nil {
		return "", err
	}
	if len(qids) == 0 {
		log.Printf("Author name '%s' not found \n", name)
		return "", nil
	}
	if len(qids) > 1 {
		log.Printf("Multiple values found for author name '%s' : %s \n", name, strings.Join(qids, ","))
	}

	return qids[0], nil
}

// get QIDs from json query result
// variable for item must be named "item" : SELECT ?item,...
func getQidsFromJSON(js string) ([]string, error) {
	if !gjson.Valid(js) {
		return []string{}, errors.New("getQidsFromJSON : Invalid json")
	}
	/*  no result :
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
	}

	1 result "Georges de La Tour" (Q203371)
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
	}

	*/
	qids := []string{}
	if !gjson.Get(js, "results.bindings.0.item.value").Exists() {
		return qids, nil
	}

	values := gjson.Get(js, "results.bindings.#.item.value")
	for _, value := range values.Array() {
		uri := value.String()
		qids = append(qids, getQid(uri))
	}

	return qids, nil
}

// http://www.wikidata.org/entity/Q203371 => Q203371
func getQid(uri string) string {
	i := strings.LastIndex(uri, "/") + 1
	qid := uri[i:]
	return qid
}

// FindMuseumByMuseoID : lookup for museum item by museo number
func FindMuseumByMuseoID(museo string) (string, error) {
	const endpoint = "https://query.wikidata.org/sparql?query=%s&format=json"
	query := fmt.Sprintf(`SELECT ?item WHERE {?item wdt:P539 "%s"}`, museo)
	url := fmt.Sprintf(endpoint, url.PathEscape(query))
	js, err := api.CallAPI(url)
	if err != nil {
		return "", err
	}

	qids, err := getQidsFromJSON(js)
	if err != nil {
		return "", err
	}
	if len(qids) == 0 {
		log.Printf("Museum not found for museo id '%s'  \n", museo)
		return "", nil
	}
	if len(qids) > 1 {
		log.Printf("Multiple values found for museo id '%s' : %s \n", museo, strings.Join(qids, ","))
	}

	return qids[0], nil
}

func FindArtworkByInventory(inv string, museum string) (string, error) {
	sparql := `
		SELECT ?artwork WHERE {
			?artwork wdt:P217 "MG 2998"
		}          
		
	`
	// +++ si 0 ou 1 ok si >1 faire condition sur musée
	return sparql, nil
}
