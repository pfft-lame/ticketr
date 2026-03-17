package apiresponse

type ApiResponse struct {
	Success    bool `json:"success"`
	StatusCode int  `json:"status_code"`
	Message    any  `json:"message,omitempty"`
	Errors     any  `json:"errors,omitempty"`
	Body       any  `json:"body,omitempty"`
}
