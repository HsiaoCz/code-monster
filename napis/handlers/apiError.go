package handlers

type APIError struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (a APIError) Error() string {
	return a.Message
}
