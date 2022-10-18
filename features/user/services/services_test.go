package services

import (
	"testing"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
func TestShowAllUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get All User", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{{ID: uint(1), Nama: "same", HP: "0822"}}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 1, "Data kembalian harus lebih dari 0")
		assert.Equal(t, uint(1), res[0].ID, "Data seharusnya berisi 1")
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All User", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.Core{}, nil).Once()
		srv := New(repo)
		res, err := srv.ShowAllUser()
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, "no data", "Pesan error kurang sesuai")
		repo.AssertExpectations(t)
	})
}
*/

func TestLoginUser(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Success Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Username: "aziz123", Password: "123456789"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Username: "aziz123", Password: "123456789"}
		res, err := srv.Login(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.Username, "username should have value")
		// assert.Equal(t, res.Username, input.Username, "username should have same value")
		assert.Equal(t, res.Password, input.Password, "password should have same value")
		repo.AssertExpectations(t)

	})
	t.Run("Failed Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(domain.Core{Username: "aziz123", Password: "123456789"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Username: "aziz123", Password: "123456789"}
		res, err := srv.Login(input)
		assert.NotNil(t, err)
		assert.NotEqual(t, err, "error")
		assert.Equal(t, res.Username, "failed")
		repo.AssertExpectations(t)

	})
}
