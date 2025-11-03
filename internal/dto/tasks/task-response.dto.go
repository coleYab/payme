package tasks

import (
	"payme/internal/models"
	"time"
)

type TestCaseResponse struct {
	ID string `json:"id"`
	Input string `json:"input"`
	Output string `json:"output"`
}


type SubmissionResponse struct {
	ID string `json:"submissionId"`
	Language string `json:"language"`
	SourceCode string `json:"sourceCode"`
	CreatedAt time.Time `json:"createdAt"`
}

func FromSubmission(submission *models.Submission) SubmissionResponse {
	return SubmissionResponse{
		Language: submission.Language,
		ID: submission.ID,
		SourceCode: submission.SourceCode,
		CreatedAt: submission.CreatedAt,
	}
}


type PrivateTaskResponse struct {
	ID string `json:"id"`
	CreatedBy string `json:"createdBy"`
	Statement string `json:"statement"`
	TesterCode string     `json:"testerCode"`
	Status     string     `json:"status"` // public, private
	TestCases []TestCaseResponse `json:"testCases"`
}

type PublicTaskResponse struct {
	ID string `json:"id"`
	CreatedBy string `json:"createdBy"`
	Statement string `json:"statement"`
	TestCases []TestCaseResponse `json:"testCases"`
}

func FromPublicTask(task *models.Task) PublicTaskResponse {
	testCaseResponseDto := []TestCaseResponse{}
	for _, testCase := range task.TestCases {
		testCaseResponseDto = append(testCaseResponseDto, TestCaseResponse{
			ID: testCase.ID,
			Input: testCase.Input,
			Output: testCase.Output,
		})
	}

	return PublicTaskResponse{
		ID: task.ID,
		CreatedBy: task.CreatedBy,
		Statement: task.Statement,
		TestCases: testCaseResponseDto,
	}
}


func FromPrivateTask(task *models.Task) PrivateTaskResponse {
	testCaseResponseDto := []TestCaseResponse{}
	for _, testCase := range task.TestCases {
		testCaseResponseDto = append(testCaseResponseDto, TestCaseResponse{
			ID: testCase.ID,
			Input: testCase.Input,
			Output: testCase.Output,
		})
	}

	return PrivateTaskResponse{
		ID: task.ID,
		Status: task.Status,
		TesterCode: task.TesterCode,
		CreatedBy: task.CreatedBy,
		Statement: task.Statement,
		TestCases: testCaseResponseDto,
	}
}

