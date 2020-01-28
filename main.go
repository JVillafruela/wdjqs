package main

import (
	"log"
	"time"

	"github.com/JVillafruela/wdjqs/joconde"
	"github.com/JVillafruela/wdjqs/wd"
)

func main() {

	const (
		PInstanceOf    = "P31"
		PCreator       = "P170"
		PAdminLocation = "P131" // localisation administrative
		PLocation      = "P276"
		PInventory     = "P217"
		PMaterial      = "P186"
		PJocondeID     = "P347"
	)

	ref := "09940004427"

	a, err := joconde.GetArtwork(ref)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Printf("Found ref %s in joconde database : %s\n", ref, a.Title)

	item := wd.Item{Lang: "fr"}
	item.Label = joconde.GetMainTitle(a.Title)

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
	item.Add(PCreator, qid)

	qid, err = wd.FindMuseumByMuseoID(a.Museo)
	log.Printf("Looking for museum %s QID=%s", a.Museo, qid)
	if err != nil {
		log.Println("Error : ", err)
	}
	if qid != "" {
		item.Add(PLocation, qid)
	}

	log.Printf("Looking for inventory id %s", a.Inventory)
	qid, err = wd.FindArtworkByInventory(a.Inventory, qid)
	if err != nil {
		log.Println("Error : ", err)
	}
	if qid != "" {
		log.Fatalf("Error : artwork with same inventory found %s\n", qid)
	}
	item.Add(PInventory, a.Inventory)

	log.Printf("Looking for city '%s' ", a.City)
	qid, err = wd.FindCityByName(a.City)
	if err != nil {
		log.Println("Error : ", err)
	} else {
		log.Printf("Found city %s", qid)
		item.Add(PAdminLocation, qid)
	}

	log.Printf("Looking for domain '%s' ", a.Domain)
	qid, err = wd.FindDomain(a.Domain)
	if err != nil {
		log.Println("FindDomain : ", err)
	}
	if qid != "" {
		log.Printf("Found domain %s", qid)
		item.Add(PInstanceOf, qid)
	}

	log.Printf("Looking for materials %s", a.Materials)
	qid, err = wd.FindMaterial(a.Materials)
	if err != nil {
		log.Println("FindMaterial : ", err)
	}
	if qid != "" {
		log.Printf("Found material %s", qid)
		item.Add(PMaterial, qid)
	}

	item.Add(PJocondeID, a.Reference)

	log.Printf("Item : %v\n", item)
}
