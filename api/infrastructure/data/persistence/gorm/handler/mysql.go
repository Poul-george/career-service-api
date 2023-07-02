package handler

import (
	"context"
	"fmt"
	"github.com/Poul-george/go-api/api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"sync"
)

type Handler struct {
	handler *gorm.DB
}

var once sync.Once
var handler *Handler

func (h *Handler) Reader(ctx context.Context) *gorm.DB {
	return h.handler.Clauses(dbresolver.Read).WithContext(ctx)
}

func (h *Handler) Writer(ctx context.Context) *gorm.DB {
	return h.handler.WithContext(ctx)
}

func NewHandler(c config.MySQL) *Handler {
	once.Do(func() {
		res, err := gorm.Open(mysql.Open(GetDSN(c)), &gorm.Config{})
		if err != nil {
			// panic(err.Error())
			panic(fmt.Errorf("failed to initialize database, got error %s, dsn=%s", err.Error(), GetDSN(c)))
		}
		handler = &Handler{handler: res}
	})
	return handler
}

func GetDSN(c config.MySQL) string {
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

func (h *Handler) Close() {
	db, err := handler.handler.DB()
	if err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}
	once = sync.Once{}
}
