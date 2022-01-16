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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:id": {
            "get": {
                "description": "return route object by route id or error if there is not one",
                "tags": [
                    "routes"
                ],
                "summary": "GetRouteByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "route ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routers.RouteModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    }
                }
            }
        },
        "/:id/passports": {
            "get": {
                "description": "return all passports are belonged the route",
                "tags": [
                    "routes"
                ],
                "summary": "GetRoutePassports",
                "parameters": [
                    {
                        "type": "string",
                        "description": "route ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routers.RoutePassportsModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    }
                }
            }
        },
        "/:passportId": {
            "get": {
                "description": "return passport by ID from database if there is one, or return error object with status code 200",
                "tags": [
                    "passports"
                ],
                "summary": "GetPassportByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "passport ID",
                        "name": "passportId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passport.Model"
                        }
                    }
                }
            }
        },
        "/:passportId/towers": {
            "get": {
                "description": "return all towers of passport by id",
                "tags": [
                    "towers"
                ],
                "summary": "GetTowersOfPassport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "passport ID",
                        "name": "passportId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passport.TowersModel"
                        }
                    }
                }
            }
        },
        "/:passportId/towers/:towerId": {
            "get": {
                "description": "return certain tower of the passport by ID",
                "tags": [
                    "towers"
                ],
                "summary": "GetPassportTowerByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "passport ID",
                        "name": "passportId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tower ID",
                        "name": "towerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passport.TowerModel"
                        }
                    }
                }
            }
        },
        "/:passportId/towers/findtower": {
            "get": {
                "description": "return the closest tower belonged the passport by coordinates",
                "tags": [
                    "towers"
                ],
                "summary": "FindTowerByCoordinate",
                "parameters": [
                    {
                        "type": "number",
                        "description": "latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passport.TowerModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    }
                }
            }
        },
        "/all": {
            "get": {
                "description": "return all routes from database",
                "tags": [
                    "routes"
                ],
                "summary": "GetAllRoutes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routers.ListRoutesModel"
                        }
                    }
                }
            }
        },
        "/passport": {
            "post": {
                "description": "save passport in database",
                "tags": [
                    "passports"
                ],
                "summary": "SavePassport",
                "parameters": [
                    {
                        "description": "xml structure of passport",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/passport.Model"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passport.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    }
                }
            }
        },
        "/route": {
            "post": {
                "description": "Save route in database",
                "tags": [
                    "routes"
                ],
                "summary": "Save",
                "parameters": [
                    {
                        "description": "xml doc of route",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routers.RouteModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routers.RouteModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errs.ErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errs.ErrorModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "passport.ExpirationModel": {
            "type": "object",
            "properties": {
                "daysUntilExpiration": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "passport.HeaderModel": {
            "type": "object",
            "properties": {
                "CHANGEDATE": {
                    "type": "string"
                },
                "Sequence": {
                    "type": "string"
                },
                "currentWay": {
                    "type": "string"
                },
                "currentWayID": {
                    "type": "string"
                },
                "echName": {
                    "type": "string"
                },
                "echkName": {
                    "type": "string"
                },
                "initialKM": {
                    "type": "string"
                },
                "initialM": {
                    "type": "string"
                },
                "initialMeter": {
                    "type": "string"
                },
                "initialPK": {
                    "type": "string"
                },
                "locationId": {
                    "type": "string"
                },
                "plotLength": {
                    "type": "string"
                },
                "sectionId": {
                    "type": "string"
                },
                "sectionName": {
                    "type": "string"
                },
                "siteId": {
                    "type": "string"
                },
                "suspensionAmount": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "wayAmount": {
                    "type": "string"
                },
                "workType": {
                    "type": "string"
                }
            }
        },
        "passport.Model": {
            "type": "object",
            "properties": {
                "Expiration": {
                    "type": "object",
                    "$ref": "#/definitions/passport.ExpirationModel"
                },
                "Header": {
                    "type": "object",
                    "$ref": "#/definitions/passport.HeaderModel"
                },
                "Towers": {
                    "type": "object",
                    "$ref": "#/definitions/passport.TowersModel"
                },
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "passport.TowerModel": {
            "type": "object",
            "properties": {
                "CountWire": {
                    "type": "string"
                },
                "Gabarit": {
                    "type": "string"
                },
                "RADIUS": {
                    "type": "string"
                },
                "SPEED": {
                    "type": "string"
                },
                "TF_TYPE": {
                    "type": "string"
                },
                "TURN": {
                    "type": "string"
                },
                "WireType": {
                    "type": "string"
                },
                "assetNum": {
                    "type": "string"
                },
                "catenary": {
                    "type": "string"
                },
                "distance": {
                    "type": "string"
                },
                "grounded": {
                    "type": "string"
                },
                "height": {
                    "type": "string"
                },
                "idtf": {
                    "type": "string"
                },
                "km": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                },
                "mapper": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "offset": {
                    "type": "string"
                },
                "pk": {
                    "type": "string"
                },
                "stopSeq": {
                    "type": "string"
                },
                "suspensionType": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "zigzag": {
                    "type": "string"
                }
            }
        },
        "passport.TowersModel": {
            "type": "object",
            "properties": {
                "TowerModel": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/passport.TowerModel"
                    }
                },
                "sectionID": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "routers.InformationRouteModel": {
            "type": "object",
            "properties": {
                "car": {
                    "type": "string"
                },
                "carID": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "eigthnum": {
                    "type": "string"
                },
                "masterPmNum": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "tripChangeData": {
                    "type": "string"
                },
                "tripType": {
                    "type": "string"
                },
                "viksRouteID": {
                    "type": "string"
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "routers.ListRoutesModel": {
            "type": "object",
            "properties": {
                "routeModel": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/routers.InformationRouteModel"
                    }
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "routers.RouteModel": {
            "type": "object",
            "properties": {
                "car": {
                    "type": "string"
                },
                "carID": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "eigthnum": {
                    "type": "string"
                },
                "masterPmNum": {
                    "type": "string"
                },
                "sectionSetModel": {
                    "type": "object",
                    "$ref": "#/definitions/routers.SectionSetModel"
                },
                "text": {
                    "type": "string"
                },
                "tripChangeData": {
                    "type": "string"
                },
                "tripType": {
                    "type": "string"
                },
                "viksRouteID": {
                    "type": "string"
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "routers.RoutePassportsModel": {
            "type": "object",
            "properties": {
                "passports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/passport.Model"
                    }
                },
                "viksRouteID": {
                    "type": "string"
                },
                "xmlname": {
                    "type": "string"
                }
            }
        },
        "routers.SectionModel": {
            "type": "object",
            "properties": {
                "changeData": {
                    "type": "string"
                },
                "sectionId": {
                    "type": "string"
                },
                "sectionName": {
                    "type": "string"
                },
                "sequence": {
                    "type": "string"
                },
                "siteId": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "workType": {
                    "type": "string"
                }
            }
        },
        "routers.SectionSetModel": {
            "type": "object",
            "properties": {
                "section": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/routers.SectionModel"
                    }
                },
                "text": {
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
	Host:        "localhost:80",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Passports and Routes of railways store server",
	Description: "Store server for passports and routes.",
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
