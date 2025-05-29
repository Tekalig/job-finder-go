package handlers

import (
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/Tekalig/job-finder-go/pkg/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	hasuraClient *hasura.Client
	jwtSecret    string
}

func (h *LoginHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := `
        query GetUser($email: String!) {
            users(where: {email: {_eq: $email}}) {
                user_id
                email
                password
                role
            }
        }
    `

	var response struct {
		Users []struct {
			UserID   string `json:"user_id"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Role     string `json:"role"`
		} `json:"users"`
	}

	if err := h.hasuraClient.Execute(query, map[string]interface{}{
		"email": input.Email,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(response.Users) == 0 {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	user := response.Users[0]

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateToken(user.UserID, user.Role, h.jwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func Login(c *gin.Context, hasuraClient *hasura.Client, jwtSecret string) {
	handler := LoginHandler{
		hasuraClient: hasuraClient,
		jwtSecret:    jwtSecret,
	}
	handler.Login(c)
}
