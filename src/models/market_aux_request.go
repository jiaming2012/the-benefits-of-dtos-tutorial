package models

type MarketAuxRequest struct {
	Meta   MarketAuxMeta    `json:"meta"`
	Data   []MarketAuxData   `json:"data"`
}