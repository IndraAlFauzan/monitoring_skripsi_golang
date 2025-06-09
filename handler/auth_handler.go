package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/indraalfauzan/monitoring_skripsi_golang/apperror"
	domain "github.com/indraalfauzan/monitoring_skripsi_golang/domain/user"
	"github.com/indraalfauzan/monitoring_skripsi_golang/entity"
	"github.com/indraalfauzan/monitoring_skripsi_golang/response"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	UserUsecase domain.UserUseCase
}

func NewAuthHandler(u domain.UserUseCase) *AuthHandler {
	return &AuthHandler{UserUsecase: u}
}

// ================== Register Mahasiswa ==================
func (h *AuthHandler) RegisterMhs(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("invalid request body"))
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.WriteJSONResponse(c, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	user, err = h.UserUsecase.RegisterMhs(user)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	res := response.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		RoleName: user.Role.Name,
	}

	response.WriteJSONResponse(c, http.StatusCreated, "registered", res)
}

// ================== Register User (Generic) ==================
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		RoleID   int    `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("invalid request body"))
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.WriteJSONResponse(c, http.StatusInternalServerError, "Failed to hash password", nil)
		return
	}

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   req.RoleID,
	}

	user, err = h.UserUsecase.RegisterUser(user)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	res := response.RegisterResponse{
		Username: user.Username,
		Email:    user.Email,
		RoleName: user.Role.Name,
	}

	response.WriteJSONResponse(c, http.StatusCreated, "registered", res)
}

// ================== Login ==================
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		code, msg := apperror.DetermineErrorType(apperror.ValidationError("invalid request body"))
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	user, err := h.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		code, msg := apperror.DetermineErrorType(err)
		response.WriteJSONResponse(c, code, msg, nil)
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		response.WriteJSONResponse(c, http.StatusInternalServerError, "JWT secret not set", nil)
		return
	}

	// Build claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role.Name,
		"exp":     time.Now().Add(30 * time.Minute).Unix(),
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		response.WriteJSONResponse(c, http.StatusInternalServerError, "Token generation failed", nil)
		return
	}

	authResp := response.LoginResponse{
		Token:    tokenString,
		UserName: user.Username,
		Email:    user.Email,
		RoleName: user.Role.Name,
	}

	response.WriteJSONResponse(c, http.StatusOK, "Login Success", authResp)
}
