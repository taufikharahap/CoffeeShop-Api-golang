package config

type Metas struct {
	Next  interface{} `json:"next"`
	Prev  interface{} `json:"prev"`
	Total int         `json:"total"`
}

type Result struct {
	Data    interface{}
	Meta    interface{}
	Message interface{}
}
