package dto

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
	// errors don't define JSON marshaling
}
