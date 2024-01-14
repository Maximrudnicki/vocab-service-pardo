package model

import "time"

type Word struct {
	Id         uint32    `gorm:"type:int;primary_key"`
	Word       string    `gorm:"type:varchar;not null"`
	Definition string    `gorm:"type:varchar;not null"`
	UserId     uint32    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:now()"`

	IsLearned       bool `gorm:"default:false"`
	Cards           bool `gorm:"default:false"`
	WordTranslation bool `gorm:"default:false"`
	Constructor     bool `gorm:"default:false"`
	WordAudio       bool `gorm:"default:false"`
}
