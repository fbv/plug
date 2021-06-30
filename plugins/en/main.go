package main

import (
	"log"
	"plug/module"
)

var Entry module.Entry

func init() {
	log.Println("EN Module Initialization")
	Entry = module.Entry{
		Lang: "EN",
	}
}
