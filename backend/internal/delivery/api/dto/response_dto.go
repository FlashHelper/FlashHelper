package dto

type SuccessResponse struct {
	Success bool `json:"success" example:"true" extensions:"x-order=1"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false" extensions:"x-order=1"`
	Error   string `json:"error" example:"unauthorized" extensions:"x-order=2"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Error:   message,
	}
}
