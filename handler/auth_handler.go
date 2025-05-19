package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	"github.com/indraalfauzan/monitoring_skripsi_golang/domain"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
)

type AuthHandler struct {
	UserUsecase domain.UserUseCase
}

func NewAuthHandler(u domain.UserUseCase) *AuthHandler {
	return &AuthHandler{UserUsecase: u}
}

func (h *AuthHandler) RegisterMhs(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("request body"))
		WriteJSONResponse(c, code, msg, nil)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := h.UserUsecase.RegisterMhs(user)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		WriteJSONResponse(c, code, msg, nil)
		return
	}

	// Fetch Role Name (safe)
	roleName := ""
	if user.Role.ID != 0 {
		roleName = user.Role.Name
	}

	res := RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		RoleName: roleName,
	}

	WriteJSONResponse(c, http.StatusCreated, "registered", res)
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   int    `json:"role_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("request body"))
		WriteJSONResponse(c, code, msg, nil)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		RoleID:   req.RoleID,
	}

	user, err := h.UserUsecase.RegisterUser(user)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		WriteJSONResponse(c, code, msg, nil)
		return
	}

	// Fetch Role Name (safe)
	roleName := ""
	if user.Role.ID != 0 {
		roleName = user.Role.Name
	}

	res := RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		RoleName: roleName,
	}

	WriteJSONResponse(c, http.StatusCreated, "registered", res)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("request body"))
		WriteJSONResponse(c, code, msg, nil)
		return
	}

	user, err := h.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		WriteJSONResponse(c, code, msg, nil)
		return
	}
	// jwt token ini akan di generate
	// jika user berhasil login
	// token dari id, role, dan expired
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	secret := os.Getenv("JWT_SECRET") // ini untuk ambil secret key dari env
	tokenString, _ := token.SignedString([]byte(secret))

	authResp := LoginResponse{
		Token:    tokenString,
		UserName: user.Username,
		Email:    user.Email,
		RoleName: user.Role.Name,
	}
	WriteJSONResponse(c, http.StatusOK, "login success", authResp)
}
