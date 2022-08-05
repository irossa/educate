package db

import (
	"context"
	"testing"
	"time"

	"github.com/irossa/educate/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomName(),
		HashedPassword: "secret",
		FullName:       util.RandomName(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:             user1.ID,
		Username:       util.RandomName(),
		HashedPassword: util.RandomName(),
		FullName:       util.RandomName(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.NotEqual(t, user1.Username, user2.Username)
	require.NotEqual(t, user1.HashedPassword, user2.HashedPassword)
	require.NotEqual(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)

}

//func TestDeleteUser(t *testing.T) {
//	district1 := createRandomDistrict(t)
//	err := testQueries.DeleteDistrict(context.Background(), district1.ID)
//	require.NoError(t, err)

//	district2, err := testQueries.GetDistrict(context.Background(), district1.ID)
//	require.Error(t, err)
//	require.EqualError(t, err, sql.ErrNoRows.Error())
//	require.Empty(t, district2)
//}

//func TestGetAllUsers(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		createRandomDistrict(t)
//	}

//	arg := GetAllDistrictsParams{
//		Limit:  5,
//		Offset: 5,
//	}
//	districts, err := testQueries.GetAllDistricts(context.Background(), arg)
//	require.NoError(t, err)
//	require.Len(t, districts, 5)

//	for _, district := range districts {
//		require.NotEmpty(t, district)
//	}
//}
