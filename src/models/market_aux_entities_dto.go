package models

type MarketAuxEntitiesDTO struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	MatchScore     float64 `json:"match_score"`
	SentimentScore float64 `json:"sentiment_score"`
}
