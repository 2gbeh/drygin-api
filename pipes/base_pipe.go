package pipes

type BaseResponse struct {
	Success bool     	 	`json:"success"`
	Data   	any `json:"data,omitempty"`
	Message any	`json:"message,omitempty"`
}