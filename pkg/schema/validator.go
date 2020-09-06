package schema

import (
	"context"
	"encoding/json"
	"github.com/qri-io/jsonschema"
	"log"
)

type Validator struct {
	schema  *jsonschema.Schema
	context context.Context
}

func NewValidator() *Validator {
	v := &Validator{}

	v.schema = &jsonschema.Schema{}

	if err := json.Unmarshal(getSchemaBytes(), v.schema); err != nil {
		log.Fatalf("failed to unmarshal schema: %s", err)
	}

	v.context = context.Background()

	return v
}

func (v *Validator) Validate(rawJson []byte) ([]jsonschema.KeyError, error) {
	errs, err := v.schema.ValidateBytes(v.context, rawJson)

	if err != nil {
		return nil, err
	}

	return errs, nil

}
