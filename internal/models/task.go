package models

import "time"

type TestCase struct {
	ID string `bson:"testId"`
	Input string `bson:"input"`
	Output string `bson:"output"`
}

type Submission struct {
	ID string `bson:"submissionId"`
	Language string `bson:"language"`
	SourceCode string `bson:"sourceCode"`
	CreatedAt time.Time `bson:"createdAt"`
}

type Task struct {
	ID string `bson:"id"`
	CreatedBy string `bson:"createdBy"`
	Statement string `bson:"statement"`
	TesterCode string `bson:"testerCode"`
	Status string `bson:"status"` // public, private
	TestCases []TestCase `bson:"testCases"`
	Submissions []Submission `bson:"submissions"`
}

func NewTask(id, createdBy, statement, testerCode, status string, testCases []TestCase) (*Task, error) {
	return &Task{
		ID: id,
		CreatedBy: createdBy,
		Statement: statement,
		Status: status,
		TesterCode: testerCode,
		TestCases: testCases,
		Submissions: []Submission{},
	}, nil
}

func NewSubmission(id, language, sourceCode string)Submission {
	return Submission{
		Language: language,
		ID: id,
		SourceCode: sourceCode,
		CreatedAt: time.Now(),
	}
}

func NewTestCase(id, input, output string) TestCase {
	return TestCase{
		ID: id,
		Input: input,
		Output: output,
	}
}


func (t *Task) UpdateStatus(status string) {
	t.Status = status
}

func (t *Task) AddSubmission(submission Submission) {
	t.Submissions = append(t.Submissions, submission)
}

