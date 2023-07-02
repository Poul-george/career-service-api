package testhelper

import (
	"context"
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/go-testfixtures/testfixtures/v3"
	"testing"
)

func LoadFixture(t *testing.T, option func(*testfixtures.Loader) error) {
	c := config.GetMySQLConfig()
	h := handler.NewHandlerForTestFixture(t, c)
	db, err := h.Writer(context.Background()).DB()
	if err != nil {
		t.Fatal(err)
	}
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.SkipResetSequences(),
		option,
	)
	if err != nil {
		t.Fatal(err)
	}
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
