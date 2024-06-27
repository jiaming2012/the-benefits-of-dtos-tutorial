package models

type MarketAuxMeta struct {
	Found int `json:"found"`
	Returned int `json:"returned"`
	Limit int `json:"limit"`
	Page int `json:"page"`
}
