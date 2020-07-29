# Go-Swagger基础

## 介绍swagger

+ 给接口添加注释
+ 根据注释自动生成json文件
+ 根据json文件自动生成文档
+ 通过文档进行接口调用测试

## 安装

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
cd ./go-swagger
go install ./cmd/swagger
```

## 在接口中添加注释

注释语法请参考[官方手册](https://bfanger.nl/swagger-explained/#swaggerObject)

```
// Package classification User API.
//
// The purpose of this service is to provide an application
// that is using plain go code to define an API
//
//      Host: localhost
//      Version: 0.0.1
//
// swagger:meta

package main

// swagger:parameters GetExampleRequest
type GetExampleRequest struct {
	// in: query
	Id    string `json:"id"`
}

// getexample api response
// swagger:response GetExampleResponse
type GetExampleResponse struct {
	// response body
	// in: body
	Body struct {
		// the code og response
		//
		// Required: true
		Code    int             `json:"code"`
		// the message of response
		//
		// Required: true
		Message string          `json:"message"`
		// response data
		Data    interface{}     `json:"data"`
	}
}

// swagger:route GET /getexample tag GetExampleRequest
//
// getexample route
//
// This is an getexample route
//
// Responses:
// 200: GetExampleResponse

// swagger:parameters PostExampleRequest
type PostExampleRequest struct {
	// in: body
	Body struct {
		Id    string `json:"id"`
		Name  string  `json:"name"`
	}
}

// postexample api response
// swagger:response PostExampleResponse
type PostExampleResponse struct {
	// response body
	// in: body
	Body struct {
		// the code og response
		//
		// Required: true
		Code    int             `json:"code"`
		// the message of response
		//
		// Required: true
		Message string          `json:"message"`
		// response data
		Data    interface{}     `json:"data"`
	}
}

// swagger:route POST /postexample tag PostExampleRequest
//
// postexample route
//
// This is an postexample route
//
// Responses:
// 200: PostExampleResponse
```
以上是一个具有标准化注释的接口文件，有两个例子接口，一个get，在request中带有参数，一个是post，提交json格式的数据。

## 生成标准swagger json

```
swagger generate spec -o ./swagger.json
```
该命令从项目的main.go开始扫描，解析所有的swagger注释，最后在项目的跟路径下生成一个swagger.json文件。

上面的注释会生成以下json
```
{
  "swagger": "2.0",
  "info": {},
  "host": "40.115.153.174:30335",
  "paths": {
    "/getexample": {
      "get": {
        "description": "This is an getexample route",
        "tags": [
          "tag"
        ],
        "summary": "getexample route",
        "operationId": "GetExampleRequest",
        "parameters": [
          {
            "type": "string",
            "x-go-name": "Id",
            "name": "id",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/GetExampleResponse"
          }
        }
      }
    },
    "/postexample": {
      "post": {
        "description": "This is an postexample route",
        "tags": [
          "tag"
        ],
        "summary": "postexample route",
        "operationId": "PostExampleRequest",
        "parameters": [
          {
            "name": "Body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "type": "string",
                  "x-go-name": "Id"
                },
                "name": {
                  "type": "string",
                  "x-go-name": "Name"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/PostExampleResponse"
          }
        }
      }
    }
  },
  "responses": {
    "GetExampleResponse": {
      "description": "getexample api response",
      "schema": {
        "type": "object",
        "required": [
          "code",
          "message"
        ],
        "properties": {
          "code": {
            "description": "the code og response",
            "type": "integer",
            "format": "int64",
            "x-go-name": "Code"
          },
          "data": {
            "description": "response data",
            "type": "object",
            "x-go-name": "Data"
          },
          "message": {
            "description": "the message of response",
            "type": "string",
            "x-go-name": "Message"
          }
        }
      }
    },
    "PostExampleResponse": {
      "description": "postexample api response",
      "schema": {
        "type": "object",
        "required": [
          "code",
          "message"
        ],
        "properties": {
          "code": {
            "description": "the code og response",
            "type": "integer",
            "format": "int64",
            "x-go-name": "Code"
          },
          "data": {
            "description": "response data",
            "type": "object",
            "x-go-name": "Data"
          },
          "message": {
            "description": "the message of response",
            "type": "string",
            "x-go-name": "Message"
          }
        }
      }
    }
  }
}
```

## 通过swagger json进行接口调用测试

### 使用chrome插件

+ 开启文件服务，使得通过url可以访问到swagger json文档
+ 下载chrome插件swagger ui console
+ 通过url访问swagger json文档，如http://40.115.153.174:30335/swagger/swagger.json
+ 开始进行接口调用测试

### 使用swaager-ui


## 遇到的问题

1. 默认情况下，调用接口的ip地址是和swagger json文档服务的ip地址一样的，但很多时候服务接口的ip地址和swagger json文档的ip地址不一样，因为往往不会部署在一台机器上。如何配置调用接口的ip地址？
在swagger json中增加host配置项即可，如"host": "40.115.153.174:30335"

