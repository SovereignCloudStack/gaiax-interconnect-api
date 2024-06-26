// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/vpn/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VPNs"
                ],
                "summary": "Add a specific VPN",
                "parameters": [
                    {
                        "description": "VPN name",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "VPN type",
                        "name": "type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "VPN localAsNumber",
                        "name": "localAsNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "VPN remoteAsNumber",
                        "name": "remoteAsNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "VPN vni",
                        "name": "vni",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.IDResponse"
                        }
                    }
                }
            }
        },
        "/api/vpn/delete/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VPNs"
                ],
                "summary": "Delete a specific VPN",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VPN ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetVPNResponse"
                        }
                    }
                }
            }
        },
        "/api/vpn/update": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VPNs"
                ],
                "summary": "Update a specific VPN",
                "parameters": [
                    {
                        "description": "VPN name",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "VPN type",
                        "name": "type",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "VPN localAsNumber",
                        "name": "localAsNumber",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "VPN remoteAsNumber",
                        "name": "remoteAsNumber",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "VPN vni",
                        "name": "vni",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.IDResponse"
                        }
                    }
                }
            }
        },
        "/api/vpn/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VPNs"
                ],
                "summary": "Get a specific VPN",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VPN ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetVPNResponse"
                        }
                    }
                }
            }
        },
        "/api/vpns": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "VPNs"
                ],
                "summary": "Get all VPNs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.GetVPNResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GetVPNResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "local_as_number": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "remote_as_number": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "vni": {
                    "type": "integer"
                }
            }
        },
        "model.IDResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go Rest Api",
	Description:      "Api Endpoints for Go Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
