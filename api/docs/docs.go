// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-10-19 15:42:05.466075529 +0200 CEST m=+0.058501774

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "GetTrades",
        "title": "diadata.org API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "api.diadata.org",
    "basePath": "/",
    "paths": {
        "/kafka/filtersBlock": {
            "get": {
                "description": "GetFiltersBlock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kafka"
                ],
                "summary": "GetFiltersBlock get messages in topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset, default is last",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "number of elements",
                        "name": "elements",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dia.FiltersBlock"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/kafka/trades": {
            "get": {
                "description": "GetTrades",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kafka"
                ],
                "summary": "GetTrades get messages in topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset, default is last",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "number of elements",
                        "name": "elements",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dia.TradesBlock"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/kafka/tradesBlock": {
            "get": {
                "description": "GetTradesBlock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kafka"
                ],
                "summary": "GetTradesBlock get messages in topic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "offset, default is last",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "number of elements",
                        "name": "elements",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dia.TradesBlock"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/chartPoints/": {
            "get": {
                "description": "Get Symbol Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get chart points for",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Some exchange",
                        "name": "exchange",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Some filter",
                        "name": "filter",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "scale 5m 30m 1h 4h 1d 1w",
                        "name": "scale",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/diaApi.points"
                        }
                    },
                    "404": {
                        "description": "Symbol not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/chartPointsAllExchanges/": {
            "get": {
                "description": "Get Symbol Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get Symbol Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Some filter",
                        "name": "filter",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "scale 5m 30m 1h 4h 1d 1w",
                        "name": "scale",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/diaApi.points"
                        }
                    },
                    "404": {
                        "description": "Symbol not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/coins": {
            "get": {
                "description": "GetCoins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get coins",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/diaApi.Coins"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/pairs/": {
            "get": {
                "description": "et pairs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get pairs",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/diaApi.Pairs"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/quotation/": {
            "get": {
                "description": "GetQuotation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get quotation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Quotation"
                        }
                    },
                    "404": {
                        "description": "Symbol not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/supply": {
            "post": {
                "description": "Post the circulating supply",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Post the circulating supply",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coin symbol",
                        "name": "Symbol",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "number of coins in circulating supply",
                        "name": "CirculatingSupply",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dia.Supply"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/supply/": {
            "get": {
                "description": "GetSupply",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get supply",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/dia.Supply"
                        }
                    },
                    "404": {
                        "description": "Symbol not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        },
        "/v1/symbol/": {
            "get": {
                "description": "Get Symbol Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dia"
                ],
                "summary": "Get Symbol Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some symbol",
                        "name": "symbol",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/diaApi.SymbolDetails"
                        }
                    },
                    "404": {
                        "description": "Symbol not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/restApi.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dia.FilterPoint": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "dia.FiltersBlock": {
            "type": "object",
            "properties": {
                "blockHash": {
                    "type": "string"
                },
                "filtersBlockData": {
                    "type": "object",
                    "$ref": "#/definitions/dia.FiltersBlockData"
                }
            }
        },
        "dia.FiltersBlockData": {
            "type": "object",
            "properties": {
                "beginTime": {
                    "type": "string"
                },
                "endTime": {
                    "type": "string"
                },
                "filterPoints": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dia.FilterPoint"
                    }
                },
                "filtersNumber": {
                    "type": "integer"
                },
                "tradesBlockHash": {
                    "type": "string"
                }
            }
        },
        "dia.Pair": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string"
                },
                "foreignName": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "dia.Supply": {
            "type": "object",
            "properties": {
                "circulatingSupply": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "dia.Trade": {
            "type": "object",
            "properties": {
                "estimatedUSDPrice": {
                    "type": "number"
                },
                "foreignTradeID": {
                    "type": "string"
                },
                "pair": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "source": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "dia.TradesBlock": {
            "type": "object",
            "properties": {
                "blockHash": {
                    "type": "string"
                },
                "tradesBlockData": {
                    "type": "object",
                    "$ref": "#/definitions/dia.TradesBlockData"
                }
            }
        },
        "dia.TradesBlockData": {
            "type": "object",
            "properties": {
                "beginTime": {
                    "type": "string"
                },
                "endTime": {
                    "type": "string"
                },
                "trades": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dia.Trade"
                    }
                },
                "tradesNumber": {
                    "type": "integer"
                }
            }
        },
        "diaApi.Coin": {
            "type": "object",
            "properties": {
                "circulatingSupply": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "priceYesterday": {
                    "type": "number"
                },
                "symbol": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "volumeYesterdayUSD": {
                    "type": "number"
                }
            }
        },
        "diaApi.CoinSymbolAndName": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "diaApi.Coins": {
            "type": "object",
            "properties": {
                "change": {
                    "type": "object",
                    "$ref": "#/definitions/models.Change"
                },
                "coins": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/diaApi.Coin"
                    }
                },
                "completeCoinList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/diaApi.CoinSymbolAndName"
                    }
                }
            }
        },
        "diaApi.Pairs": {
            "type": "object",
            "properties": {
                "pairs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dia.Pair"
                    }
                }
            }
        },
        "diaApi.SymbolDetails": {
            "type": "object",
            "properties": {
                "change": {
                    "type": "object",
                    "$ref": "#/definitions/models.Change"
                },
                "coin": {
                    "type": "object",
                    "$ref": "#/definitions/diaApi.Coin"
                },
                "exchanges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SymbolExchangeDetails"
                    }
                }
            }
        },
        "diaApi.points": {
            "type": "object",
            "properties": {
                "dataPoints": {
                    "type": "string"
                }
            }
        },
        "models.Change": {
            "type": "object",
            "properties": {
                "usd": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CurrencyChange"
                    }
                }
            }
        },
        "models.CurrencyChange": {
            "type": "object",
            "properties": {
                "rate": {
                    "type": "number"
                },
                "rateYesterday": {
                    "type": "number"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "models.Quotation": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "priceYesterday": {
                    "type": "number"
                },
                "source": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "volumeYesterdayUSD": {
                    "type": "number"
                }
            }
        },
        "models.SymbolExchangeDetails": {
            "type": "object",
            "properties": {
                "lastTrades": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dia.Trade"
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "priceYesterday": {
                    "type": "number"
                },
                "time": {
                    "type": "string"
                },
                "volumeYesterdayUSD": {
                    "type": "number"
                }
            }
        },
        "restApi.APIError": {
            "type": "object",
            "properties": {
                "errorcode": {
                    "type": "integer"
                },
                "errormessage": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
