// handlers/employer.go
package handlers

import (
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/gin-gonic/gin"
)

type EmployerHandler struct {
	hasuraClient *hasura.Client
}

func (h *EmployerHandler) Signup(c *gin.Context) {
	var input struct {
		CompanyName        string `json:"companyName"`
		Email              string `json:"email"`
		Password           string `json:"password"`
		CompanyDescription string `json:"companyDescription"`
		ContactNumber      string `json:"contactNumber"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create employer mutation
	mutation := `
        mutation CreateEmployer($input: employers_insert_input!) {
            insert_employers_one(object: $input) {
                company_id
                email
            }
        }
    `

	// Execute mutation
	var response struct {
		InsertEmployersOne struct {
			CompanyID int    `json:"company_id"`
			Email     string `json:"email"`
		} `json:"insert_employers_one"`
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
func EmployerSignup(c *gin.Context) {
	handler := EmployerHandler{
		hasuraClient: hasura.NewClient("your-hasura-endpoint", "your-admin-secret"),
	}
	handler.Signup(c)
}
