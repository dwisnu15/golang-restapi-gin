package models

//used as a returning response with data
type ResponseBody struct {
	Success bool        `json:"success"`
	Message string 		`json:"message"`
	Data    interface{} `json:"data"`
}

//normal response with no data
type ResponseMessage struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
}
