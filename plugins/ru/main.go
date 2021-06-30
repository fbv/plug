package main

import (
	"log"
	"plug/module"
)

var Entry module.Entry

func init() {
	log.Println("RU Module Initialization")
	Entry = module.Entry{
		Lang: "RU",
	}
}
