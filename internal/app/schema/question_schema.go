package schema

type AddQuestion struct {
	Title   string     `json:"title"`
	Content string     `json:"content"`
	Tag     []string   `json:"tag"`
	Degree  uint       `json:"degree"`
	IO      []TestData `json:"io"`
}

type TestData struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type QuestionMsg struct {
	ID     uint     `json:"id"`
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
	ID     int    `json:"id"`
	Answer string `json:"answer"`
}
