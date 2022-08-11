package api

import (
	"testing"

	db "github.com/irossa/educate/db/sqlc"
	"github.com/irossa/educate/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		ID:				util.RandomInt(1, 100000),
		Username:       util.RandomName(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomName(),
	}
	return
}
