package user

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/utils"
)

type testConfig struct {
	mockRepo *MockRepository
	serv     *Service
	ctx      context.Context
}

func initTestConfig(t *testing.T) testConfig {
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)
	s := NewService(mockRepo)
	ctx := context.TODO()
	return testConfig{mockRepo, s, ctx}
}

var (
	pass = "password"
	passSHA512 = "b109f3bbbc244eb82441917ed06d618b9008dd09b3befd1b5e07394c706a8bb980b1d7785e5976ec049b46df5f1326af5a2ea6d103fd07c95385ffab0cacbc86"
)

func TestAuthenticate(t *testing.T) {
	t.Run("user not found", func(t *testing.T) {
		tc := initTestConfig(t)
		u := NewUser("testuser", "testpassword")

		call := tc.mockRepo.EXPECT().Find(gomock.Any(), gomock.Eq(u.Username)).Times(1)
		call.Return(nil, ErrNotFound)

		_, err := tc.serv.Authenticate(tc.ctx, u)
		utils.Equals(t, ErrNotFound, err)
	})

	t.Run("passwords don't match", func(t *testing.T) {
		tc := initTestConfig(t)
		auth := NewUser("testuser", "somehashedpassword")
		dbUser := NewUser("testuser", "anotherhashedpassword")

		findCall := tc.mockRepo.EXPECT().Find(gomock.Any(), gomock.Eq(auth.Username)).Times(1)
		findCall.Return(dbUser, nil)

		_, err := tc.serv.Authenticate(tc.ctx, auth)
		utils.Equals(t, ErrIncorrectPassword, err)
	})

	t.Run("successfully authenticated", func(t *testing.T) {
		tc := initTestConfig(t)
		auth := NewUser("testuser", pass)
		dbUser := NewUser("testuser", passSHA512)

		findCall := tc.mockRepo.EXPECT().Find(gomock.Any(), gomock.Eq(auth.Username)).Times(1)
		findCall.Return(dbUser, nil)

		u, err := tc.serv.Authenticate(tc.ctx, auth)
		utils.Ok(t, err)
		utils.Equals(t, dbUser, u)
	})
}

func TestAddUser(t *testing.T) {
	t.Run("user already exists", func(t *testing.T) {
		tc := initTestConfig(t)
		u := NewUser("testuser", "testpassword")

		call := tc.mockRepo.EXPECT().Save(gomock.Any(), gomock.Eq(u)).Times(1)
		call.Return(ErrAlreadyExists)

		err := tc.serv.AddUser(tc.ctx, u)
		utils.Equals(t, ErrAlreadyExists, err)
	})

	t.Run("successful user save", func(t *testing.T) {
		tc := initTestConfig(t)
		u := NewUser("testuser", pass)

		call := tc.mockRepo.EXPECT().Save(gomock.Any(), gomock.Eq(u)).Times(1)
		call.Return(nil)

		err := tc.serv.AddUser(tc.ctx, u)
		utils.Ok(t, err)
		utils.Equals(t, u.Password, passSHA512)
	})
}
