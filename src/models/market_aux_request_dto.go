package models

type MarketAuxRequestDTO struct {
	Meta   MarketAuxMeta    `json:"meta"`
	Data   []MarketAuxData   `json:"data"`
}