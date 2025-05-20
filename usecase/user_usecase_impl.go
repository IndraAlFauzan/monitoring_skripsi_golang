package usecase

import (
	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/user"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"github.com/indraalfauzan/monitoring_skripsi_golang/utils"
)

type userUseCaseImpl struct {
	userRepo domain.UserRepository
}

// RegisterUser implements domain.UserUseCase.
func (u *userUseCaseImpl) RegisterUser(user *entity.User) (*entity.User, error) {
	// valiadation input
	if user.Email == "" {
		return &entity.User{}, apperror.ValidationError("Nama Tidak Boleh Kosong")
	}

	if user.Username == "" {
		return &entity.User{}, apperror.ValidationError("Username Tidak Boleh Kosong")
	}

	if user.Password == "" {
		return &entity.User{}, apperror.ValidationError("Password Tidak Boleh Kosong")
	}

	if user.RoleID == 0 {
		return &entity.User{}, apperror.ValidationError("RoleID Tidak Boleh Kosong")
	}

	// check if user already exists with the same role
	existingUser, err := u.userRepo.GetUserByEmailandRole(user.Email, user.RoleID)
	if err != nil {
		return &entity.User{}, apperror.InternalServerErrorWithMessage("Failed to check user existence")
	}
	if existingUser != nil {
		return &entity.User{}, apperror.ValidationError("User already exists")
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return &entity.User{}, apperror.InternalServerErrorWithMessage("Failed to hash password")
	}

	user.Password = hash

	user, err = u.userRepo.CreateUser(user)
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil
}

// Login implements domain.UserUserCase.
func (u *userUseCaseImpl) Login(email string, password string) (*entity.User, error) {
	// valiadation input
	if email == "" {
		return &entity.User{}, apperror.ValidationError("Email Tidak Boleh Kosong")
	}
	if password == "" {
		return &entity.User{}, apperror.ValidationError("Password Tidak Boleh Kosong")
	}

	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return &entity.User{}, apperror.InternalServerErrorWithMessage("Failed to get user")
	}
	if user == nil {
		return &entity.User{}, apperror.ValidationError("User not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return &entity.User{}, apperror.ValidationError("Invalid password")
	}

	return user, nil
}

// Register implements domain.UserUserCase.
func (u *userUseCaseImpl) RegisterMhs(user *entity.User) (*entity.User, error) {
	// valiadation input
	if user.Email == "" {
		return &entity.User{}, apperror.ValidationError("Nama Tidak Boleh Kosong")
	}

	if user.Username == "" {
		return &entity.User{}, apperror.ValidationError("Username Tidak Boleh Kosong")
	}

	if user.Password == "" {
		return &entity.User{}, apperror.ValidationError("Password Tidak Boleh Kosong")
	}

	// check if user already exists
	existingUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return &entity.User{}, apperror.InternalServerErrorWithMessage("Failed to check user existence")
	}
	if existingUser != nil {
		return &entity.User{}, apperror.ValidationError("User already exists")
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return &entity.User{}, apperror.InternalServerErrorWithMessage("Failed to hash password")
	}

	user.Password = hash
	user.RoleID = 2 // Set default role ID for MHS
	user, err = u.userRepo.CreateUser(user)
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil

}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCaseImpl{
		userRepo: userRepo,
	}
}
