package table

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint64
	ExternalUserID string
	Name           string
	Password       string
	MailAddress    string
	Comments       string
	UserDetail     UserDetail
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type UserDetail struct {
	UserID               uint64
	Name                 string
	NameKana             string
	Profession           string // 職業
	Occupation           string // 職種
	CountryOfCitizenship uint8  // 国籍
	Age                  uint8
	BirthDate            string
	Gender               uint8
	Height               uint16
	BirthPlace           string
	Residence            string // 居住地
	WorkLocation         string
	AnnualIncome         uint64
	TitleComment         string
	ProfileComment       string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
}
