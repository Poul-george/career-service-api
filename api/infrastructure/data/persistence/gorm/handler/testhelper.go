package handler

import (
	"fmt"
	"github.com/Poul-george/go-api/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func NewHandlerForTestFixture(t *testing.T, c config.MySQL) *Handler {
	t.Helper()
	res, err := gorm.Open(mysql.Open(getDSNForTestFixture(t, c)))
	if err != nil {
		panic(err)
	}
	return &Handler{handler: res}
}

func getDSNForTestFixture(t *testing.T, c config.MySQL) string {
	t.Helper()
	return fmt.Sprintf(
		// "%s:%s@tcp(%s:%d)/%s?charset%s&parseTime=True&loc=Asia/Tokyo",
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		// c.Encoding,
	)
}
