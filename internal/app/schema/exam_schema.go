package schema

type AddExam struct {
	Name    string   `json:"name"`
	Student []string `json:"student"`
}

type AddQuestionToExam struct {
	ID     []int `json:"id"`
	ExamID int   `json:"exam_id"`
}
