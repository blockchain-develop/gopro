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
