package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/indraalfauzan/monitoring_skripsi_golang/response"
)

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	// Fetch user profile from the database using userID
	roleName := c.MustGet("role").(string)

	result := struct {
		Id       int    `json:"id"`
		RoleName string `json:"role"`
	}{
		Id:       userID,
		RoleName: roleName,
	}

	response.WriteJSONResponse(c, 200, "Success", result)
}
