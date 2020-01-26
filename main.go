package main

import (
	"log"

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

	log.Printf("Looking for author %s\n", joconde.ReverseName(a.Author))
	qid, err := wd.FindAuthor(joconde.ReverseName(a.Author))
	if err != nil {
		log.Fatal("Error : ", err)
	}
	if qid != "" {
		log.Printf("Found author : %s \n", qid)
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
