package schema

type AddExam struct {
	Name    string   `json:"name"`
	Student []string `json:"student"`
}
