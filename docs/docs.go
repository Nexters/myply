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
            "email": "minkj1992@gmail.com"
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
        "/members/": {
            "get": {
                "description": "get myply member defail information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.MemberResponse"
                        }
                    },
                    "401": {
                        "description": "Unautorized"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "join myply members",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "sign up body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.signUpDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.BaseResponse"
                        }
                    },
                    "409": {
                        "description": "Account already exist"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/musics": {
            "get": {
                "description": "플레이리스트 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "musics"
                ],
                "summary": "Retrieve music playlist",
                "parameters": [
                    {
                        "type": "string",
                        "name": "nextToken",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListMusicResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/musics/search": {
            "get": {
                "description": "플레이리스트 검색",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "musics"
                ],
                "summary": "Search music playlist",
                "parameters": [
                    {
                        "type": "string",
                        "name": "nextToken",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListMusicResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/tags/recommend": {
            "get": {
                "description": "get tags recommended by myply",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Get recommended tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.RecommendResponse"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.BaseResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.ListMusicResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.MusicResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.MemberResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/controller.memberResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.MusicResponse": {
            "type": "object",
            "properties": {
                "isMemoed": {
                    "type": "boolean"
                },
                "thumbnailURL": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "videoDeepLink": {
                    "type": "string"
                },
                "youtubeTags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "youtubeVideoID": {
                    "type": "string"
                }
            }
        },
        "controller.RecommendResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/controller.RecommendResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.RecommendResponseData": {
            "type": "object",
            "properties": {
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controller.memberResponseData": {
            "type": "object",
            "properties": {
                "deviceToken": {
                    "type": "string"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.signUpDTO": {
            "type": "object",
            "properties": {
                "deviceToken": {
                    "type": "string"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
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
	Title:            "MYPLY SERVER",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
