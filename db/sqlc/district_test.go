package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/irossa/educate/util"
	"github.com/stretchr/testify/require"
)

func createRandomDistrict(t *testing.T) District {
	argName := util.RandomName()

	district, err := testQueries.CreateDistrict(context.Background(), argName)
	require.NoError(t, err)
	require.NotEmpty(t, district)

	require.Equal(t, argName, district.Name)
	require.NotZero(t, district.ID)
	require.NotZero(t, district.CreatedAt)

	return district
}

func TestCreateDistrict(t *testing.T) {
	createRandomDistrict(t)
}

func TestGetDistrict(t *testing.T) {
	district1 := createRandomDistrict(t)
	district2, err := testQueries.GetDistrict(context.Background(), district1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, district2)

	require.Equal(t, district1.ID, district2.ID)
	require.Equal(t, district1.Name, district2.Name)
	require.WithinDuration(t, district1.CreatedAt, district2.CreatedAt, time.Second)
}

func TestUpdateDistrict(t *testing.T) {
	district1 := createRandomDistrict(t)

	arg := UpdateDistrictParams{
		ID:   district1.ID,
		Name: util.RandomName(),
	}

	district2, err := testQueries.UpdateDistrict(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, district2)

	require.Equal(t, district1.ID, district2.ID)
	require.NotEqual(t, district1.Name, district2.Name)
	require.WithinDuration(t, district1.CreatedAt, district2.CreatedAt, time.Second)

}

func TestDeleteDistrict(t *testing.T) {
	district1 := createRandomDistrict(t)
	err := testQueries.DeleteDistrict(context.Background(), district1.ID)
	require.NoError(t, err)

	district2, err := testQueries.GetDistrict(context.Background(), district1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, district2)
}

func TestGetAllDistricts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDistrict(t)
	}

	arg := GetAllDistrictsParams{
		Limit:  5,
		Offset: 5,
	}
	districts, err := testQueries.GetAllDistricts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, districts, 5)

	for _, district := range districts {
		require.NotEmpty(t, district)
	}
}
