package utils

type MyResponse struct {
	Success bool   `json:"success" xml:"success"`
	Message string `json:"message" xml:"message"`
	Data    any    `json:"data" xml:"data"`
}

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

func FailedValidation(errors any) map[string]any {
	return map[string]any{
		"message": "Invalid data",
		"errors":  errors,
	}
}
