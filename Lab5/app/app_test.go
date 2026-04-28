package app

import (
	"errors"
	"testing"

	"lab5/app/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	errDB := errors.New("database error")

	tests := []struct {
		Name        string
		InputName   string
		SetupMock   func(db *mocks.MockDB)
		ExpectedMsg []string
		ExpectedErr error
	}{
		{
			Name:      "Успешное добавление пользователя",
			InputName: "Alice",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil)
				db.EXPECT().Insert("Alice").Return(nil)
				db.EXPECT().GetAll().Return([]string{"Alice"}, nil)
			},
			ExpectedMsg: []string{"Alice"},
			ExpectedErr: nil,
		},
		{
			Name:      "Пустое имя - не добавляем",
			InputName: "",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil)
				db.EXPECT().GetAll().Return([]string{}, nil)
			},
			ExpectedMsg: []string{},
			ExpectedErr: nil,
		},
		{
			Name:      "Ошибка при CreateTable",
			InputName: "Test",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(errDB)
			},
			ExpectedMsg: nil,
			ExpectedErr: errDB,
		},
		{
			Name:      "Ошибка при Insert",
			InputName: "Test",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil)
				db.EXPECT().Insert("Test").Return(errDB)
			},
			ExpectedMsg: nil,
			ExpectedErr: errDB,
		},
		{
			Name:      "Ошибка при GetAll",
			InputName: "Test",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil)
				db.EXPECT().Insert("Test").Return(nil)
				db.EXPECT().GetAll().Return(nil, errDB)
			},
			ExpectedMsg: nil,
			ExpectedErr: errDB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDB := mocks.NewMockDB(ctrl)
			tt.SetupMock(mockDB)

			a := New(mockDB)
			users, err := a.Run(tt.InputName)

			require.ErrorIs(t, err, tt.ExpectedErr)
			require.Equal(t, tt.ExpectedMsg, users)
		})
	}
}

func TestClear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDB(ctrl)
	mockDB.EXPECT().DeleteAll().Return(nil)

	a := New(mockDB)
	err := a.Clear()

	require.NoError(t, err)
}
