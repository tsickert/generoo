package schema

func getSchemaBytes() []byte {
	return []byte(`
{
  "$schema": "http://json-schema.org/draft-07/schema",
  "$id": "http://generoo.armyofone.tech/schema.json",

  "title": "Generoo Configuration File Schema",
  "description": "The configuration schema for generoo generation. This file is used to define the flow of Generoo template rendering.",

  "type": "object",
  "required": [
	"version"
  ],
  "properties": {
	"version": { "$ref": "#/$defs/version" },
	"name": { "$ref": "#/$defs/name" },
	"description": { "$ref": "#/$defs/description" },
    "constants": {
      "type": "array",
      "title": "The inputs schema",
      "description": "Inputs are values that are taken from the user to rendering in the template.",
      "default": [],
      "items": {
        "$ref": "#/$defs/constant"
      }
    },
    "inputs": {
      "type": "array",
      "title": "The inputs schema",
      "description": "Inputs are values that are taken from the user to rendering in the template.",
      "default": [],
      "items": {
        "$ref": "#/$defs/input"
      }
    }
  },

  "$defs": {
    "constant": {
      "type": "object",
      "title": "The name schema",
      "description": "An input to be rendered that does not need to be taken from the user.",
      "properties": {
        "name": { "$ref": "#/$defs/name" },
        "value": { "$ref": "#/$defs/value" }
      },
      "required": [
        "name",
        "value"
      ]
    },
    "input": {
      "type": "object",
      "title": "The inputs schema",
      "description": "Input required from user.",
      "required": [
        "text",
        "type"
      ],
      "properties": {
        "name": { "$ref": "#/$defs/name" },
        "text": { "$ref": "#/$defs/text" },
        "type": { "$ref": "#/$defs/type" },
        "default": { "$ref": "#/$defs/default" },
		"options": {
          "type": "array",
          "title": "The options schema",
          "description": "List of valid options that the user can pick from.",
          "default": [],
          "items": { 
			"type": "string"
		  }
        },
        "validations": {
          "type": "array",
          "title": "The validations schema",
          "description": "List of validations to perform on the user input.",
          "default": [],
          "items": { "$ref": "#/$defs/validation" }
        },
        "follow_ups": {
          "type": "array",
          "title": "The follow_ups schema",
          "description": "Inputs that become required based on the user input.",
          "default": [],
          "items": { "$ref": "#/$defs/input" }
        }
      }
    },

    "validation": {
      "type": "object",
      "title": "The first anyOf schema",
      "description": "Validations are checked before rendering occurs to make sure valid inputs are provided.",
      "default": {},
      "required": [
        "type",
        "value"
      ],
      "properties": {
        "type": { "$ref": "#/$defs/type" },
        "value": { "$ref": "#/$defs/value" }
      }
    },

    "name": {
      "type": "string",
      "title": "The name schema",
      "description": "The name of an input. This could be used as a reference later."
    },
    "text": {
      "type": "string",
      "title": "The text schema",
      "description": "The text that a user sees during the prompt phase.",
      "default": ""
    },
    "type": {
      "type": "string",
      "title": "The type schema",
      "description": "The variable type of the input. This is used for type coercion and validation.",
      "default": ""
    },
    "default": {
      "type": "string",
      "title": "The default schema",
      "description": "The default value an option can have. This will be shown in the prompt.",
      "default": ""
    },
    "value": {
      "type": "string",
      "title": "The value schema",
      "description": "The value against which the validation takes place."
    },
	"version": {
	  "type": "string",
	  "pattern": "^(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$",
 	  "title": "The version schema",
	  "description": "Semantic version of the template."
	}
  }
}
	`)
}
