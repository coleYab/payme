package services

import (
	"fmt"
	"payme/internal/models"
	"payme/internal/repository"

	"github.com/google/uuid"
)

type TaskServices struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) *TaskServices {
	return &TaskServices{taskRepository}
}

func (us *TaskServices) SaveTask(task *models.Task) error {
	return us.taskRepository.Save(task)
}

func (us *TaskServices) FindPublicTasks(skip, limit int) ([]*models.Task, error) {
	return us.taskRepository.FindPublicTasks()
}

func (us *TaskServices) CreateTask(id, createdBy, statement, testerCode, status  string, testCases []struct {
	Input string
	Output string
}) (*models.Task, error) {
	newTestCases := []models.TestCase{}
	for _, testCase := range testCases {
		newTestCases = append(newTestCases, models.NewTestCase(uuid.NewString(), testCase.Input, testCase.Output))
	}
	task, _ := models.NewTask(uuid.NewString(), createdBy, statement, testerCode, status, newTestCases)
	if err := us.taskRepository.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (us *TaskServices) FindTasks(skip, limit int) ([]*models.Task, error) {
	return us.taskRepository.FindTasks()
}

func (us *TaskServices) FindTaskByID(id string) (*models.Task, error) {
	return us.taskRepository.FindTaskByID(id)
}

func (us *TaskServices) UpdateStatus(id, newStatus string) (*models.Task, error) {
	task, err := us.taskRepository.FindTaskByID(id)
	if err != nil {
		return task, err
	}

	task.UpdateStatus(newStatus)
	if err := us.taskRepository.Update(task); err != nil {
		return task, err
	}

	return task, err
}

func (us *TaskServices) UpdateTask(id, firstName, lastName string) (*models.Task, error) {
	task, err := us.taskRepository.FindTaskByID(id)
	if err != nil {
		return nil, err
	}

	// task.UpdateTask(firstName, lastName)
	if err := us.taskRepository.Update(task); err != nil {
		return nil, err
	}

	return task, err
}

func (us *TaskServices) DeleteTask(id string) error {
	return us.taskRepository.Delete(id)
}

func (us *TaskServices) CreateSubmission(taskId, id, language, sourceCode string) (*models.Submission, error) {
	task, err := us.taskRepository.FindTaskByID(taskId)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}
	submission := models.NewSubmission(id, language, sourceCode)
	task.AddSubmission(submission)
	us.taskRepository.Update(task)
	return &submission, nil
}
