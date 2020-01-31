package main

import (
	"log"
	"time"

	"github.com/JVillafruela/wdjqs/joconde"
	"github.com/JVillafruela/wdjqs/wd"
)

const (
	PInstanceOf    = "P31"
	PCreator       = "P170"
	PAdminLocation = "P131" // localisation administrative
	PLocation      = "P276"
	PInventory     = "P217"
	PCollection    = "P195"
	PMaterial      = "P186"
	PJocondeID     = "P347"
	PTitle         = "P1476"
	//used for sources
	PStatedIn     = "P248"
	PReferenceURL = "P854"
	PRetrieved    = "P813"
)

func main() {

	ref := "09940004427"

	a, err := joconde.GetArtwork(ref)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Printf("Found ref %s in joconde database : %s\n", ref, a.Title)

	item := wd.Item{Lang: "fr"}
	item.Label = joconde.GetMainTitle(a.Title)
	item.AddProperty(PTitle, item.Label)

	// TODO call to WDQS does not work reliably (2020-01-27) ?
	// in WDQS sometimes I got "Query timeout limit reached"
	var qid string
	name := joconde.ReverseName(a.Author)
	for i := 0; i < 3; i++ {
		log.Printf("Looking for author %s (try #%d)\n", name, i+1)
		qid, err = wd.FindAuthor(name)
		if err != nil {
			log.Println("Error : ", err)
		}
		if qid != "" {
			log.Printf("Found author : %s \n", qid)
			break
		}
		time.Sleep(10 * time.Second)
	}
	item.Description = a.Domain + " de " + name // "peinture de Georges de LA TOUR" ou Denomination "tableau" ?
	item.AddProperty(PCreator, qid)

	qid, err = wd.FindMuseumByMuseoID(a.Museo)
	log.Printf("Looking for museum %s QID=%s", a.Museo, qid)
	if err != nil {
		log.Println("Error : ", err)
	}
	if qid != "" {
		item.AddProperty(PCollection, qid)
		item.AddProperty(PLocation, qid)
	}

	log.Printf("Looking for inventory id %s", a.Inventory)
	qid, err = wd.FindArtworkByInventory(a.Inventory, qid)
	if err != nil {
		log.Println("Error : ", err)
	}
	if qid != "" {
		log.Fatalf("Error : artwork with same inventory number found %s\n", qid)
	}
	item.AddProperty(PInventory, a.Inventory)
	if item.Properties[PLocation].Value != "" {
		item.AddQualifier(PInventory, PCollection, item.Properties[PCollection].Value)
	}

	log.Printf("Looking for city '%s' ", a.City)
	qid, err = wd.FindCityByName(a.City)
	if err != nil {
		log.Println("Error : ", err)
	} else {
		log.Printf("Found city %s", qid)
		item.AddProperty(PAdminLocation, qid)
	}

	log.Printf("Looking for domain '%s' ", a.Domain)
	qid, err = wd.FindDomain(a.Domain)
	if err != nil {
		log.Println("FindDomain : ", err)
	}
	if qid != "" {
		log.Printf("Found domain %s", qid)
		item.AddProperty(PInstanceOf, qid)
	}

	log.Printf("Looking for materials %s", a.Materials)
	qid, err = wd.FindMaterial(a.Materials)
	if err != nil {
		log.Println("FindMaterial : ", err)
	}
	if qid != "" {
		log.Printf("Found material %s", qid)
		item.AddProperty(PMaterial, qid)
	}

	item.AddProperty(PJocondeID, a.Reference)

	addSources(&item, a.ReferenceURL)

	item.Qid = wd.QSandbox
	item.WriteQS("qs.txt")
	if err != nil {
		log.Println("WriteQS : ", err)
	}

	log.Println("Statements written in qs.txt")
}

func addSources(it *wd.Item, url string) {
	today := time.Now().Format("+2006-01-02T00:00:00Z/11") // YYYY-MM-DD
	sources := map[string]string{
		PStatedIn:     wd.QJoconde,
		PReferenceURL: url,
		PRetrieved:    today,
	}
	it.AddSourcesToAll(sources)
}
