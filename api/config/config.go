package config

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/spf13/viper"
)

type config struct {
	MySQL   MySQL   `mapstructure:"mysql"`
	Server  Server  `mapstructure:"server"`
	Service Service `mapstructure:"service"`
	Redis   Redis   `mapstructure:"redis"`
}

var once sync.Once
var cfg config
var buildMode = os.Getenv("BUILD_MODE")

//go:embed default/production.yml
var defaultProduction []byte

//go:embed default/local.yml
var defaultLocal []byte

//go:embed default/test.yml
var defaultTest []byte

func getConfig() config {
	once.Do(func() {
		viper.SetConfigType("yml")
		// これでaws上の環境変数を取得しているっぽい
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		def := defaultLocal
		if isTesting() {
			def = defaultTest
		} else {
			switch buildMode {
			case "production":
				def = defaultProduction
			}
		}

		// 設定ファイルを読み込みます
		err := viper.ReadConfig(bytes.NewBuffer(def))
		if err != nil {
			fmt.Println("Failed to read yml file:", err)
			os.Exit(1) //プログラムを終了する関数
		}

		cfg = config{}
		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})
	return cfg

}

func isTesting() bool {
	return flag.Lookup("test.v") != nil || strings.HasSuffix(os.Args[0], ".test")
}

// OverrideValueForTest test時のmutex処理のための上書き設定
type OverrideValueForTest struct {
	MySQLDatabase string
}

func (v OverrideValueForTest) bind(c *config) {
	if v.MySQLDatabase != "" {
		c.MySQL.Database = v.MySQLDatabase
	}
}

var overrideValue OverrideValueForTest

func SetOverrideValueForTest(t *testing.T, v OverrideValueForTest) {
	overrideValue = v
	t.Cleanup(func() {
		overrideValue = OverrideValueForTest{}
	})
}
