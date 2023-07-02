package user

import (
	"context"
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
	"github.com/Poul-george/go-api/api/util/testhelper"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestRepository_FindByID(t *testing.T) {
	testhelper.Lock(t)
	testhelper.LoadFixture(t, testfixtures.Directory("testdata/fixtures/findbyid"))
	ctx := context.Background()
	r := NewRepository(handler.NewHandler(config.GetMySQLConfig()))

	//testLoc, _ := time.LoadLocation("Asia/Tokyo")

	type UserInfo struct {
		externalUserID identifier.ExternalUserID
		userID         identifier.UserID
	}
	tests := []struct {
		name    string
		args    UserInfo
		want    *model.User
		wantErr bool
	}{
		{
			name: "(正)externalUserIDだけを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("111"),
				userID:         identifier.UserID(0),
			},
			want: model.ReConstructorUser(
				1,
				"111",
				"test1",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(正)UserIDだけを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID(""),
				userID:         identifier.UserID(2),
			},
			want: model.ReConstructorUser(
				2,
				"222",
				"test2",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(正)externalUserID&UserIDを指定してuser情報が取得できる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("333"),
				userID:         identifier.UserID(3),
			},
			want: model.ReConstructorUser(
				3,
				"333",
				"test3",
				"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
				"test@gmail.com",
				"test comments",
				time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
			),
			wantErr: false,
		},
		{
			name: "(異)externalUserID&存在しないUserIDを指定してエラーになる",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("333"),
				userID:         identifier.UserID(100),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "(異)削除済みのユーザーを取得できない",
			args: UserInfo{
				externalUserID: identifier.ExternalUserID("444"),
				userID:         identifier.UserID(4),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FindByID(ctx, tt.args.externalUserID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestUserRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opt := cmp.AllowUnexported(model.User{})
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "updatedAt"), opt); diff != "" {
				t.Errorf("TestUserRepository.FindByID() diff(-got +want)\n%s", diff)
			}
		})
	}
}

func TestRepository_FindByIDs(t *testing.T) {
	testhelper.Lock(t)
	testhelper.LoadFixture(t, testfixtures.Directory("testdata/fixtures/findbyids"))
	ctx := context.Background()
	r := NewRepository(handler.NewHandler(config.GetMySQLConfig()))

	var tests = []struct {
		name    string
		want    []model.User
		wantErr bool
	}{
		{
			name: "userの情報が複数取得できる",
			want: []model.User{
				*model.ReConstructorUser(
					1,
					"111",
					"test1",
					"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
					"test@gmail.com",
					"test comments",
					time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
				),
				*model.ReConstructorUser(
					2,
					"222",
					"test2",
					"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
					"test@gmail.com",
					"test comments",
					time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
				),
				*model.ReConstructorUser(
					3,
					"333",
					"test3",
					"$2a$10$tZN5qGGheum3BL9up8VhbOXpojUnlyb5vQEehb.rkPqV8VeP57aHu",
					"test@gmail.com",
					"test comments",
					time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FindByIDs(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestUserRepository.FindByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(model.User{})
			if diff := cmp.Diff(got, tt.want, cmpopts.IgnoreFields(model.User{}, "updatedAt"), opt); diff != "" {
				t.Errorf("TestUserRepository.FindByIDs() diff(-got +want)\n%s", diff)
			}
		})
	}
}

func TestRepository_Create(t *testing.T) {
	testhelper.Lock(t)
	ctx := context.Background()
	r := NewRepository(handler.NewHandler(config.GetMySQLConfig()))

	user1, _ := model.NewUser(
		"1234",
		"test_user",
		"test_pass",
		"test@gmail.com",
		"test comments",
	)

	tests := []struct {
		name    string
		args    *model.User
		want    table.User
		wantErr bool
	}{
		{
			name: "正常に登録ができる",
			args: user1,
			want: table.User{
				ID:             1,
				ExternalUserID: "1234",
				Name:           "test_user",
				MailAddress:    "test@gmail.com",
				Comments:       "test comments",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// テーブル情報を削除
		_ = r.clearCreateTable(ctx)
		t.Run(tt.name, func(t *testing.T) {
			if err := r.Create(ctx, tt.args); (err != nil) != tt.wantErr {
				t.Errorf("TestUserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		var user table.User
		r.handler.Writer(ctx).Where("id = 1").First(&user)
		if diff := cmp.Diff(user, tt.want, cmpopts.IgnoreFields(table.User{}, "Password", "CreatedAt", "UpdatedAt", "DeletedAt")); diff != "" {
			t.Errorf("TestUserRepository.FindByIDs() diff(-got +want)\n%s", diff)
		}
	}
}

func (r *Repository) clearCreateTable(ctx context.Context) error {
	err := r.handler.Writer(ctx).Transaction(func(tx *gorm.DB) error {
		// テストデータの物理削除
		// coordinatesテーブルはcreated_at、updated_atにデフォルトで0000-00-00 00:00:00といった不正な値を許容する
		// SQL_MODEを変更してから削除する
		if err := tx.Exec("SET SQL_MODE='ALLOW_INVALID_DATES'").Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM users;").Error; err != nil {
			return err
		}
		if err := tx.Exec("SET SQL_MODE=''").Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Where("1 = 1").Delete(&table.UserDetail{}).Error; err != nil {
			return err
		}
		// AUTO_INCREMENTのリセット
		if err := tx.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
