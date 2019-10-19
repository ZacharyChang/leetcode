package template

type Question struct {
	Id          int
	Title       string
	LeetcodeUrl string
	GithubUrl   string
	Difficulty  string
	Tags        string
}

type Page struct {
	Title     string
	Questions []Question
	Topics    []TopicTag
}

type TopicTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
