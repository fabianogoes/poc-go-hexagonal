package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response represents a response body format
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// errorResponse represents an error response body format
type errorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
}

// newResponse is a helper function to create a response body
func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(message string) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
	}
}

// handleSuccess sends a success response with the specified status code and optional data
func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleBadRequestError(ctx *gin.Context, err error) {
	errRsp := newErrorResponse(err.Error())

	ctx.JSON(http.StatusBadRequest, errRsp)
}
