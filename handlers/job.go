package handlers

import (
	"github.com/Tekalig/job-finder-go/hasura"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	hasuraClient *hasura.Client
}

func (h *JobHandler) CreateJob(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CompanyID   int    `json:"companyId"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mutation := `
        mutation CreateJob($input: jobs_insert_input!) {
            insert_jobs_one(object: $input) {
                job_id
                title
            }
        }
    `

	var response struct {
		InsertJobsOne struct {
			JobID int    `json:"job_id"`
			Title string `json:"title"`
		} `json:"insert_jobs_one"`
	}

	if err := h.hasuraClient.Execute(mutation, map[string]interface{}{
		"input": input,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, response)
}

func (h *JobHandler) GetJobs(c *gin.Context) {
	query := `
        query GetJobs {
            jobs {
                job_id
                title
                description
                company_id
            }
        }
    `

	var response struct {
		Jobs []struct {
			JobID       int    `json:"job_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			CompanyID   int    `json:"company_id"`
		} `json:"jobs"`
	}

	if err := h.hasuraClient.Execute(query, nil, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}

func (h *JobHandler) UpdateJob(c *gin.Context) {
	var input struct {
		JobID       int    `json:"jobId"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mutation := `
        mutation UpdateJob($id: Int!, $input: jobs_set_input!) {
            update_jobs_by_pk(pk_columns: {job_id: $id}, _set: $input) {
                job_id
                title
            }
        }
    `

	var response struct {
		UpdateJobsByPk struct {
			JobID int    `json:"job_id"`
			Title string `json:"title"`
		} `json:"update_jobs_by_pk"`
	}

	if err := h.hasuraClient.Execute(mutation, map[string]interface{}{
		"id":    input.JobID,
		"input": input,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}

func (h *JobHandler) DeleteJob(c *gin.Context) {
	var input struct {
		JobID int `json:"jobId"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	mutation := `
        mutation DeleteJob($id: Int!) {
            delete_jobs_by_pk(job_id: $id) {
                job_id
            }
        }
    `

	var response struct {
		DeleteJobsByPk struct {
			JobID int `json:"job_id"`
		} `json:"delete_jobs_by_pk"`
	}

	if err := h.hasuraClient.Execute(mutation, map[string]interface{}{
		"id": input.JobID,
	}, &response); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}

func CreateJobHandler(c *gin.Context, hasuraClient *hasura.Client) {
	handler := JobHandler{
		hasuraClient: hasuraClient,
	}
	handler.CreateJob(c)
}

func GetJobsHandler(c *gin.Context, hasuraClient *hasura.Client) {
	handler := JobHandler{
		hasuraClient: hasuraClient,
	}
	handler.GetJobs(c)
}

func UpdateJobHandler(c *gin.Context, hasuraClient *hasura.Client) {
	handler := JobHandler{
		hasuraClient: hasuraClient,
	}
	handler.UpdateJob(c)
}

func DeleteJobHandler(c *gin.Context, hasuraClient *hasura.Client) {
	handler := JobHandler{
		hasuraClient: hasuraClient,
	}
	handler.DeleteJob(c)
}
