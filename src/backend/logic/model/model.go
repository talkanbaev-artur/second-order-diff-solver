package model

type Numericals struct {
	MethodID          int     `json:"id"`
	MethodTitle       string  `json:"title"`
	MethodDescription string  `json:"description"`
	MethodIcon        *string `json:"icon"`
}
