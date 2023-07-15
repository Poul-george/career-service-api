package list_test

import (
	"encoding/json"
	"github.com/Poul-george/go-api/api/config"
	useCase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/repository/user"
	"github.com/Poul-george/go-api/api/presentation/appapi/controller/user/list"
	customContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	cmdConfig "github.com/Poul-george/go-api/api/presentation/cmd/config"
	"github.com/Poul-george/go-api/api/util/testhelper"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestController_Get(t *testing.T) {
	testhelper.Lock(t)
	testhelper.LoadFixture(t, testfixtures.Directory("testdata/fixtures"))

	h := handler.NewHandler(config.GetMySQLConfig())
	c := list.NewController(
		*useCase.NewUseCase(user.NewRepository(h)),
	)

	tests := []struct {
		name           string
		wantResponse   list.Response
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "削除されていないs、user情報が全権取得できること",
			wantResponse: list.Response{
				Items: []list.UserResponse{
					{
						ID:             1,
						ExternalUserID: "111",
						Name:           "user1",
						MailAddress:    "user@lcc.com",
						Comments:       "user comments 1",
						UpdatedAt:      time.Date(2023, 1, 1, 00, 00, 00, 00, time.Local),
					},
					{
						ID:             2,
						ExternalUserID: "222",
						Name:           "user2",
						MailAddress:    "user@lcc.com",
						Comments:       "user comments 2",
						UpdatedAt:      time.Date(2023, 1, 2, 00, 00, 00, 00, time.Local),
					},
					{
						ID:             4,
						ExternalUserID: "444",
						Name:           "user4",
						MailAddress:    "user@lcc.com",
						Comments:       "user comments 4",
						UpdatedAt:      time.Date(2023, 1, 4, 00, 00, 00, 00, time.Local),
					},
				},
			},
			wantStatusCode: 200,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := cmdConfig.NewEchoServer()
			req := httptest.NewRequest(http.MethodGet, "/dummy", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			customCtx := customContext.Context{Context: ctx}

			if err := c.Get(customCtx); (err != nil) != tt.wantErr {
				t.Errorf("Controller.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if rec.Code != tt.wantStatusCode {
				t.Errorf("Controller.Get() statusCode = %v, wantStatusCode = %v", rec.Code, tt.wantStatusCode)
				return
			}

			if tt.wantStatusCode == http.StatusOK {
				var gotResponse list.Response
				if err := json.Unmarshal(rec.Body.Bytes(), &gotResponse); err != nil {
					t.Error(err)
					return
				}
				// TODO: updateAtのレスポンスカラムがtimezoneの影響で00:00:00が09:00:00になるのを修正する
				if diff := cmp.Diff(gotResponse, tt.wantResponse, cmpopts.IgnoreFields(list.UserResponse{}, "UpdatedAt")); diff != "" {
					t.Errorf("Controller.Get() diff(-gotResponse +wantResponse)\n%s", diff)
				}
			}

		})
	}
}
