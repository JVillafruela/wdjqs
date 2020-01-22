package main

import (
	"fmt"
	"log"

	"github.com/JVillafruela/wdjqs/wd"
)

func main() {
	/*
		ref := "09940004427"
		a, err := joconde.GetArtwork(ref)
		if err != nil {
			log.Fatal("Error : ", err)
		}
		fmt.Printf("%v \n", a)
	*/

	/*
		js, err := wd.FindAuthor("Georges de La Tour")
		if err != nil {
			log.Fatal("Error : ", err)
		}
		fmt.Printf("%v \n", js)
	*/

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
	qid, err := wd.FindCityByName("Vienne")
	if err != nil {
		log.Fatal("Error : ", err)
	}
	fmt.Printf("%v \n", qid)

}
