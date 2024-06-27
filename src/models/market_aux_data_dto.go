package models

type MarketAuxDataDTO struct {
	Uuid        string              `json:"uuid"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Url         string              `json:"url"`
	PublishedAt string              `json:"published_at"`
	Source      string              `json:"source"`
	Entities    []MarketAuxEntities `json:"entities"`
}
