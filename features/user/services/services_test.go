package services

import (
	"errors"
	"testing"

	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/config"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/domain"
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUserr(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Username: "same", Email: "skjsa"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.AddUser(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID, "Seharusnya ada ID yang berhasil dibuat")
		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
		assert.Equal(t, input.Username, res.Username, "Nama user harus sesuai")
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.AddUser(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DUPLICATED_DATA, "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
	t.Run("Gaga; Hash", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("cannot encript password")).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: ""}
		res, err := srv.AddUser(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "cannot encript password", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{ID: uint(1), Username: "same", Email: "skjsa"}, nil).Once()
		srv := New(repo)
		input := domain.Core{ID: 1, Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.UpdateProfile(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID, "Seharusnya ada ID user yang diupdate")
		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
		assert.Equal(t, input.Username, res.Username, "Nama user harus sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Update", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New("rejected from database")).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.UpdateProfile(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "rejected from database", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}
