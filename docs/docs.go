// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "Eder Velázquez",
            "url": "https://www.linkedin.com/in/oscar-eder-vel%C3%A1zquez-pineda/",
            "email": "edervzz.work@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/confirm/email/{token}": {
            "post": {
                "description": "Confirm user created",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Confirm sign up",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Confirmation token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/EmailConfirmUser"
                        }
                    },
                    "400": {
                        "description": "Invalid token"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login User Information",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Login"
                        }
                    },
                    "400": {
                        "description": "Login failed"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/login/refresh/{token}": {
            "post": {
                "description": "Refresh login session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Refresh Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Refresh token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/RefreshLogin"
                        }
                    },
                    "400": {
                        "description": "Login failed"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/me/password": {
            "put": {
                "description": "Change current password including a valid access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Change password",
                "parameters": [
                    {
                        "description": "Change password Information",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ChangePassword"
                        }
                    },
                    "400": {
                        "description": "Login failed"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/password/forgot": {
            "post": {
                "description": "Forgot password flow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Forgot password flow",
                "parameters": [
                    {
                        "description": "New Forgot Password info",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewPasswordForgot"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/PasswordForgot"
                        }
                    },
                    "400": {
                        "description": "Something is wrong"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/password/reset": {
            "post": {
                "description": "Forgot password flow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Forgot password flow",
                "parameters": [
                    {
                        "description": "Reset password information",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewResetPasswordForgot"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ResetPasswordForgot"
                        }
                    },
                    "400": {
                        "description": "Something is wrong"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/sign-up": {
            "post": {
                "description": "Register a new user with password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "description": "User Information",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/NewSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SignUp"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "409": {
                        "description": "User already exists."
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "ChangePassword": {
            "description": "Result of change current password",
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "expires_in": {
                    "type": "integer",
                    "example": 1800
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "token_type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "EmailConfirmUser": {
            "description": "User confirmed via email",
            "type": "object",
            "properties": {
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "Login": {
            "description": "Login",
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "expires_in": {
                    "type": "integer",
                    "example": 1800
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "token_type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "NewChangePassword": {
            "description": "Change current password",
            "type": "object",
            "required": [
                "new_password"
            ],
            "properties": {
                "current_password": {
                    "type": "string"
                },
                "new_password": {
                    "description": "\"One upper, one lower, one number, one special @#$%\u0026, len: 8-16\"",
                    "type": "string",
                    "format": "password",
                    "maxLength": 16,
                    "minLength": 8
                }
            }
        },
        "NewLogin": {
            "description": "Confirmation token",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@mail.com"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "NewPasswordForgot": {
            "description": "New Password forgot flow",
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "NewResetPasswordForgot": {
            "description": "Reset password forgot",
            "type": "object",
            "required": [
                "new_password"
            ],
            "properties": {
                "new_password": {
                    "description": "\"One upper, one lower, one number, one special @#$%\u0026, len: 8-16\"",
                    "type": "string",
                    "format": "password",
                    "maxLength": 16,
                    "minLength": 8
                },
                "reset_password_token": {
                    "type": "string",
                    "example": "2eba30ff-adb8-478b-913f-ace363acbd34"
                }
            }
        },
        "NewSignUp": {
            "description": "User account information",
            "type": "object",
            "required": [
                "email",
                "fullname",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "johndoe@mail.com"
                },
                "fullname": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 5,
                    "example": "John Doe"
                },
                "password": {
                    "description": "\"One upper, one lower, one number, one special @#$%\u0026, len: 8-16 \"",
                    "type": "string",
                    "format": "password",
                    "maxLength": 16,
                    "minLength": 8
                }
            }
        },
        "PasswordForgot": {
            "description": "Password forgot flow",
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "2eba30ff-adb8-478b-913f-ace363acbd34"
                }
            }
        },
        "RefreshLogin": {
            "description": "Refresh Login",
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "expires_in": {
                    "type": "integer",
                    "example": 1800
                },
                "refresh_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                },
                "token_type": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "ResetPasswordForgot": {
            "description": "Result Reset password forgot",
            "type": "object",
            "properties": {
                "isSuccess": {
                    "type": "boolean"
                }
            }
        },
        "SignUp": {
            "description": "Confirmation token",
            "type": "object",
            "properties": {
                "confirm_token": {
                    "type": "string",
                    "example": "2eba30ff-adb8-478b-913f-ace363acbd34"
                },
                "email": {
                    "type": "string",
                    "example": "johndoe@mail.com"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:6001",
	BasePath:         "/security/api/v1",
	Schemes:          []string{},
	Title:            "Security API",
	Description:      "Security API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
