package schema

type AddQuestion struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tag     []string `json:"tag"`
	Degree  uint     `json:"degree"`
}

type QuestionMsg struct {
	Title  string   `json:"title"`
	Tag    []string `json:"tag"`
	Degree uint     `json:"degree"`
}
