package response

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role"`
}

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}

func WriteJSONResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	// Set the content type to application/json
	c.Header("Content-Type", "application/json")
	// Set the status code
	c.Status(statusCode)
	c.JSON(statusCode, Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
