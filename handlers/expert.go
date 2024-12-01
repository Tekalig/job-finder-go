package handlers

import (
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/gin-gonic/gin"
)

type ExpertHandler struct {
	hasuraClient *hasura.Client
}

func (h *ExpertHandler) Signup(c *gin.Context) {
	var input struct {
		FullName string `json:"fullName"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create expert mutation
	mutation := `
        mutation CreateExpert($input: experts_insert_input!) {
            insert_experts_one(object: $input) {
                expert_id
                email
            }
        }
    `

	// Execute mutation
	var response struct {
		InsertExpertsOne struct {
			ExpertID int    `json:"expert_id"`
			Email    string `json:"email"`
		} `json:"insert_experts_one"`
	}

	if err := h.hasuraClient.Execute(mutation, map[string]interface{}{
		"input": input,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

// Add this function to fix the error in routes
func ExpertSignup(c *gin.Context) {
	handler := ExpertHandler{
		hasuraClient: hasura.NewClient("your-hasura-endpoint", "your-admin-secret"),
	}
	handler.Signup(c)
}
