package model

import (
	"time"

	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id             identifier.UserID
	externalUserID identifier.ExternalUserID
	name           string
	password       string
	mailAddress    string
	comments       string
	updatedAt      time.Time
	userDetail     UserDetail
}

type UserDetail struct {
	UserID               *identifier.UserID
	Name                 *string
	NameKana             *string
	Profession           string // 職業
	Occupation           string // 職種
	CountryOfCitizenship *uint8 // 国籍
	Age                  *uint8
	BirthDate            *string
	Gender               *uint8
	Height               uint16
	BirthPlace           string
	Residence            string // 居住地
	WorkLocation         string
	AnnualIncome         uint64
	TitleComment         string
	ProfileComment       string
}

func NewUser(
	externalUserID string,
	name string,
	password string,
	mailAddress string,
	comments string,
) (*User, error) {
	//password生成
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		externalUserID: identifier.ExternalUserID(externalUserID),
		name:           name,
		password:       string(pass),
		mailAddress:    mailAddress,
		comments:       comments,
	}

	return &user, nil
}

func ReConstructorUser(
	id identifier.UserID,
	externalUserID identifier.ExternalUserID,
	name string,
	password string,
	mailAddress string,
	comments string,
	updateAt time.Time,
	// userDetail UserDetail,
) *User {
	return &User{
		id:             id,
		externalUserID: externalUserID,
		name:           name,
		password:       password,
		mailAddress:    mailAddress,
		comments:       comments,
		updatedAt:      updateAt,
	}
}

func (u *User) ID() identifier.UserID                     { return u.id }
func (u *User) ExternalUserID() identifier.ExternalUserID { return u.externalUserID }
func (u *User) Name() string                              { return u.name }
func (u *User) Password() string                          { return u.password }
func (u *User) MailAddress() string                       { return u.mailAddress }
func (u *User) Comments() string                          { return u.comments }
func (u *User) UpdatedAt() time.Time                      { return u.updatedAt }

func (u *User) UserDetail() UserDetail { return u.userDetail }
