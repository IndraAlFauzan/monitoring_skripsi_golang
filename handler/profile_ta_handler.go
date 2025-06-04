package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/mahasiswa"
	"github.com/indraalfauzan/monitoring_skripsi_golang/response"
)

type ProfilTAHandler struct {
	usecase domain.ProfileTAUseCase
}

func NewProfilTAHandler(uc domain.ProfileTAUseCase) *ProfilTAHandler {
	return &ProfilTAHandler{uc}
}

func (h *ProfilTAHandler) AjukanTA(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	var req struct {
		Judul    string `json:"judul"`
		Dosen1ID int    `json:"dosen1_id"`
		Dosen2ID int    `json:"dosen2_id"`
	}

	if err := c.ShouldBind(&req); err != nil {
		response.WriteJSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	result, err := h.usecase.AjukanTA(userID, req.Judul, req.Dosen1ID, req.Dosen2ID)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	response.WriteJSONResponse(c, http.StatusCreated, "Pengajuan TA berhasil", result)
}
