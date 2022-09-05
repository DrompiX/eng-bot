package term

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/utils"
)

type testConfig struct {
	mockRepo *MockRepository
	serv     *service
	ctx      context.Context
}

func initTestConfig(t *testing.T, termLimit int) testConfig {
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepository(ctrl)
	s := NewService(mockRepo, termLimit)
	ctx := context.Background()
	return testConfig{mockRepo, s, ctx}
}

// Replace atomic call logic with simple function call
func patchAtomic(repo *MockRepository) {
	repo.EXPECT().Atomic(gomock.Any(), gomock.Any()).DoAndReturn(
		func (ctx context.Context, fn func(r Repository) error) error {
			return fn(repo)
		},
	)
}

func TestGetCollection(t *testing.T) {
	t.Run("get empty collection", func(t *testing.T) {
		tc := initTestConfig(t, 10)
		userId := user.NewUserID()
		emptyCollection := make([]*Term, 0)

		mockCall := tc.mockRepo.EXPECT().GetAllTerms(gomock.Any(), gomock.Eq(userId))
		mockCall.Times(1).Return(emptyCollection, nil)

		c, err := tc.serv.GetCollection(tc.ctx, userId)
		utils.Ok(t, err)
		utils.Equals(t, emptyCollection, c)
	})

	t.Run("get non-empty collection", func(t *testing.T) {
		tc := initTestConfig(t, 10)

		userId := user.NewUserID()
		collection := []*Term{
			{"user1", "hello", "привет"},
			{"user1", "kidney", "почка"},
			{"user1", "garden", "сад"},
			{"user1", "button", "кнопка"},
		}

		mockCall := tc.mockRepo.EXPECT().GetAllTerms(gomock.Any(), gomock.Eq(userId))
		mockCall.Times(1).Return(collection, nil)

		c, err := tc.serv.GetCollection(tc.ctx, userId)
		utils.Ok(t, err)
		utils.Equals(t, collection, c)
	})
}

func TestAddTerm(t *testing.T) {
	t.Run("term limit exceeded", func(t *testing.T) {
		tc := initTestConfig(t, 1)
		patchAtomic(tc.mockRepo)
		
		tm, _ := NewTerm("hello", "привет", "testuser")
		tc.mockRepo.EXPECT().GetTermCount(gomock.Any(), gomock.Eq(tm.Uid)).Times(1).Return(1, nil)
		tc.mockRepo.EXPECT().AddTerm(gomock.Any(), gomock.Any()).Times(0)

		err := tc.serv.AddTerm(tc.ctx, tm)
		utils.Equals(t, TermLimitExceeded, err)
	})

	t.Run("term already exist", func(t *testing.T) {
		tc := initTestConfig(t, 2)
		patchAtomic(tc.mockRepo)
		
		tm, _ := NewTerm("hello", "привет", "testuser")
		tc.mockRepo.EXPECT().GetTermCount(gomock.Any(), gomock.Eq(tm.Uid)).Times(1).Return(1, nil)
		tc.mockRepo.EXPECT().AddTerm(gomock.Any(), gomock.Any()).Times(1).Return(TermAlreadyExists)

		err := tc.serv.AddTerm(tc.ctx, tm)
		utils.Equals(t, TermAlreadyExists, err)
	})

	t.Run("add term successfully", func(t *testing.T) {
		tc := initTestConfig(t, 2)
		patchAtomic(tc.mockRepo)
		
		tm, _ := NewTerm("hello", "привет", "testuser")
		tc.mockRepo.EXPECT().GetTermCount(gomock.Any(), gomock.Eq(tm.Uid)).Times(1).Return(1, nil)
		tc.mockRepo.EXPECT().AddTerm(gomock.Any(), gomock.Any()).Times(1).Return(nil)

		err := tc.serv.AddTerm(tc.ctx, tm)
		utils.Ok(t, err)
	})
}
