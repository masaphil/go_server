package boundary

type AddSourceInputDto struct {
	Title     string
	Statement string
	QuizList  []*SourceQuizDto
	SkillIds  []string
	Average   float32
	Hyojun    float32
}

type SourceQuizDto struct {
	Title     string
	Statement string
	Options   []string
	Answer    string
}
