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
            "name": "Podgrab Github",
            "url": "https://www.github.com/akhilrex/podgrab"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/podcasts": {
            "get": {
                "description": "Get all Podcasts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all Podcasts",
                "operationId": "get-all-podcasts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sort by property",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by asc/desc",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Podcast"
                            }
                        }
                    }
                }
            }
        },
        "/podcasts/{id}": {
            "get": {
                "description": "Get single podcast by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get single podcast by ID",
                "operationId": "get-podcast-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Podcast id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Podcast"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete single podcast by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete single podcast by ID",
                "operationId": "delete-podcast-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Podcast id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Podcast": {
            "type": "object",
            "properties": {
                "allEpisodesCount": {
                    "type": "integer"
                },
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "downloadedEpisodesCount": {
                    "type": "integer"
                },
                "downloadingEpisodesCount": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "lastEpisode": {
                    "type": "string"
                },
                "podcastItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.PodcastItem"
                    }
                },
                "summary": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Tag"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "db.PodcastItem": {
            "type": "object",
            "properties": {
                "bookmarkDate": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "downloadDate": {
                    "type": "string"
                },
                "downloadPath": {
                    "type": "string"
                },
                "downloadStatus": {
                    "type": "integer"
                },
                "duration": {
                    "type": "integer"
                },
                "episodeType": {
                    "type": "string"
                },
                "fileURL": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isPlayed": {
                    "type": "boolean"
                },
                "localImage": {
                    "type": "string"
                },
                "podcast": {
                    "$ref": "#/definitions/db.Podcast"
                },
                "podcastID": {
                    "type": "string"
                },
                "pubDate": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "db.Tag": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "podcasts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Podcast"
                    }
                },
                "updatedAt": {
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
	Title:       "Podgrab API",
	Description: "This is the api documentation for Podgrab.",
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