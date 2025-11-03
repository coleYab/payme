package task

import (
	"net/http"
	taskdto "payme/internal/dto/tasks"
	"payme/internal/services"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	taskService *services.TaskServices
}

// NewTaskHandler creates a new TaskHandler instance.
func NewTaskHandler(taskService *services.TaskServices) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// FindPublicTasks godoc
// @Summary Get public tasks
// @Description Retrieve a paginated list of all public tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} map[string]interface{} "List of public tasks"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks [get]
func (u *TaskHandler) FindPublicTasks(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 50
	}

	tasks, err := u.taskService.FindPublicTasks(page*limit+1, limit)
	tasksDto := []taskdto.PublicTaskResponse{}
	for _, task := range tasks {
		tasksDto = append(tasksDto, taskdto.FromPublicTask(task))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"tasks":   tasksDto,
		"metadata": map[string]any{
			"first":    0,
			"next":     page + 1,
			"current":  page,
			"previous": max(page-1, 0),
			"count":    len(tasksDto),
			"last":     100,
			"limit":    limit,
		},
	})
}

// FindTasks godoc
// @Summary Get private tasks
// @Description Retrieve a paginated list of all private tasks (requires authentication)
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} map[string]interface{} "List of private tasks"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks/private [get]
func (u *TaskHandler) FindTasks(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 0
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 50
	}

	tasks, err := u.taskService.FindTasks(page*limit+1, limit)
	tasksDto := []taskdto.PrivateTaskResponse{}
	for _, task := range tasks {
		tasksDto = append(tasksDto, taskdto.FromPrivateTask(task))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"tasks":   tasksDto,
		"metadata": map[string]any{
			"first":    0,
			"next":     page + 1,
			"current":  page,
			"previous": max(page-1, 0),
			"count":    len(tasksDto),
			"last":     100,
			"limit":    limit,
		},
	})
}

// FindTaskByID godoc
// @Summary Get task by ID
// @Description Retrieve details of a specific task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} taskdto.PublicTaskResponse "Task found"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /api/v1/tasks/{id} [get]
func (u *TaskHandler) FindTaskByID(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	task, err := u.taskService.FindTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"task":    taskdto.FromPublicTask(task),
	})
}

// UpdateTaskStatus godoc
// @Summary Update task status
// @Description Update the status of a specific task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param data body taskdto.UpdateTaskStatusDto true "New task status"
// @Success 200 {object} taskdto.PublicTaskResponse "Updated task"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks/{id}/status [patch]
func (u *TaskHandler) UpdateTaskStatus(ctx *gin.Context) {
	var dto taskdto.UpdateTaskStatusDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse dto", "error": err.Error()})
		return
	}

	id := ctx.Param("id")
	newTask, err := u.taskService.UpdateStatus(id, dto.NewStatus)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to update task status", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"task":    taskdto.FromPublicTask(newTask),
	})
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with statement, test cases, and metadata
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param data body taskdto.CreateTaskDto true "Task data"
// @Success 200 {object} taskdto.PrivateTaskResponse "Created task"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks [post]
func (u *TaskHandler) CreateTask(ctx *gin.Context) {
	var createTaskDto taskdto.CreateTaskDto
	if err := ctx.ShouldBindJSON(&createTaskDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse dto", "error": err.Error()})
		return
	}

	testCases := []struct {
		Input  string
		Output string
	}{}
	for _, t := range createTaskDto.TestCases {
		testCases = append(testCases, struct {
			Input  string
			Output string
		}{Input: t.Input, Output: t.Output})
	}

	newTask, err := u.taskService.CreateTask(uuid.NewString(), createTaskDto.CreatedBy, createTaskDto.Statement, createTaskDto.TesterCode, createTaskDto.Status, testCases)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to create task", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "task": taskdto.FromPrivateTask(newTask)})
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]string "Deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /api/v1/tasks/{id} [delete]
func (u *TaskHandler) DeleteTask(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is required"})
		return
	}

	if err := u.taskService.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// FindSubmissionsByTask godoc
// @Summary Get submissions for a task
// @Description Retrieve all submissions related to a specific task
// @Tags submissions
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]interface{} "List of submissions"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks/{id}/submissions [get]
func (u *TaskHandler) FindSubmissionsByTask(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task id is required"})
		return
	}

	task, err := u.taskService.FindTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task not found"})
		return
	}

	submissionsDto := []taskdto.SubmissionResponse{}
	for _, submission := range task.Submissions {
		submissionsDto = append(submissionsDto, taskdto.FromSubmission(&submission))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "submissions": submissionsDto})
}

// CreateSubmission godoc
// @Summary Create a submission for a task
// @Description Submit a new solution for a specific task
// @Tags submissions
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param data body taskdto.CreateSubmissionDto true "Submission data"
// @Success 200 {object} taskdto.SubmissionResponse "Created submission"
// @Failure 400 {object} map[string]string "Invalid request"
// @Router /api/v1/tasks/{id}/submit [post]
func (u *TaskHandler) CreateSubmission(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), " ")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "task id is required"})
		return
	}

	var dto taskdto.CreateSubmissionDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse dto", "error": err.Error()})
		return
	}

	newSubmission, err := u.taskService.CreateSubmission(id, uuid.NewString(), dto.Language, dto.SourceCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to create submission", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "submission": taskdto.FromSubmission(newSubmission)})
}

