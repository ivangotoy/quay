// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jonathan King",
            "email": "joking@redhat.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "description": "This endpoint will load the config bundle mounted by the config-tool into memory. This state can then be modified, validated, downloaded, and optionally committed to a Quay operator instance.",
                "produces": [
                    "application/json"
                ],
                "summary": "Returns the mounted config bundle.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/editor.ConfigBundle"
                        }
                    }
                }
            }
        },
        "/config/download": {
            "post": {
                "description": "This endpoint will download the config bundle in the request body as a tar.gz",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "summary": "Downloads a config bundle as a tar.gz",
                "parameters": [
                    {
                        "description": "JSON Representing Config Bundle",
                        "name": "configBundle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/editor.ConfigBundle"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/config/operator": {
            "post": {
                "description": "Handles an HTTP POST request containing a new ` + "`" + `config.yaml` + "`" + `, adds any uploaded certs, and calls an API endpoint on the Quay Operator to create a new ` + "`" + `Secret` + "`" + `.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Commits a config bundle to a Quay operator instance.",
                "parameters": [
                    {
                        "description": "JSON Representing Config Bundle",
                        "name": "configBundle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/editor.ConfigBundle"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/config/validate": {
            "post": {
                "description": "This endpoint will validate the config bundle contained in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Validates a config bundle.",
                "parameters": [
                    {
                        "description": "JSON Representing Config Bundle",
                        "name": "configBundle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/editor.ConfigBundle"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "editor.ConfigBundle": {
            "type": "object",
            "properties": {
                "certs": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "config.yaml": {
                    "type": "object",
                    "additionalProperties": true
                },
                "managedFieldGroups": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{"http"},
	Title:       "Config Tool Editor API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
