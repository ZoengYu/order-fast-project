package db

import (
	"context"
	"database/sql"
	"testing"

	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashpassword, err := util.HashPassword("password")
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomUser(),
		HashedPassword: hashpassword,
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	get_user, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.Equal(t, user, get_user)
}

func TestDelAccount(t *testing.T) {
	user := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.Username)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), user.Username)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)
	updated_arg := UpdateUserParams{
		Username: user.Username,
		HashedPassword: sql.NullString{
			String: "hardcode_hashedPassword",
			Valid:  true,
		},
	}
	updated_user, err := testQueries.UpdateUser(context.Background(), updated_arg)
	require.NoError(t, err)
	require.Equal(t, updated_user.Username, updated_arg.Username)
	require.Equal(t, updated_user.HashedPassword, updated_arg.HashedPassword.String)
}
