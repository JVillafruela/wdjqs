package main

import (
	"log"
	"time"

	"github.com/JVillafruela/wdjqs/joconde"
	"github.com/JVillafruela/wdjqs/wd"
)

func main() {

	ref := "09940004427"

	a, err := joconde.GetArtwork(ref)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	log.Printf("Found ref %s in joconde database : %s\n", ref, a.Title)

	// TODO call to WDQS does not work reliably ?
	name := joconde.ReverseName(a.Author)
	for i := 0; i < 3; i++ {
		log.Printf("Looking for author %s (try #%d)\n", name, i+1)
		qid, err := wd.FindAuthor(name)
		if err != nil {
			log.Println("Error : ", err)
		}
		if qid != "" {
			log.Printf("Found author : %s \n", qid)
			break
		}
		time.Sleep(10 * time.Second)
	}

	/*
		qid, err := wd.FindMuseumByMuseoID("M0994")
		if err != nil {
			log.Fatal("Error : ", err)
		}
		fmt.Printf("%v \n", qid)
	*/
	/*
		qid, err := wd.FindArtworkByInventory("MG 2998", "")
		if err != nil {
			log.Fatal("Error : ", err)
		}
		fmt.Printf("%v \n", qid)
	*/
	/*
		qid, err := wd.FindCityByName("Vienne")
		if err != nil {
			log.Fatal("Error : ", err)
		}
		fmt.Printf("%v \n", qid)
	*/
	/*
		qid, err := wd.FindDomain("zz")
		if err != nil {
			log.Println("FindDomain : ", err)
		}
		fmt.Printf("%v \n", qid)
	*/
}
