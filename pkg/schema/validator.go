package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/qri-io/jsonschema"
)

type Validator struct {
	schema  *jsonschema.Schema
	context context.Context
}

func NewValidator() *Validator {
	v := &Validator{}

	v.schema = &jsonschema.Schema{}

	if err := json.Unmarshal(getSchemaBytes(), v.schema); err != nil {
		fmt.Println("The Generoo configuration schema cannot be read. Please add an issue in the Generoo repository " +
			"here: https://github.com/army-of-one/generoo/issues. Thanks, and sorry for the inconvenience.")
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
