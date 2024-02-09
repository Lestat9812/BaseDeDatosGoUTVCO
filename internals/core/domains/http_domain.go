package domains

type ResultResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Errors struct {
	Message error `json:"message"`
}
