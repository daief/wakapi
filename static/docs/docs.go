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
            "name": "Ferdinand Mütsch",
            "url": "https://github.com/muety",
            "email": "ferdinand@muetsch.io"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://github.com/muety/wakapi/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/compat/shields/v1/{user}/{interval}/{filter}": {
            "get": {
                "description": "Retrieve total time for a given entity (e.g. a project) within a given range (e.g. one week) in a format compatible with [Shields.io](https://shields.io/endpoint). Requires public data access to be allowed.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "badges"
                ],
                "summary": "Get badge data",
                "operationId": "get-badge",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch data for",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "today",
                            "yesterday",
                            "week",
                            "month",
                            "year",
                            "7_days",
                            "last_7_days",
                            "30_days",
                            "last_30_days",
                            "12_months",
                            "last_12_months",
                            "any"
                        ],
                        "type": "string",
                        "description": "Interval to aggregate data for",
                        "name": "interval",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter to apply (e.g. 'project:wakapi' or 'language:Go')",
                        "name": "filter",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.BadgeData"
                        }
                    }
                }
            }
        },
        "/compat/wakatime/v1/users/{user}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mimics https://wakatime.com/developers#users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wakatime"
                ],
                "summary": "Retrieve the given user",
                "operationId": "get-wakatime-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch (or 'current')",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.UserViewModel"
                        }
                    }
                }
            }
        },
        "/compat/wakatime/v1/users/{user}/all_time_since_today": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mimics https://wakatime.com/developers#all_time_since_today",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wakatime"
                ],
                "summary": "Retrieve summary for all time",
                "operationId": "get-all-time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch data for (or 'current')",
                        "name": "user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.AllTimeViewModel"
                        }
                    }
                }
            }
        },
        "/compat/wakatime/v1/users/{user}/projects": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mimics https://wakatime.com/developers#projects",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wakatime"
                ],
                "summary": "Retrieve and fitler the user's projects",
                "operationId": "get-wakatime-projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch data for (or 'current')",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Query to filter projects by",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.ProjectsViewModel"
                        }
                    }
                }
            }
        },
        "/compat/wakatime/v1/users/{user}/stats/{range}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mimics https://wakatime.com/developers#stats",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wakatime"
                ],
                "summary": "Retrieve statistics for a given user",
                "operationId": "get-wakatimes-tats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch data for (or 'current')",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "today",
                            "yesterday",
                            "week",
                            "month",
                            "year",
                            "7_days",
                            "last_7_days",
                            "30_days",
                            "last_30_days",
                            "12_months",
                            "last_12_months",
                            "any"
                        ],
                        "type": "string",
                        "description": "Range interval identifier",
                        "name": "range",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.StatsViewModel"
                        }
                    }
                }
            }
        },
        "/compat/wakatime/v1/users/{user}/summaries": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mimics https://wakatime.com/developers#summaries.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wakatime"
                ],
                "summary": "Retrieve WakaTime-compatible summaries",
                "operationId": "get-wakatime-summaries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID to fetch data for (or 'current')",
                        "name": "user",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "today",
                            "yesterday",
                            "week",
                            "month",
                            "year",
                            "7_days",
                            "last_7_days",
                            "30_days",
                            "last_30_days",
                            "12_months",
                            "last_12_months",
                            "any"
                        ],
                        "type": "string",
                        "description": "Range interval identifier",
                        "name": "range",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start date (e.g. '2021-02-07')",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End date (e.g. '2021-02-08')",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.SummariesViewModel"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "misc"
                ],
                "summary": "Check the application's health status",
                "operationId": "get-health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/heartbeat": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "heartbeat"
                ],
                "summary": "Push a new heartbeat",
                "operationId": "post-heartbeat",
                "parameters": [
                    {
                        "description": "A heartbeat",
                        "name": "heartbeat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Heartbeat"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            }
        },
        "/summary": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "summary"
                ],
                "summary": "Retrieve a summary",
                "operationId": "get-summary",
                "parameters": [
                    {
                        "enum": [
                            "today",
                            "yesterday",
                            "week",
                            "month",
                            "year",
                            "7_days",
                            "last_7_days",
                            "30_days",
                            "last_30_days",
                            "12_months",
                            "last_12_months",
                            "any"
                        ],
                        "type": "string",
                        "description": "Interval identifier",
                        "name": "interval",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start date (e.g. '2021-02-07')",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End date (e.g. '2021-02-08')",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Whether to recompute the summary from raw heartbeat or use cache",
                        "name": "recompute",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Summary"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Heartbeat": {
            "type": "object",
            "properties": {
                "branch": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "description": "https://gorm.io/docs/conventions.html#CreatedAt",
                    "type": "number"
                },
                "editor": {
                    "description": "ignored because editor might be parsed differently by wakatime",
                    "type": "string"
                },
                "entity": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_write": {
                    "type": "boolean"
                },
                "language": {
                    "type": "string"
                },
                "machine": {
                    "description": "ignored because wakatime api doesn't return machines currently",
                    "type": "string"
                },
                "operating_system": {
                    "description": "ignored because os might be parsed differently by wakatime",
                    "type": "string"
                },
                "project": {
                    "type": "string"
                },
                "time": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Summary": {
            "type": "object",
            "properties": {
                "editors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SummaryItem"
                    }
                },
                "from": {
                    "type": "string",
                    "format": "date",
                    "example": "2006-01-02 15:04:05.000"
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SummaryItem"
                    }
                },
                "machines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SummaryItem"
                    }
                },
                "operating_systems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SummaryItem"
                    }
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SummaryItem"
                    }
                },
                "to": {
                    "type": "string",
                    "format": "date",
                    "example": "2006-01-02 15:04:05.000"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.SummaryItem": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "v1.AllTimeData": {
            "type": "object",
            "properties": {
                "is_up_to_date": {
                    "description": "true if the stats are up to date; when false, a 202 response code is returned and stats will be refreshed soon\u003e",
                    "type": "boolean"
                },
                "range": {
                    "$ref": "#/definitions/v1.AllTimeRange"
                },
                "text": {
                    "description": "total time logged since account created as human readable string\u003e",
                    "type": "string"
                },
                "total_seconds": {
                    "description": "total number of seconds logged since account created",
                    "type": "number"
                }
            }
        },
        "v1.AllTimeRange": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                }
            }
        },
        "v1.AllTimeViewModel": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/v1.AllTimeData"
                }
            }
        },
        "v1.BadgeData": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "schemaVersion": {
                    "type": "integer"
                }
            }
        },
        "v1.Project": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "repository": {
                    "type": "string"
                }
            }
        },
        "v1.ProjectsViewModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.Project"
                    }
                }
            }
        },
        "v1.StatsData": {
            "type": "object",
            "properties": {
                "daily_average": {
                    "type": "number"
                },
                "days_including_holidays": {
                    "type": "integer"
                },
                "editors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "end": {
                    "type": "string"
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "machines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "operating_systems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "start": {
                    "type": "string"
                },
                "total_seconds": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "v1.StatsViewModel": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/v1.StatsData"
                }
            }
        },
        "v1.SummariesData": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "dependencies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "editors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "grand_total": {
                    "$ref": "#/definitions/v1.SummariesGrandTotal"
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "machines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "operating_systems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesEntry"
                    }
                },
                "range": {
                    "$ref": "#/definitions/v1.SummariesRange"
                }
            }
        },
        "v1.SummariesEntry": {
            "type": "object",
            "properties": {
                "digital": {
                    "type": "string"
                },
                "hours": {
                    "type": "integer"
                },
                "minutes": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "percent": {
                    "type": "number"
                },
                "seconds": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "total_seconds": {
                    "type": "number"
                }
            }
        },
        "v1.SummariesGrandTotal": {
            "type": "object",
            "properties": {
                "digital": {
                    "type": "string"
                },
                "hours": {
                    "type": "integer"
                },
                "minutes": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "total_seconds": {
                    "type": "number"
                }
            }
        },
        "v1.SummariesRange": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "end": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                }
            }
        },
        "v1.SummariesViewModel": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.SummariesData"
                    }
                },
                "end": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                }
            }
        },
        "v1.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "display_name": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_email_confirmed": {
                    "type": "boolean"
                },
                "is_email_public": {
                    "type": "boolean"
                },
                "last_heartbeat_at": {
                    "type": "string"
                },
                "last_plugin_name": {
                    "type": "string"
                },
                "last_project": {
                    "type": "string"
                },
                "modified_at": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                }
            }
        },
        "v1.UserViewModel": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/v1.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Wakapi API",
	Description: "REST API to interact with [Wakapi](https://wakapi.dev)\n\n## Authentication\nSet header `Authorization` to your API Key encoded as Base64 and prefixed with `Basic`\n**Example:** `Basic ODY2NDhkNzQtMTljNS00NTJiLWJhMDEtZmIzZWM3MGQ0YzJmCg==`",
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
