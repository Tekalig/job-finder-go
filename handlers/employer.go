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
		mutation CreateEmployer($companyName: String!, $email: String!, $password: String!, $companyDescription: String!, $contactNumber: String!) {
			insert_employers_one(object: {
				company_name: $companyName,
				email: $email,
				password: $password,
				company_description: $companyDescription,
				contact_number: $contactNumber
			}) {
				company_id
				company_name
				email
				company_description
				contact_number
			}
		}
	`

	// Execute mutation
	var response struct {
		InsertEmployersOne struct {
			CompanyID          int    `json:"company_id"`
			CompanyName        string `json:"company_name"`
			Email              string `json:"email"`
			CompanyDescription string `json:"company_description"`
			ContactNumber      string `json:"contact_number"`
		} `json:"insert_employers_one"`
	}

	if err := h.hasuraClient.Execute(mutation, map[string]interface{}{
		"companyName":        input.CompanyName,
		"email":              input.Email,
		"password":           input.Password,
		"companyDescription": input.CompanyDescription,
		"contactNumber":      input.ContactNumber,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"company_id":          response.InsertEmployersOne.CompanyID,
		"company_name":        response.InsertEmployersOne.CompanyName,
		"email":               response.InsertEmployersOne.Email,
		"company_description": response.InsertEmployersOne.CompanyDescription,
		"contact_number":      response.InsertEmployersOne.ContactNumber,
	})
}

func EmployerSignup(c *gin.Context, hasuraClient *hasura.Client) {
	handler := EmployerHandler{
		hasuraClient: hasuraClient,
	}
	handler.Signup(c)
}
