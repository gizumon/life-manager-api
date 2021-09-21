package responses

type ResultData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  string       `json:"status"`
	Results []ResultData `array:"results"`
}

func SuccessEvents(data []ResultData) *SuccessResponse {
	return &SuccessResponse{Status: "OK", Results: data}
}
