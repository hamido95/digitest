package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Worker struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func main() {
	r := gin.Default()
	r.GET("/worker/:id", getWorker)
	r.Run()
}

func getWorker(c *gin.Context) {
	var workers = map[string]Worker{
		"1": {
			ID:          "1",
			Name:        "ahmad",
			Description: ("greetings from ahmad jan"),
			Status:      "active",
		},
		"2": {
			ID:          "2",
			Name:        "mahmood",
			Description: ("greetings from mahmood jan"),
			Status:      "inactive",
		},
	}

	workerID := c.Param("id")
	worker, ok := workers[workerID]
	if ok {
		c.JSON(http.StatusOK, worker)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Worker %s not found", workerID)})
	}
}
