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
    "description": "microservice for the sbp pay system of self-service car washes",
    "title": "wash-sbp",
    "version": "1.0.0"
  },
  "paths": {
    "/cancel": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "cancel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cancel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health_check": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
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
          "wash"
        ],
        "operationId": "notification",
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
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pay": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "pay",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pay"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/payResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "security": [
          {}
        ],
        "tags": [
          "wash"
        ],
        "operationId": "signup",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/FirebaseToken"
            }
          }
        }
      }
    },
    "/wash/": {
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashCreate"
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashDelete"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "update",
        "parameters": [
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/wash/list": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "list",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pagination"
            }
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/wash/{id}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "getWash",
        "parameters": [
          {
            "type": "string",
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FirebaseToken": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "Notification": {
      "type": "object",
      "properties": {
        "AccountToken": {
          "description": "Идентификатор привязки счета, назначаемый банком-эмитентом",
          "type": "string"
        },
        "Amount": {
          "type": "integer"
        },
        "BankMemberId": {
          "description": "Идентификатор банка-эмитента клиента, который будет совершать оплату по привязанному счету",
          "type": "string"
        },
        "BankMemberName": {
          "description": "Наименование банка-эмитента",
          "type": "string"
        },
        "CardId": {
          "type": "integer"
        },
        "ErrorCode": {
          "type": "string"
        },
        "ExpDate": {
          "type": "string"
        },
        "Message": {
          "description": "Краткое описание ошибки",
          "type": "string"
        },
        "NotificationType": {
          "description": "Код ошибки (\u003c= 20 символов)",
          "type": "string"
        },
        "OrderID": {
          "type": "string"
        },
        "Pan": {
          "type": "string"
        },
        "PaymentID": {
          "type": "string"
        },
        "RequestKey": {
          "description": "Идентификатор запроса на привязку счета",
          "type": "string"
        },
        "Status": {
          "type": "string"
        },
        "Success": {
          "type": "boolean"
        },
        "TerminalKey": {
          "type": "string"
        },
        "Token": {
          "description": "Подпись запроса",
          "type": "string"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "integer",
          "format": "int64",
          "maximum": 100
        },
        "offset": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Pay": {
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
    "Wash": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "WashCreate": {
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "WashDelete": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
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
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "cancel": {
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
    "error": {
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
    "payResponse": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
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
    "description": "microservice for the sbp pay system of self-service car washes",
    "title": "wash-sbp",
    "version": "1.0.0"
  },
  "paths": {
    "/cancel": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "cancel",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cancel"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/health_check": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
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
          "wash"
        ],
        "operationId": "notification",
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
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pay": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "pay",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pay"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/payResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/signup": {
      "post": {
        "security": [
          {}
        ],
        "tags": [
          "wash"
        ],
        "operationId": "signup",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/FirebaseToken"
            }
          }
        }
      }
    },
    "/wash/": {
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashCreate"
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashDelete"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "update",
        "parameters": [
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/wash/list": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "list",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Pagination"
            }
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/wash/{id}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash"
        ],
        "operationId": "getWash",
        "parameters": [
          {
            "type": "string",
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
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "403": {
            "description": "Access denied",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Wash not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "FirebaseToken": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "Notification": {
      "type": "object",
      "properties": {
        "AccountToken": {
          "description": "Идентификатор привязки счета, назначаемый банком-эмитентом",
          "type": "string"
        },
        "Amount": {
          "type": "integer"
        },
        "BankMemberId": {
          "description": "Идентификатор банка-эмитента клиента, который будет совершать оплату по привязанному счету",
          "type": "string"
        },
        "BankMemberName": {
          "description": "Наименование банка-эмитента",
          "type": "string"
        },
        "CardId": {
          "type": "integer"
        },
        "ErrorCode": {
          "type": "string"
        },
        "ExpDate": {
          "type": "string"
        },
        "Message": {
          "description": "Краткое описание ошибки",
          "type": "string"
        },
        "NotificationType": {
          "description": "Код ошибки (\u003c= 20 символов)",
          "type": "string"
        },
        "OrderID": {
          "type": "string"
        },
        "Pan": {
          "type": "string"
        },
        "PaymentID": {
          "type": "string"
        },
        "RequestKey": {
          "description": "Идентификатор запроса на привязку счета",
          "type": "string"
        },
        "Status": {
          "type": "string"
        },
        "Success": {
          "type": "boolean"
        },
        "TerminalKey": {
          "type": "string"
        },
        "Token": {
          "description": "Подпись запроса",
          "type": "string"
        }
      }
    },
    "Pagination": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "integer",
          "format": "int64",
          "maximum": 100
        },
        "offset": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "Pay": {
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
    "Wash": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "WashCreate": {
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "WashDelete": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
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
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "terminal_key": {
          "type": "string"
        },
        "terminal_password": {
          "type": "string"
        }
      }
    },
    "cancel": {
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
    "error": {
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
    "payResponse": {
      "type": "object",
      "properties": {
        "orderID": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
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
