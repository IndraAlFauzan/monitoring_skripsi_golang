package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"github.com/indraalfauzan/monitoring_skripsi_golang/response"
)

type MahasiswaHandler struct {
	usecase domain.MahasiswaProfileUseCase
}

func NewMahasiswaHandler(uc domain.MahasiswaProfileUseCase) *MahasiswaHandler {
	return &MahasiswaHandler{uc}
}

func (h *MahasiswaHandler) CreateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	nim := c.PostForm("nim")
	nama := c.PostForm("nama")
	noHP := c.PostForm("no_hp")

	// Handle file
	file, err := c.FormFile("photo")
	if err != nil {
		response.WriteJSONResponse(c, http.StatusBadRequest, "Photo is required", nil)
		return
	}

	// Simpan file
	path := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		response.WriteJSONResponse(c, http.StatusInternalServerError, "Failed to upload file", nil)
		return
	}

	profile := &entity.MahasiswaProfile{
		UserID:    userID,
		NIM:       nim,
		Nama:      nama,
		NoHP:      noHP,
		PhotoPath: path,
	}

	result, err := h.usecase.CreateProfile(profile)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		if code == http.StatusBadRequest {
			response.WriteJSONResponse(c, code, msg, nil)
			return
		}
		if code == http.StatusInternalServerError {
			response.WriteJSONResponse(c, http.StatusInternalServerError, "Failed to save profile", nil)
			return
		}
	}

	response.WriteJSONResponse(c, http.StatusCreated, "Profile created", result)
}

func (h *MahasiswaHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	profile, err := h.usecase.GetProfile(userID)
	if err != nil {
		response.WriteJSONResponse(c, http.StatusNotFound, "Profile not found", nil)
		return
	}

	response.WriteJSONResponse(c, http.StatusOK, "Profile retrieved", profile)
}

func (h *MahasiswaHandler) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	// Ambil form values
	nim := c.PostForm("nim")
	nama := c.PostForm("nama")
	noHP := c.PostForm("no_hp")

	// Ambil file jika dikirim
	photoPath := ""
	file, err := c.FormFile("photo")
	if err == nil {
		path := "uploads/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			response.WriteJSONResponse(c, http.StatusInternalServerError, "Failed to upload file", nil)
			return
		}
		photoPath = path
	}

	// Buat object baru untuk validasi
	profile := &entity.MahasiswaProfile{
		UserID:    userID,
		NIM:       nim,
		Nama:      nama,
		NoHP:      noHP,
		PhotoPath: photoPath,
	}

	// Kirim ke usecase (biar validasi di dalam)
	updated, err := h.usecase.UpdateProfile(profile)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	response.WriteJSONResponse(c, http.StatusOK, "Profile updated", updated)
}
