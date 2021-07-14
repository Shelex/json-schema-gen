package main

import (
	"flag"
	"log"

	"github.com/Shelex/json-schema-gen/pkg/fixture"
	"github.com/Shelex/json-schema-gen/pkg/prompt"
	"github.com/Shelex/json-schema-gen/pkg/schema"
	"github.com/spf13/afero"
)

func main() {
	var selectedFixture, baseURL string
	flag.StringVar(&selectedFixture, "fixture", "", "selected fixture path")
	flag.StringVar(&baseURL, "url", "", "base url")
	flag.Parse()

	os := afero.Afero{Fs: afero.NewOsFs()}

	prompt := &prompt.UI{Validate: fixture.Validate}
	fixtureFile, err := fixture.Select(prompt, os, selectedFixture)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := schema.NewSchema(fixtureFile.Content, fixtureFile.Path)
	if err != nil {
		log.Fatalf("failed to generate schema: %s\n", err)
	}

	results := fixture.NewResultsHandler(schema, prompt, os)
	if err := results.Write(); err != nil {
		log.Fatalf("failed to handle results: %s\n", err)
	}
}
