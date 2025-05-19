package handler

import "github.com/gin-gonic/gin"

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	// Fetch user profile from the database using userID
	roleName := c.MustGet("role").(string)

	result := struct {
		Id       int    `json:"id"`
		RoleName string `json:"role"`
	}{
		Id: userID,

		RoleName: roleName,
	}

	WriteJSONResponse(c, 200, "Success", result)
}
