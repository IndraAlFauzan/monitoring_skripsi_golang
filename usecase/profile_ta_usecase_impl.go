package usecase

import (
	"time"

	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"

	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
)

type profilTAUsecase struct {
	repo domain.ProfileTARepository
}

func NewProfilTAUseCase(r domain.ProfileTARepository) domain.ProfileTAUseCase {
	return &profilTAUsecase{r}
}

// GetProfileTA implements domain.ProfileTAUseCase.
func (u *profilTAUsecase) GetProfileTA(userID int) (*entity.ProfileTAWithPembimbing, error) {
	profile, err := u.repo.GetProfileTA(userID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (u *profilTAUsecase) AjukanTA(userID int, judul string, dosen1ID int, dosen2ID int) (*entity.ProfielTA, error) {
	if judul == "" {
		return nil, apperror.ValidationError("Judul Skripsi Tidak Boleh Kosong")
	}
	if dosen1ID == 0 {
		return nil, apperror.ValidationError("Pembimbing Satu Tidak Boleh Kosong")
	}
	if dosen2ID == 0 {
		return nil, apperror.ValidationError("Pembimbing Dua Tidak Boleh Kosong")
	}
	if dosen1ID == dosen2ID {
		return nil, apperror.ValidationError("Pembimbing tidak boleh sama")
	}

	profil := &entity.ProfielTA{
		UserID:          userID,
		JudulSkripsi:    judul,
		StatusBimbingan: "diajukan",
		CreatedAt:       time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
	}

	createdProfil, err := u.repo.CreateProfilTA(profil)
	if err != nil {
		return nil, err
	}

	// Insert pembimbing
	pembimbings := []*entity.ProfileTADosen{
		{ProfilTAID: createdProfil.IDProfile, DosenID: dosen1ID, StatusDosenID: 1},
		{ProfilTAID: createdProfil.IDProfile, DosenID: dosen2ID, StatusDosenID: 2},
	}

	if err := u.repo.CreatePembimbing(pembimbings); err != nil {
		return nil, err
	}

	return createdProfil, nil
}
