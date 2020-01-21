package main

import (
	"fmt"
	"log"

	"github.com/JVillafruela/wdjqs/joconde"
)

func main() {
	ref := "09940004427"
	a, err := joconde.GetArtwork(ref)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	fmt.Printf("%v \n", a)
}
