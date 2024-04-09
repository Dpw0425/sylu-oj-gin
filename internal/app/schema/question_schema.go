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

type Question struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tag     []string `json:"tag"`
	Degree  int      `json:"degree"`
}

type Answer struct {
	Answer string `json:"answer"`
}
