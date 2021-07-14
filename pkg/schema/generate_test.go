package schema

import (
	"testing"
)

const schemaPath = "/cypress/fixtures/newApp/newSchema.js"

func TestGenerateEmptyResponse(t *testing.T) {
	s, err := NewSchema(nil, schemaPath)
	if err == nil {
		t.Errorf("expected error for empty response, got: %s, err: %s", string(s.Data), err)
	}
}
