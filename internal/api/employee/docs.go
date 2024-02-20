package employee

import "template-gin-api/internal/response"

// swagger:route PATCH /employees/products employee updateQuantityProductRequest
// Edit Employee Product
//
// ---
// responses:
//
// 200: updateQuantityProductResponse
// 400: updateQuantityProductErrResponse
// 500: updateQuantityProductErrResponse

// swagger:parameters updateQuantityProductRequest
type _ struct {
	// in:body
	Body UpdateQuantityProductRequest
}

// type swaggerUpdateQuantityProductParamsWrapper struct {
// 	// in:body
// 	Body UpdateQuantityProductRequest
// }

// swagger:response updateQuantityProductResponse
type _ struct{}

// type swaggerUpdateQuantityProductResponseWrapper struct{}

// swagger:response updateQuantityProductErrResponse
type _ struct {
	// in:body
	Body response.ErrResponse
}

// type swaggerEmployeeBadRequestResponseWrapper struct {
// 	// in:body
// 	Body response.ErrResponse
// }

// swagger:route GET /employees/{id} employee inquiryEmployeeByIdRequest
// Get Employee by id
//
// ---
// responses:
//
// 200: inquiryEmployeeByIdResponse

// swagger:parameters inquiryEmployeeByIdRequest
type _ struct {
	// Id of a Employee
	//
	// in:path
	// example: 1
	Id int `json:"id"`
}

// swagger:response inquiryEmployeeByIdResponse
type _ struct {
	// in:body
	Body struct {
		// example: 200
		Code uint64 `json:"code"`
		// example: Success.
		Message string                  `json:"message"`
		Data    InquiryEmployeeResponse `json:"data"`
	}
}

// swagger:route POST /employees employee upsertEmployeeRequest
// Insert & Update Employee
//
// ---
// responses:
//
// 200: upsertEmployeeResponse
// 400: upsertEmployeeErrResponse

// swagger:parameters upsertEmployeeRequest
type _ struct {
	// in:body
	Body struct {
		// example: 1
		Id int `json:"id"`
		// example: username
		Username string `json:"username"`
		// example: email@email.com
		Email string `json:"email"`
	}
}

// swagger:response upsertEmployeeResponse
type _ struct{}

// swagger:response upsertEmployeeErrResponse
type _ struct {
	// in:body
	Body struct {
		// example: 400
		Code uint64 `json:"code"`
		// example: Invalid request.
		Message string `json:"message"`
		Error   string `json:"error"`
	}
}
