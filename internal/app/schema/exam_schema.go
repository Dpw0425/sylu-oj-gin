package schema

type AddExam struct {
	Name    string   `json:"name"`
	Student []string `json:"student"`
}

type AddQuestionToExam struct {
	ID     []int `json:"id"`
	ExamID int   `json:"exam_id"`
}

type ExamStatusResp struct {
	ID       int    `json:"id"`       // 用户 id
	Username string `json:"username"` // 用户名 (学号)
	Status   string `json:"status"`
}

type ExamSummary struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	QuestionNum int64  `json:"question_num"`
	StudentNum  int64  `json:"student_num"`
}
