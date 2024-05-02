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
