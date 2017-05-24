package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"sabey.co/ishtargate"
)

var (
	settings_input   = flag.String("settings-input", "", "Raw JSON Settings")
	settings_file    = flag.String("settings-file", "ishtargate-settings.json", "File of JSON Settings")
	ishtargate_input = flag.String("ishtargate-input", "", "Raw JSON IshtarGate")
	ishtargate_file  = flag.String("ishtargate-file", "ishtargate.json", "File of JSON IshtarGate")
	server           = flag.String("server", "", "Individual Server to Parse")
)

func main() {
	// parse flags
	if !flag.Parsed() {
		flag.Parse()
	}
	if *settings_input == "" && *settings_file == "" {
		log.Fatalln("settings input and settings file are both empty")
		return
	}
	if *ishtargate_input == "" && *ishtargate_file == "" {
		log.Fatalln("ishtargate input and ishtargate file are both empty")
		return
	}
	// settings
	settings := &ishtargate.Settings{}
	if *settings_input != "" {
		// parse input first if not empty
		if err := json.Unmarshal([]byte(*settings_input), settings); err != nil {
			log.Fatalf("failed to unmarshal settings input: \"%s\"\n", err)
			return
		}
	} else if *settings_file != "" {
		// parse file secondary
		bs, err := ioutil.ReadFile(*settings_file)
		if err != nil {
			log.Fatalf("failed to open settings file: \"%s\"\n", err)
			return
		}
		if err := json.Unmarshal(bs, settings); err != nil {
			log.Fatalf("failed to unmarshal settings file: \"%s\"\n", err)
			return
		}
	}
	// ishtargate
	ish := &ishtargate.IshtarGate{}
	if *ishtargate_input != "" {
		// parse input first if not empty
		if err := json.Unmarshal([]byte(*ishtargate_input), ish); err != nil {
			log.Fatalf("failed to unmarshal ishtargate input: \"%s\"\n", err)
			return
		}
	} else if *ishtargate_file != "" {
		// parse file secondary
		bs, err := ioutil.ReadFile(*ishtargate_file)
		if err != nil {
			log.Fatalf("failed to open ishtargate file: \"%s\"\n", err)
			return
		}
		if err := json.Unmarshal(bs, ish); err != nil {
			log.Fatalf("failed to unmarshal ishtargate file: \"%s\"\n", err)
			return
		}
	}
	if *server != "" {
		// parse individual server
		if !ish.BuildServer(settings, *server) {
			log.Fatalf("failed to build IshtarGate Server: \"%s\"\n", *server)
			return
		}
	} else {
		// parse all
		if !ish.Build(settings) {
			log.Fatalln("failed to build IshtarGate")
			return
		}
	}
	log.Println("success!")
}
