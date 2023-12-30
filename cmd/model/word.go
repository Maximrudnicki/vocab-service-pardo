package model

type Word struct {
	Id         uint32 `gorm:"type:int;primary_key"`
	Word       string `gorm:"type:varchar;not null"`
	Definition string `gorm:"type:varchar;not null"`
	UserId     uint32 `gorm:"not null"`
}
