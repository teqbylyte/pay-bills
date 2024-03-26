package utils

type MyResponse struct {
	Success bool   `json:"success" xml:"success"`
	Message string `json:"message" xml:"message"`
	Data    any    `json:"data" xml:"data"`
	Errors  any    `json:"errors"`
}

// SuccessResponse - The response structure for a successful request call
func SuccessResponse(message string, data ...any) MyResponse {
	var payload any

	if len(data) > 0 {
		payload = &data[0] // The first value in the list or null if empty
	}

	return MyResponse{
		Success: true,
		Message: message,
		Data:    payload,
	}
}

// FailedResponse - The failed response for a failed request call
func FailedResponse(message string, data ...any) MyResponse {
	var payload any

	if len(data) > 0 {
		payload = &data[0] // The first value in the list or null if empty
	}

	return MyResponse{
		Success: false,
		Message: message,
		Data:    payload,
	}
}

// FailedValidationResponse - The response for a request which fails on data validation. Should be return with a 422 status code.
func FailedValidationResponse(errors any) MyResponse {
	return MyResponse{
		Success: false,
		Message: "Invalid data",
		Errors:  errors,
	}
}
