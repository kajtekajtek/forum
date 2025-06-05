package models

import (
	"time"
)

type Server struct {
	ID			uint			`gorm:"primaryKey" json:"id"`
	Name		string			`gorm:"type:text;not null" json:"name"`
	CreatedAt	time.Time		`gorm:"autoCreateTime" json:"createdAt"`
	Members		[]Membership	`gorm:"foreignKey:ServerID" json:"-"`
}

type Membership struct {
	ID			uint		`gorm:"primaryKey" json:"id"`
	UserID		string		`gorm:"type:text;not null;index" json:"userId"`
	ServerID	uint		`gorm:"not null;index" json:"serverId"`
	Role		string		`gorm:"type:text;not null" json:"role"`
	CreatedAt	time.Time	`gorm:"autoCreateTime" json:"createdAt"`
	// relation
	Server		*Server		`gorm:"foreignKey:ServerID" json:"-"`
}
