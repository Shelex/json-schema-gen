package fixture

import (
	"log"
	"path/filepath"

	"github.com/Shelex/json-schema-gen/pkg/prompt"
	"github.com/Shelex/json-schema-gen/pkg/schema"
	"github.com/spf13/afero"
)

// ResultsHandler - results from schema generator
type ResultsHandler struct {
	schema schema.Schema
	folder string
	prompt prompt.Executor
	os     afero.Afero
}

func NewResultsHandler(s schema.Schema, p prompt.Executor, os afero.Afero) ResultsHandler {
	return ResultsHandler{schema: s, folder: fixtureFolder, prompt: p, os: os}
}

// Write - write schema to file, add configuration
func (s *ResultsHandler) Write() error {
	isNewSchema := s.isNewSchema()
	if !isNewSchema {
		overwrite, newName, err := overwrite(s.prompt)
		if err != nil {
			return err
		}
		if !overwrite {
			s.schema.Title = newName
		}
	}
	return s.writeSchema()
}

func (s *ResultsHandler) isNewSchema() bool {
	schemaPath := filepath.Join(s.folder, s.schema.Title+".json")
	if exist, err := s.os.Exists(schemaPath); err != nil || exist {
		return false
	}
	return true
}

func (s *ResultsHandler) writeSchema() error {
	fileName := s.schema.Title + ".json"
	path := filepath.Join(s.folder, fileName)
	if err := s.os.WriteFile(path, s.schema.Data, 0644); err != nil {
		return err
	}
	log.Printf("Saved schema to: %s", path)
	return nil
}
