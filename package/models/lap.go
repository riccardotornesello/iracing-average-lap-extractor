package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Lap struct {
	gorm.Model

	EventSessionID uint
	EventSession   EventSession `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CustId int

	LapEvents pq.StringArray `gorm:"type:text[]"`
	Incident  bool
	LapTime   int
	LapNumber int
}
