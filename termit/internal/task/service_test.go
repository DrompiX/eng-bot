package task

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/term"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
	"gitlab.ozon.dev/DrompiX/homework-2/termit/utils"
)

type testConfig struct {
	taskMockRepo *MockRepository
	termMockRepo *term.MockRepository
	serv         *Service
	ctx          context.Context
}

func initTestConfig(t *testing.T) testConfig {
	ctrl := gomock.NewController(t)
	taskMockRepo := NewMockRepository(ctrl)
	termMockRepo := term.NewMockRepository(ctrl)
	s := NewService(taskMockRepo, termMockRepo)
	ctx := context.TODO()
	return testConfig{taskMockRepo, termMockRepo, s, ctx}
}

func TestGenerateTask(t *testing.T) {
	t.Run("empty collection returns error", func(t *testing.T) {
		tc := initTestConfig(t)
		uid := user.UserID("testuser")

		getAllCall := tc.termMockRepo.EXPECT().GetAllTerms(gomock.Any(), gomock.Eq(uid)).Times(1)
		getAllCall.Return(nil, ErrCollectionIsEmpty)

		_, err := tc.serv.GenerateTask(tc.ctx, uid)
		utils.Equals(t, ErrCollectionIsEmpty, err)
	})

	t.Run("successful task generate", func(t *testing.T) {
		tc := initTestConfig(t)
		uid := user.UserID("testuser")
		terms := []*term.Term{
			{Uid: uid, Data: "hello", Translation: "привет"},
		}
		expT := NewTask(terms[0].Uid, terms[0].Data, terms[0].Translation)
		
		getAllCall := tc.termMockRepo.EXPECT().GetAllTerms(gomock.Any(), gomock.Eq(uid)).Times(1)
		getAllCall.Return(terms, nil)

		createCall := tc.taskMockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1)
		createCall.Return(nil)

		genT, err := tc.serv.GenerateTask(tc.ctx, uid)
		utils.Ok(t, err)
		utils.Equals(t, expT.Uid, genT.Uid)
		utils.Equals(t, expT.Term, genT.Term)
		utils.Equals(t, expT.Expected, genT.Expected)
	})
}

func TestCheckAnswer(t *testing.T) {
	t.Run("task already answered", func(t *testing.T) {
		tc := initTestConfig(t)
		uid := user.UserID("testuser")
		success := true
		task := NewTask(uid, "hello", "привет")
		task.Success = &success
		ans := NewAnswer(task.ID, uid, task.Expected)

		createCall := tc.taskMockRepo.EXPECT().GetById(gomock.Any(), gomock.Eq(task.ID)).Times(1)
		createCall.Return(*task, nil)

		_, err := tc.serv.CheckAnswer(tc.ctx, ans)
		utils.Equals(t, ErrAlreadyAnswered, err)
	})

	t.Run("answer is correct", func(t *testing.T) {
		tc := initTestConfig(t)
		uid := user.UserID("testuser")
		task := NewTask(uid, "hello", "привет")
		ans := NewAnswer(task.ID, uid, task.Expected)

		createCall := tc.taskMockRepo.EXPECT().GetById(gomock.Any(), gomock.Eq(task.ID)).Times(1)
		createCall.Return(*task, nil)

		updCall := tc.taskMockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1)
		updCall.Return(nil)

		exp := &AnswerCheck{Success: true, Expected: task.Expected}
		ac, err := tc.serv.CheckAnswer(tc.ctx, ans)
		utils.Ok(t, err)
		utils.Equals(t, exp, ac)
	})

	t.Run("answer is not correct", func(t *testing.T) {
		tc := initTestConfig(t)
		uid := user.UserID("testuser")
		task := NewTask(uid, "hello", "привет")
		ans := NewAnswer(task.ID, uid, "incorrect answer")

		createCall := tc.taskMockRepo.EXPECT().GetById(gomock.Any(), gomock.Eq(task.ID)).Times(1)
		createCall.Return(*task, nil)

		updCall := tc.taskMockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Times(1)
		updCall.Return(nil)

		exp := &AnswerCheck{Success: false, Expected: task.Expected}
		ac, err := tc.serv.CheckAnswer(tc.ctx, ans)
		utils.Ok(t, err)
		utils.Equals(t, exp, ac)
	})
}