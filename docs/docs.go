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
            "name": "API Support",
            "url": "https://www.lifemiles.dev/support",
            "email": "jose.regalado@lifemiles.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/getCurrencies": {
            "get": {
                "description": "Retrieve a list of currencies and rates to USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Currency"
                ],
                "summary": "Return CurrencyResponse",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/currency.CurrencyResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/milesRate": {
            "post": {
                "description": "Retrieve Rate Per Store Response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AccRequest"
                ],
                "summary": "Return RatePerStoreResponse",
                "parameters": [
                    {
                        "description": "AccRequest",
                        "name": "accRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rate.AccRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rate.RatePerStoreResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "currency.Currency": {
            "type": "object",
            "properties": {
                "currencyCode": {
                    "type": "string"
                },
                "rateToUSD": {
                    "type": "number"
                }
            }
        },
        "currency.CurrencyResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/currency.Currency"
                    }
                },
                "code": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                }
            }
        },
        "rate.AccRequest": {
            "type": "object",
            "properties": {
                "partnerCode": {
                    "type": "string",
                    "example": "TOTCO"
                },
                "sender": {
                    "type": "string",
                    "example": "RBMCO"
                },
                "terminalID": {
                    "type": "string",
                    "example": "0010077923"
                }
            }
        },
        "rate.MilesRate": {
            "type": "object",
            "properties": {
                "accrualRate": {
                    "type": "integer"
                },
                "sender": {
                    "type": "string"
                },
                "terminalId": {
                    "type": "string"
                }
            }
        },
        "rate.RatePerStoreResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "object",
                    "$ref": "#/definitions/rate.MilesRate"
                },
                "code": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                }
            }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger lmgp-rates-svc API",
	Description: "This is the documentation from lmgp-rates-svc service.",
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