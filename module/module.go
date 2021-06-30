package module

import (
	"io/fs"
	"log"
	"path/filepath"
	"plugin"
	"strings"
)

type Entry struct {
	Lang string
}

var Entries = make([]*Entry, 0)

func Load() error {
	if len(Entries) > 0 {
		log.Println("Modules already loaded")
		return nil
	}
	err := filepath.Walk("./lib", loadFile)
	if err != nil {
		return err
	}
	log.Println("Modules loaded")
	return nil
}

func loadFile(path string, info fs.FileInfo, err error) error {
	if err != nil {
		log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if !info.IsDir() && strings.HasSuffix(info.Name(), ".so") {
		log.Println("Loading", path)
		p, err := plugin.Open(path)
		if err != nil {
			return err
		}
		s, err := p.Lookup("Entry")
		if err != nil {
			return err
		}
		e, ok := s.(*Entry)
		if ok {
			Entries = append(Entries, e)
		} else {
			log.Println("Unable to cast Entry")
		}
	}
	return nil
}
