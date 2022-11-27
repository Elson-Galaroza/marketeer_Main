package models

type GenericResponse struct {
	Code    int32 `json:"code"`
	Status  string
	Message string `json:"message,omitempty"`
	//Data    interface{} `json:",omitempty"`
}
