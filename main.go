package main

import (
	"fmt"
	"log"

	"github.com/JVillafruela/wdjqs/api"
)

func main() {
	ref := "09940004427"
	a, err := api.GetArtwork(ref)
	if err != nil {
		log.Fatal("Error : ", err)
	}
	fmt.Printf("%v \n", a)
}
