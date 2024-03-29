// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the sbp system of self-service car washes",
    "title": "wash-sbp",
    "version": "1.0.1"
  },
  "paths": {
    "/groups/{groupId}/washes/{washId}": {
      "post": {
        "tags": [
          "washes",
          "groups"
        ],
        "operationId": "assignWashToGroup",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "washId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/healthcheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "standard"
        ],
        "operationId": "healthcheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/notification": {
      "post": {
        "security": [
          {}
        ],
        "tags": [
          "notifications"
        ],
        "operationId": "receiveNotification",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Notification"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/payments/cancel": {
      "post": {
        "tags": [
          "payments"
        ],
        "operationId": "cancelPayment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PaymentCancellation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/payments/init": {
      "post": {
        "tags": [
          "payments"
        ],
        "operationId": "initPayment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Payment"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PaymentResponse"
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/transactions": {
      "get": {
        "tags": [
          "transactions"
        ],
        "operationId": "getTransactions",
        "parameters": [
          {
            "$ref": "#/parameters/page"
          },
          {
            "$ref": "#/parameters/pageSize"
          },
          {
            "enum": [
              "new",
              "authorized",
              "confirmed_not_synced",
              "confirmed",
              "canceling",
              "canceled",
              "refunded",
              "unknown"
            ],
            "type": "string",
            "name": "status",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "organizationId",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "washId",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "postId",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/TransactionPage"
            }
          },
          "403": {
            "$ref": "#/responses/GenericError"
          },
          "404": {
            "$ref": "#/responses/GenericError"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/washes": {
      "get": {
        "tags": [
          "washes"
        ],
        "operationId": "getWashes",
        "parameters": [
          {
            "$ref": "#/parameters/offset"
          },
          {
            "$ref": "#/parameters/limit"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Wash"
              }
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "post": {
        "tags": [
          "washes"
        ],
        "operationId": "createWash",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashCreation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success creation",
            "schema": {
              "$ref": "#/definitions/Wash"
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    },
    "/washes/{id}": {
      "get": {
        "tags": [
          "washes"
        ],
        "operationId": "getWashById",
        "parameters": [
          {
            "$ref": "#/parameters/uuid"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Wash"
            }
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "delete": {
        "tags": [
          "washes"
        ],
        "operationId": "deleteWash",
        "parameters": [
          {
            "$ref": "#/parameters/uuid"
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      },
      "patch": {
        "tags": [
          "washes"
        ],
        "operationId": "updateWash",
        "parameters": [
          {
            "$ref": "#/parameters/uuid"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashUpdate"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success update"
          },
          "default": {
            "$ref": "#/responses/GenericError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Group": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Notification": {
      "type": "object",
      "properties": {
        "Amount": {
          "description": "Payment amount",
          "type": "integer"
        },
        "CardId": {
          "type": "integer",
          "x-nullable": true
        },
        "ErrorCode": {
          "description": "Error code",
          "type": "string"
        },
        "ExpDate": {
          "type": "string",
          "x-nullable": true
        },
        "OrderId": {
          "description": "Order ID",
          "type": "string"
        },
        "Pan": {
          "description": "PAN (Primary Account Number)",
          "type": "string"
        },
        "PaymentId": {
          "description": "Payment ID",
          "type": "integer"
        },
        "Status": {
          "description": "Payment status",
          "type": "string"
        },
        "Success": {
          "description": "Indicates whether the payment was successful",
          "type": "boolean"
        },
        "TerminalKey": {
          "description": "Terminal key",
          "type": "string"
        },
        "Token": {
          "description": "Payment token",
          "type": "string"
        }
      }
    },
    "Organization": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Payment": {
      "type": "object",
      "properties": {
        "amount": {
          "type": "integer"
        },
        "orderId": {
          "type": "string"
        },
        "postId": {
          "type": "string"
        },
        "washId": {
          "type": "string"
        }
      }
    },
    "PaymentCancellation": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "postID": {
          "type": "string"
        },
        "washID": {
          "type": "string"
        }
      }
    },
    "PaymentResponse": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "SimpleWash": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "id",
        "createdAt",
        "amount",
        "status",
        "postId",
        "wash",
        "group",
        "organization"
      ],
      "properties": {
        "amount": {
          "type": "integer",
          "format": "int64"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "group": {
          "$ref": "#/definitions/Group"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "organization": {
          "$ref": "#/definitions/Organization"
        },
        "postId": {
          "type": "integer"
        },
        "status": {
          "$ref": "#/definitions/TransactionStatus"
        },
        "wash": {
          "$ref": "#/definitions/SimpleWash"
        }
      }
    },
    "TransactionPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        },
        "page": {
          "type": "integer",
          "minimum": 1
        },
        "pageSize": {
          "type": "integer",
          "maximum": 100,
          "minimum": 1
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    },
    "TransactionStatus": {
      "type": "string",
      "enum": [
        "new",
        "authorized",
        "confirmed_not_synced",
        "confirmed",
        "canceling",
        "canceled",
        "refunded",
        "unknown"
      ]
    },
    "User": {
      "description": "User profile",
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "organization": {
          "type": "object",
          "properties": {
            "deleted": {
              "type": "boolean"
            },
            "description": {
              "type": "string"
            },
            "displayName": {
              "type": "string"
            },
            "id": {
              "type": "string",
              "format": "uuid"
            },
            "name": {
              "type": "string"
            }
          },
          "x-nullable": true
        },
        "role": {
          "$ref": "#/definitions/UserRole"
        }
      }
    },
    "UserRole": {
      "type": "string",
      "enum": [
        "systemManager",
        "admin",
        "noAccess"
      ]
    },
    "Wash": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "groupId": {
          "type": "string",
          "format": "uuid"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "organizationId": {
          "type": "string",
          "format": "uuid"
        },
        "password": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    },
    "WashCreation": {
      "required": [
        "name",
        "description",
        "terminalKey",
        "terminalPassword",
        "groupId"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "groupId": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    },
    "WashUpdate": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "limit": {
      "minimum": 1,
      "type": "integer",
      "format": "int64",
      "default": 100,
      "description": "Maximum number of records to return",
      "name": "limit",
      "in": "query"
    },
    "offset": {
      "type": "integer",
      "format": "int64",
      "default": 0,
      "description": "Number of records to skip for pagination",
      "name": "offset",
      "in": "query"
    },
    "page": {
      "minimum": 1,
      "type": "integer",
      "default": 1,
      "name": "page",
      "in": "query"
    },
    "pageSize": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 10,
      "name": "pageSize",
      "in": "query"
    },
    "userRole": {
      "enum": [
        "systemManager",
        "admin",
        "noAccess"
      ],
      "type": "string",
      "name": "role",
      "in": "query"
    },
    "uuid": {
      "type": "string",
      "format": "uuid",
      "name": "id",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the sbp system of self-service car washes",
    "title": "wash-sbp",
    "version": "1.0.1"
  },
  "paths": {
    "/groups/{groupId}/washes/{washId}": {
      "post": {
        "tags": [
          "washes",
          "groups"
        ],
        "operationId": "assignWashToGroup",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "washId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/healthcheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "standard"
        ],
        "operationId": "healthcheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/notification": {
      "post": {
        "security": [
          {}
        ],
        "tags": [
          "notifications"
        ],
        "operationId": "receiveNotification",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Notification"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/payments/cancel": {
      "post": {
        "tags": [
          "payments"
        ],
        "operationId": "cancelPayment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PaymentCancellation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/payments/init": {
      "post": {
        "tags": [
          "payments"
        ],
        "operationId": "initPayment",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Payment"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/PaymentResponse"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/transactions": {
      "get": {
        "tags": [
          "transactions"
        ],
        "operationId": "getTransactions",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "maximum": 100,
            "minimum": 1,
            "type": "integer",
            "default": 10,
            "name": "pageSize",
            "in": "query"
          },
          {
            "enum": [
              "new",
              "authorized",
              "confirmed_not_synced",
              "confirmed",
              "canceling",
              "canceled",
              "refunded",
              "unknown"
            ],
            "type": "string",
            "name": "status",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "organizationId",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "washId",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "name": "postId",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/TransactionPage"
            }
          },
          "403": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/washes": {
      "get": {
        "tags": [
          "washes"
        ],
        "operationId": "getWashes",
        "parameters": [
          {
            "minimum": 0,
            "type": "integer",
            "format": "int64",
            "default": 0,
            "description": "Number of records to skip for pagination",
            "name": "offset",
            "in": "query"
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "default": 100,
            "description": "Maximum number of records to return",
            "name": "limit",
            "in": "query"
          },
          {
            "type": "string",
            "format": "uuid",
            "name": "groupId",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Wash"
              }
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "washes"
        ],
        "operationId": "createWash",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashCreation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success creation",
            "schema": {
              "$ref": "#/definitions/Wash"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/washes/{id}": {
      "get": {
        "tags": [
          "washes"
        ],
        "operationId": "getWashById",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Wash"
            }
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "washes"
        ],
        "operationId": "deleteWash",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "washes"
        ],
        "operationId": "updateWash",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashUpdate"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success update"
          },
          "default": {
            "description": "Generic error response",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Group": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Notification": {
      "type": "object",
      "properties": {
        "Amount": {
          "description": "Payment amount",
          "type": "integer"
        },
        "CardId": {
          "type": "integer",
          "x-nullable": true
        },
        "ErrorCode": {
          "description": "Error code",
          "type": "string"
        },
        "ExpDate": {
          "type": "string",
          "x-nullable": true
        },
        "OrderId": {
          "description": "Order ID",
          "type": "string"
        },
        "Pan": {
          "description": "PAN (Primary Account Number)",
          "type": "string"
        },
        "PaymentId": {
          "description": "Payment ID",
          "type": "integer"
        },
        "Status": {
          "description": "Payment status",
          "type": "string"
        },
        "Success": {
          "description": "Indicates whether the payment was successful",
          "type": "boolean"
        },
        "TerminalKey": {
          "description": "Terminal key",
          "type": "string"
        },
        "Token": {
          "description": "Payment token",
          "type": "string"
        }
      }
    },
    "Organization": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Payment": {
      "type": "object",
      "properties": {
        "amount": {
          "type": "integer"
        },
        "orderId": {
          "type": "string"
        },
        "postId": {
          "type": "string"
        },
        "washId": {
          "type": "string"
        }
      }
    },
    "PaymentCancellation": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "postID": {
          "type": "string"
        },
        "washID": {
          "type": "string"
        }
      }
    },
    "PaymentResponse": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "SimpleWash": {
      "type": "object",
      "required": [
        "id",
        "name",
        "deleted"
      ],
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "id",
        "createdAt",
        "amount",
        "status",
        "postId",
        "wash",
        "group",
        "organization"
      ],
      "properties": {
        "amount": {
          "type": "integer",
          "format": "int64"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "group": {
          "$ref": "#/definitions/Group"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "organization": {
          "$ref": "#/definitions/Organization"
        },
        "postId": {
          "type": "integer"
        },
        "status": {
          "$ref": "#/definitions/TransactionStatus"
        },
        "wash": {
          "$ref": "#/definitions/SimpleWash"
        }
      }
    },
    "TransactionPage": {
      "type": "object",
      "required": [
        "items",
        "page",
        "pageSize",
        "totalPages",
        "totalItems"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        },
        "page": {
          "type": "integer",
          "minimum": 1
        },
        "pageSize": {
          "type": "integer",
          "maximum": 100,
          "minimum": 1
        },
        "totalItems": {
          "type": "integer",
          "minimum": 0
        },
        "totalPages": {
          "type": "integer",
          "minimum": 0
        }
      }
    },
    "TransactionStatus": {
      "type": "string",
      "enum": [
        "new",
        "authorized",
        "confirmed_not_synced",
        "confirmed",
        "canceling",
        "canceled",
        "refunded",
        "unknown"
      ]
    },
    "User": {
      "description": "User profile",
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "organization": {
          "type": "object",
          "properties": {
            "deleted": {
              "type": "boolean"
            },
            "description": {
              "type": "string"
            },
            "displayName": {
              "type": "string"
            },
            "id": {
              "type": "string",
              "format": "uuid"
            },
            "name": {
              "type": "string"
            }
          },
          "x-nullable": true
        },
        "role": {
          "$ref": "#/definitions/UserRole"
        }
      }
    },
    "UserOrganization": {
      "type": "object",
      "properties": {
        "deleted": {
          "type": "boolean"
        },
        "description": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      },
      "x-nullable": true
    },
    "UserRole": {
      "type": "string",
      "enum": [
        "systemManager",
        "admin",
        "noAccess"
      ]
    },
    "Wash": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "groupId": {
          "type": "string",
          "format": "uuid"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "organizationId": {
          "type": "string",
          "format": "uuid"
        },
        "password": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    },
    "WashCreation": {
      "required": [
        "name",
        "description",
        "terminalKey",
        "terminalPassword",
        "groupId"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "groupId": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    },
    "WashUpdate": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminalKey": {
          "type": "string"
        },
        "terminalPassword": {
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "limit": {
      "minimum": 1,
      "type": "integer",
      "format": "int64",
      "default": 100,
      "description": "Maximum number of records to return",
      "name": "limit",
      "in": "query"
    },
    "offset": {
      "minimum": 0,
      "type": "integer",
      "format": "int64",
      "default": 0,
      "description": "Number of records to skip for pagination",
      "name": "offset",
      "in": "query"
    },
    "page": {
      "minimum": 1,
      "type": "integer",
      "default": 1,
      "name": "page",
      "in": "query"
    },
    "pageSize": {
      "maximum": 100,
      "minimum": 1,
      "type": "integer",
      "default": 10,
      "name": "pageSize",
      "in": "query"
    },
    "userRole": {
      "enum": [
        "systemManager",
        "admin",
        "noAccess"
      ],
      "type": "string",
      "name": "role",
      "in": "query"
    },
    "uuid": {
      "type": "string",
      "format": "uuid",
      "name": "id",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "GenericError": {
      "description": "Generic error response",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
}
