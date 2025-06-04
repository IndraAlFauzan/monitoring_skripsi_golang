package usecase

import (
	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
)

type mahasiswaProfileUsecase struct {
	repo domain.MahasiswaProfileRepository
}

func NewMahasiswaProfileUsecase(repo domain.MahasiswaProfileRepository) domain.MahasiswaProfileUseCase {
	return &mahasiswaProfileUsecase{repo}
}

func (u *mahasiswaProfileUsecase) CreateProfile(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error) {

	// validate the profile data
	if profile.NIM == "" {
		return nil, apperror.ValidationError("NIM tidak boleh kosong")
	}

	if profile.Nama == "" {
		return nil, apperror.ValidationError("Nama tidak boleh kosong")
	}

	if profile.NoHP == "" {
		return nil, apperror.ValidationError("NoHP tidak boleh kosong")
	}

	if profile.PhotoPath == "" {
		return nil, apperror.ValidationError("PhotoPath tidak boleh kosong")
	}

	mhsProfile, err := u.repo.Create(profile)
	if err != nil {
		return nil, err
	}
	return mhsProfile, nil
}

func (u *mahasiswaProfileUsecase) GetProfile(userID int) (*entity.MahasiswaProfile, error) {
	mhsProfile, err := u.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	if mhsProfile == nil {
		return nil, apperror.ValidationError("Profile not found")
	}
	return mhsProfile, nil
}

func (u *mahasiswaProfileUsecase) UpdateProfile(profile *entity.MahasiswaProfile) (*entity.MahasiswaProfile, error) {
	// Validasi seperti Create
	if profile.NIM == "" {
		return nil, apperror.ValidationError("NIM Tidak Boleh Kosong")
	}
	if profile.Nama == "" {
		return nil, apperror.ValidationError("Nama  Tidak Boleh Kosong")
	}
	if profile.NoHP == "" {
		return nil, apperror.ValidationError("NoHP  Tidak Boleh Kosong")
	}
	if profile.PhotoPath == "" {
		return nil, apperror.ValidationError("PhotoPath  Tidak Boleh Kosong")
	}

	// Ambil data lama
	existing, err := u.repo.GetByUserID(profile.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, apperror.ValidationError("Profile not found")
	}

	// Update field-nya
	existing.NIM = profile.NIM
	existing.Nama = profile.Nama
	existing.NoHP = profile.NoHP
	existing.PhotoPath = profile.PhotoPath

	// Simpan ke repo
	return u.repo.Update(existing)
}
