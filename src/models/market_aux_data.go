package models

import (
	"github.com/google/uuid"
)

type MarketAuxData struct {
	Uuid        uuid.UUID
	Title       string
	Description string
	Url         string
	PublishedAt string
	Source      string
	Entities    []MarketAuxEntities
}
