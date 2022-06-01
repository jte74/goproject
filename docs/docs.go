// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "julien.tellier@openwt.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "login user",
                        "name": "users",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/restricted/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Return All Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "id"
                ],
                "summary": "Return All Users",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.User": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "Age\nin: int32",
                    "type": "integer"
                },
                "date_created": {
                    "description": "DateCreated\nin: date",
                    "type": "string"
                },
                "firstname": {
                    "description": "Firstname\nin: string",
                    "type": "string"
                },
                "id": {
                    "description": "Id\nin: int32",
                    "type": "integer"
                },
                "name": {
                    "description": "Name\nin: string",
                    "type": "string"
                },
                "password": {
                    "description": "Password\nin: string",
                    "type": "string"
                },
                "token": {
                    "description": "Token\nin: string",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Article API",
	Description:      "This is a sample service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
