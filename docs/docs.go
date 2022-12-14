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
        "contact": {
            "name": "DuongBaoLong",
            "email": "longduongxx86@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "description": "gel all  todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todomodel.ToDoItem"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create  todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "newTodo",
                        "name": "dataItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todomodel.ToDoItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/todomodel.ToDoItem"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 8,
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/todomodel.ToDoItem"
                        }
                    }
                }
            },
            "put": {
                "description": "update  todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updateTodo",
                        "name": "dataItem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todomodel.ToDoItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/todomodel.ToDoItem"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "todomodel.ToDoItem": {
            "description": "User account information with user id and username",
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "is_show": {
                    "type": "boolean"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Todolist",
	Description:      "Todolist API using Swaggo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
