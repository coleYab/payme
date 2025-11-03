package tasks

type CreateTestCaseDto struct {
	Input string `json:"input"`
	Output string `json:"output"`
}

type CreateSubmissionDto struct {
	Language string `json:"language"`
	SourceCode string `json:"sourceCode"`
}

type CreateTaskDto struct {
	CreatedBy string `json:"createdBy"`
	Statement string `json:"statement"`
	TesterCode string `json:"testerCode"`
	Status string `json:"status"` // public, private
	TestCases []CreateTestCaseDto `json:"testCases"`
}

type UpdateTaskStatusDto struct {
	NewStatus string `json:"string"`
}
