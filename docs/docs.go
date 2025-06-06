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
        "/api/Category/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/account": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get all accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.AccountListSwaggerResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.APIResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create new account",
                "parameters": [
                    {
                        "description": "Account data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.AccountListSwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/account/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get account by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.AccountListSwaggerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Update account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Account data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Delete account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/balance": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Get all balances",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.BalanceListSwaggerResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.APIResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/category": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.CategoryListSwaggerResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.APIResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create new category",
                "parameters": [
                    {
                        "description": "Category data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.CategoryListSwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/category/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CategoryListSwaggerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/logout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/report/daily/{date}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reports"
                ],
                "summary": "Get report daily",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Date (format: dd-mm-yyyy)",
                        "name": "date",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.TransactionListSwaggerResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.APIResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/transaction": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get all transactions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.TransactionListSwaggerResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.APIResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Create new transaction",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.TransactionListSwaggerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        },
        "/api/transaction/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get transaction by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.TransactionListSwaggerResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Update transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Delete transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AccountDTO": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                }
            }
        },
        "dto.BalanceDTO": {
            "type": "object",
            "properties": {
                "Balance": {
                    "type": "number"
                },
                "ID": {
                    "type": "integer"
                }
            }
        },
        "dto.CategoryDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.UserDTO": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                }
            }
        },
        "request.AccountRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "descriptions": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.AuthRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.CategoryRequest": {
            "type": "object",
            "required": [
                "name",
                "type"
            ],
            "properties": {
                "descriptions": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "pemasukan",
                        "pengeluaran",
                        "pindah"
                    ]
                }
            }
        },
        "request.TransactionRequest": {
            "type": "object",
            "required": [
                "account_id",
                "amount",
                "category_id",
                "date",
                "type"
            ],
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "category_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "descriptions": {
                    "type": "string"
                },
                "target_account_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "pemasukan",
                        "pengeluaran",
                        "pindah"
                    ]
                }
            }
        },
        "response.APIResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "meta": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "response.AccountListSwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AccountResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.AccountResponse": {
            "type": "object",
            "properties": {
                "Balance": {
                    "$ref": "#/definitions/dto.BalanceDTO"
                },
                "BalanceID": {
                    "type": "integer"
                },
                "Descriptions": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                }
            }
        },
        "response.BalanceListSwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.BalanceResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.BalanceResponse": {
            "type": "object",
            "properties": {
                "Account": {
                    "$ref": "#/definitions/dto.AccountDTO"
                },
                "AccountID": {
                    "type": "integer"
                },
                "Balance": {
                    "type": "number"
                },
                "ID": {
                    "type": "integer"
                }
            }
        },
        "response.CategoryListSwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.CategoryResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.CategoryResponse": {
            "type": "object",
            "properties": {
                "Descriptions": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "Type": {
                    "type": "string"
                }
            }
        },
        "response.TransactionListSwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.TransactionResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.TransactionResponse": {
            "type": "object",
            "properties": {
                "Account": {
                    "$ref": "#/definitions/dto.AccountDTO"
                },
                "AccountID": {
                    "type": "integer"
                },
                "Amount": {
                    "type": "number"
                },
                "Category": {
                    "$ref": "#/definitions/dto.CategoryDTO"
                },
                "CategoryID": {
                    "type": "integer"
                },
                "Date": {
                    "type": "string"
                },
                "Descriptions": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "TargetAccount": {
                    "$ref": "#/definitions/dto.AccountDTO"
                },
                "TargetAccountID": {
                    "type": "integer"
                },
                "Type": {
                    "type": "string"
                },
                "User": {
                    "$ref": "#/definitions/dto.UserDTO"
                },
                "UserID": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"https"},
	Title:            "BE CP2B - REST API Docs",
	Description:      "Simply cp2b system",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
