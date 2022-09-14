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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "내 상세정보를 얻는다.\n- Device-Token 헤더값이 필요하다.",
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
                "description": "회원가입\n- deviceToken 필드를 넣어주지 않으면 자동으로 생성된다. 만약 필드가 있다면 해당 값으로 생성한다.",
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
                        "description": "body",
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
                            "$ref": "#/definitions/controller.MemberResponse"
                        }
                    },
                    "409": {
                        "description": "Account already exist"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "내 정보를 업데이트 한다.\n- Device-Token 헤더값이 필요하다.\n- 이름만 업데이트 할경우 \"name\" 필드만, 키워드만 업데이트 할 경우 \"keywords\" 필드만 넘겨주면 된다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Update name or keywords",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.updateDTO"
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
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
        "/memos/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "내 메모 리스트 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memos"
                ],
                "summary": "Get user's Memo list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ListMemoResponse"
                        }
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "메모 생성",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memos"
                ],
                "summary": "Add Memo",
                "parameters": [
                    {
                        "description": "memo request body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.AddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.MemoResponse"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/memos/{memoID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "메모 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memos"
                ],
                "summary": "Retrieve Memo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "memoID to retrieve",
                        "name": "memoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.MemoResponse"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "메모 수정",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memos"
                ],
                "summary": "Update Memo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "memoID to retrieve",
                        "name": "memoID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "memo request body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.PatchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.MemoResponse"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/memos/{youtubeVideoID}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유니크 키인 (Device Token, YoutubeVideoID) 조합으로 메모 삭제",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "memos"
                ],
                "summary": "Delete Memo by a youtube video ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "youtubeVideoID to retrieve",
                        "name": "youtubeVideoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/musics": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                        "description": ""
                    }
                }
            }
        },
        "/musics/preference": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "내 취향 플레이리스트 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "musics"
                ],
                "summary": "Get my prefer music playlist",
                "parameters": [
                    {
                        "type": "string",
                        "name": "nextToken",
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
                        "description": ""
                    }
                }
            }
        },
        "/musics/search": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                        "type": "string",
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
                        "description": ""
                    }
                }
            }
        },
        "/tags/recommend": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "controller.AddRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "youtubeVideoID": {
                    "type": "string"
                }
            }
        },
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
        "controller.ListMemoData": {
            "type": "object",
            "properties": {
                "memos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.MemoResponse"
                    }
                }
            }
        },
        "controller.ListMemoResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/controller.ListMemoData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.ListMusicData": {
            "type": "object",
            "properties": {
                "musics": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.MusicResponse"
                    }
                },
                "nextPageToken": {
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
                    "$ref": "#/definitions/controller.ListMusicData"
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
        "controller.MemoResponse": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "keywords": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "memoID": {
                    "type": "string"
                },
                "thumbnailURL": {
                    "type": "string"
                },
                "title": {
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
        "controller.PatchRequest": {
            "type": "object",
            "properties": {
                "body": {
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
        },
        "controller.updateDTO": {
            "type": "object",
            "properties": {
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
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Device-Token",
            "in": "header"
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
