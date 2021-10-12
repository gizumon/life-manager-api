package responses

type PlainErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func ErrorPlain(title string, description string) PlainErrorResponse {
	return PlainErrorResponse{
		Error:            title,
		ErrorDescription: description,
	}
}
