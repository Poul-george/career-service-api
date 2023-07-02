package testhelper

import (
	"fmt"
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"syscall"
	"testing"
	"time"
)

const maxConcurrency = 5
const mutexTimeout = 30 + time.Second

func Lock(t *testing.T) {
	i := 1
	timeout := time.Now().Add(mutexTimeout)
	for {
		if time.Now().After(timeout) {
			t.Fatal("mutex timeout")
		}

		f, err := lock(t, i)
		if err != nil {
			i = i%maxConcurrency + 1
			continue
		}

		t.Cleanup(func() {
			f(true)
		})
		return
	}
}

// テスト内で複数のテストが同時に実行されることを防止するための排他制御機能を提供
func lock(t *testing.T, i int) (func(bool), error) {
	// プロセスごとにロックファイルを生成する
	// ロックが取れなければエラーを返す
	fd, err := syscall.Open(lockfile(i), syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(fd, syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		if err := syscall.Close(fd); err != nil {
			return nil, err
		}
		return nil, err
	}
	overrideConfig(t, i)
	// ロック解除用の関数を返却する
	return func(closeDB bool) {
		reset(closeDB)
		if err := syscall.Flock(fd, syscall.LOCK_UN); err != nil {
			t.Fatal(err)
		}
		if err := syscall.Close(fd); err != nil {
			t.Fatal(err)
		}
	}, nil
}
func lockfile(i int) string {
	return fmt.Sprintf("/tmp/go-api-%d.lock", i)
}
func overrideConfig(t *testing.T, i int) {
	if i > 1 {
		config.SetOverrideValueForTest(
			t,
			config.OverrideValueForTest{
				MySQLDatabase: fmt.Sprintf("mysql_test_%d", i),
			},
		)
	} else {
		config.SetOverrideValueForTest(
			t,
			config.OverrideValueForTest{
				MySQLDatabase: "mysql_test",
			},
		)
	}
}
func reset(closeDB bool) {
	if closeDB {
		// db handlerはsingletonで保持されるのでテストごとにcloseする
		c := config.GetMySQLConfig()
		h := handler.NewHandler(c)
		h.Close()
	}
}
