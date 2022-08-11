package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockdb "github.com/irossa/educate/db/mock"
	db "github.com/irossa/educate/db/sqlc"
	"github.com/irossa/educate/util"
	"github.com/stretchr/testify/require"
)

func TestGetDistrictAPI(t *testing.T) {
	district := randomDistrict()
	user, _ := randomUser(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetDistrict(gomock.Any(), gomock.Eq(district.ID)).
		Times(1).
		Return(district, nil)

	server := NewTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/district?id=%d", district.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	addAuthorization(t, request, server.tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomDistrict() db.District {
	return db.District{
		ID:   util.RandomInt(1, 1000),
		Name: util.RandomName(),
	}
}
