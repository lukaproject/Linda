// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bags": {
            "get": {
                "description": "list bag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list bag [no implementation]",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.ListBagsResp"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new bag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add bag",
                "parameters": [
                    {
                        "description": "bag's request",
                        "name": "addBagReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.AddBagReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.AddBagResp"
                        }
                    }
                }
            }
        },
        "/bags/{bagName}": {
            "get": {
                "description": "get bag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get bag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bag's name",
                        "name": "bagName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.GetBagResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete bag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "delete bag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bag's name",
                        "name": "bagName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.DeleteBagResp"
                        }
                    }
                }
            }
        },
        "/bags/{bagName}/tasks": {
            "post": {
                "description": "add task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "add task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bag's name",
                        "name": "bagName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "add tasks's request",
                        "name": "addTaskReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.AddTaskReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.AddTaskResp"
                        }
                    }
                }
            }
        },
        "/bags/{bagName}/tasks/{taskName}": {
            "get": {
                "description": "get task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bag's name",
                        "name": "bagName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "task's name",
                        "name": "taskName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.GetTaskResp"
                        }
                    }
                }
            }
        },
        "/healthcheck": {
            "post": {
                "description": "health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "summary": "health check",
                "responses": {}
            }
        }
    },
    "definitions": {
        "apis.AddBagReq": {
            "type": "object",
            "properties": {
                "bagDisplayName": {
                    "type": "string",
                    "example": "test-bagDisplayName"
                }
            }
        },
        "apis.AddBagResp": {
            "type": "object",
            "properties": {
                "bagDisplayName": {
                    "type": "string"
                },
                "bagName": {
                    "type": "string"
                },
                "createTimeMs": {
                    "type": "integer"
                },
                "updateTimeMs": {
                    "type": "integer"
                }
            }
        },
        "apis.AddTaskReq": {
            "type": "object",
            "properties": {
                "scriptPath": {
                    "type": "string",
                    "example": "/bin/test.sh"
                },
                "taskDisplayName": {
                    "type": "string",
                    "example": "test-taskDisplayName"
                },
                "workingDir": {
                    "type": "string",
                    "example": "/bin/testWorkingDir/working"
                }
            }
        },
        "apis.AddTaskResp": {
            "type": "object",
            "properties": {
                "bagName": {
                    "type": "string"
                },
                "createTimeMs": {
                    "type": "integer"
                },
                "finishTimeMs": {
                    "type": "integer"
                },
                "nodeId": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "scheduledTimeMs": {
                    "type": "integer"
                },
                "scriptPath": {
                    "type": "string"
                },
                "taskDisplayName": {
                    "type": "string"
                },
                "taskName": {
                    "type": "string"
                },
                "workingDir": {
                    "type": "string"
                }
            }
        },
        "apis.Bag": {
            "type": "object",
            "properties": {
                "bagDisplayName": {
                    "type": "string"
                },
                "bagName": {
                    "type": "string"
                },
                "createTimeMs": {
                    "type": "integer"
                },
                "updateTimeMs": {
                    "type": "integer"
                }
            }
        },
        "apis.DeleteBagResp": {
            "type": "object",
            "properties": {
                "errorMsg": {
                    "type": "string"
                }
            }
        },
        "apis.GetBagResp": {
            "type": "object",
            "properties": {
                "bagDisplayName": {
                    "type": "string"
                },
                "bagName": {
                    "type": "string"
                },
                "createTimeMs": {
                    "type": "integer"
                },
                "updateTimeMs": {
                    "type": "integer"
                }
            }
        },
        "apis.GetTaskResp": {
            "type": "object",
            "properties": {
                "bagName": {
                    "type": "string"
                },
                "createTimeMs": {
                    "type": "integer"
                },
                "finishTimeMs": {
                    "type": "integer"
                },
                "nodeId": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "scheduledTimeMs": {
                    "type": "integer"
                },
                "scriptPath": {
                    "type": "string"
                },
                "taskDisplayName": {
                    "type": "string"
                },
                "taskName": {
                    "type": "string"
                },
                "workingDir": {
                    "type": "string"
                }
            }
        },
        "apis.ListBagsResp": {
            "type": "object",
            "properties": {
                "bags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/apis.Bag"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.dev",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "AgentCentral API",
	Description:      "This is agent central swagger API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
