package main

import (
	"log"
	"plug/module"
)

func main() {
	log.Println("Main")
	err := module.Load()
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range module.Entries {
		log.Println(m.Lang)
	}
	log.Println("Done")
}
