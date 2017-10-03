package controllers

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/gin-gonic/gin"
	"github.com/pedrocelso/go-rest-service/lib/services/task"
)

// CreateTask creates an Task
func CreateTask(c *gin.Context) {
	var tsk *task.Task
	var err error
	var output *task.Task
	ctx := appengine.NewContext(c.Request)

	if err = c.BindJSON(&tsk); err == nil {
		if output, err = task.Create(ctx, tsk); err == nil {
			c.JSON(http.StatusOK, ResponseObject{"task": output})
		}
	}

	if err != nil {
		log.Errorf(ctx, "ERROR: %v", err.Error())
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}

// GetTask based on its ID
func GetTask(c *gin.Context) {
	var err error
	var output *task.Task
	tskId := c.Param("taskId")
	ctx := appengine.NewContext(c.Request)

	if output, err = task.GetByID(ctx, tskId); err == nil {
		c.JSON(http.StatusOK, output)
	}
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, ResponseObject{"error": err.Error()})
	}
}
