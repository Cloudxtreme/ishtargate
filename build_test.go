package ishtargate

import (
	"encoding/json"
	"fmt"
	"sabey.co/unittest"
	"testing"
)

func TestBuild(t *testing.T) {
	fmt.Println("TestBuild")

	// you can build your own object here
	// or you can export your object and build your settings with the executable
	// once built, your files can be found in "settings.BuildPath/ishtar"

	// Settings Object
	settings := &Settings{
		BuildPath: "build",
	}
	unittest.NotNil(t, settings)

	// you can also export your objects like so:
	// Settings JSON
	bs, _ := json.Marshal(settings)
	unittest.Equals(t, string(bs), `{"build-path":"build"}`)
	fmt.Printf("Settings: \"%s\"\n", bs)

	// IshtarGate Object
	ish := &IshtarGate{}
	unittest.NotNil(t, ish)

	// build all
	// ish.Build(settings)

	// build individual server
	// ish.BuildServer(settings, "myserver")

	// IshtarGate JSON
	bs, _ = json.Marshal(ish)
	unittest.Equals(t, string(bs), "{}")
	fmt.Printf("IshtarGate: \"%s\"\n", bs)
}
