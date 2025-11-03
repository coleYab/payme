package task

import "github.com/gin-gonic/gin"

type TaskRoutes struct{
	handler *TaskHandler
}

func NewTaskRoutes(handler *TaskHandler) *TaskRoutes {
	return &TaskRoutes{handler: handler}
}

func (u *TaskRoutes) RegisterRoutes(routes *gin.RouterGroup) {
	routes.GET("/", u.handler.FindPublicTasks)
	routes.GET("/all", u.handler.FindTasks)
	routes.POST("/", u.handler.CreateTask)
	routes.POST("/:id/submit", u.handler.CreateSubmission)
	routes.GET("/:id/submissions", u.handler.FindSubmissionsByTask)
	routes.GET("/:id", u.handler.FindTaskByID)
	// routes.PUT("/:id", u.handler.UpdateTask)
	routes.DELETE("/:id", u.handler.DeleteTask)
}
