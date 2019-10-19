package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	pb "github.com/zacharychang/leetcode/hack/readme/leetcode"
	template2 "github.com/zacharychang/leetcode/hack/readme/template"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"text/template"
)

type AllProblemsResponse struct {
	UserName        string          `json:"user_name"`
	NumSolved       int             `json:"num_solved"`
	NumTotal        int             `json:"num_total"`
	AcEasy          int             `json:"ac_easy"`
	AcMedium        int             `json:"ac_medium"`
	AcHard          int             `json:"ac_hard"`
	StatStatusPairs StatStatusPairs `json:"stat_status_pairs"`
}

type QueryProblemResponse struct {
	Data QueryProblemData `json:"data"`
}

type QueryProblemData struct {
	Question QuestionDetail `json:"question"`
}

type QuestionDetail struct {
	QuestionId         string     `json:"questionId"`
	QuestionFrontendId string     `json:"questionFrontedId"`
	QuestionTitle      string     `json:"questionTitle"`
	QuestionTitleSlug  string     `json:"questionTitleSlug"`
	Content            string     `json:"content"`
	TopicTags          []TopicTag `json:"topicTags"`
}

type TopicTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type StatStatusPairs []StatStatusPair

type StatStatusPair struct {
	Stat       Stat       `json:"stat"`
	Status     string     `json:"status"`
	Difficulty Difficulty `json:"difficulty"`
}

type Stat struct {
	QuestionId         int    `json:"question_id"`
	QuestionTitle      string `json:"question__title"`
	QuestionTitleSlug  string `json:"question__title_slug"`
	FrontendQuestionId int    `json:"frontend_question_id"`
}

func (stat StatStatusPair) toTemplateQuestion(topics []TopicTag) template2.Question {

	url := fmt.Sprintf("https://leetcode.com/problems/%s", stat.Stat.QuestionTitleSlug)

	var str bytes.Buffer
	title := fmt.Sprintf("%03d.%s", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitleSlug)
	githubUrl := fmt.Sprintf("[Detail](https://github.com/ZacharyChang/leetcode/tree/master/%s)", title)
	str.Write([]byte(fmt.Sprintf("|%d|[%s](%s)|%s|%s|\n", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitle, url, githubUrl, stat.Difficulty.LevelName())))

	return template2.Question{
		Id:          stat.Stat.FrontendQuestionId,
		Title:       stat.Stat.QuestionTitle,
		LeetcodeUrl: url,
		GithubUrl:   githubUrl,
		Difficulty:  stat.Difficulty.LevelName(),
		Tags: func(topics []TopicTag) string {
			var str bytes.Buffer
			for i, topic := range topics {
				if i > 0 {
					str.Write([]byte(", "))
				}
				str.Write([]byte("["))
				str.Write([]byte(topic.Name))
				str.Write([]byte("]"))
			}
			return str.String()
		}(topics),
	}
}

type Difficulty struct {
	Level int `json:"level"`
}

func (d Difficulty) LevelName() string {
	switch d.Level {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Hard"
	}
	return "Unknown"
}

type AllTagsResponse struct {
	Topics []Topic `json:"topics"`
}

type Topic struct {
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Questions []int  `json:"questions"`
}

func (s StatStatusPairs) Len() int {
	return len(s)
}

func (s StatStatusPairs) Less(i, j int) bool {
	// sort by frontend question id
	if s[i].Stat.FrontendQuestionId < s[j].Stat.FrontendQuestionId {
		return true
	}
	return false
}

func (s StatStatusPairs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var (
	username    = flag.String("username", os.Getenv("LEETCODE_USERNAME"), "")
	password    = flag.String("password", os.Getenv("LEETCODE_PASSWORD"), "")
	templateDir = flag.String("template-dir", "./hack/readme/template", "")
	outputDir   = flag.String("output-dir", "./", "")
)

func getTopis() (map[int][]TopicTag, []TopicTag) {
	resp, err := http.Get("https://leetcode.com/problems/api/tags/")
	defer func() {
		err = resp.Body.Close()
		checkErr(err)
	}()
	checkErr(err)

	res, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	var data AllTagsResponse
	err = json.Unmarshal(res, &data)
	checkErr(err)

	resMap := make(map[int][]TopicTag, 0)
	resArray := make([]TopicTag, 0)
	if len(data.Topics) > 0 {
		for _, v := range data.Topics {
			for _, questionId := range v.Questions {
				//_, exist := resMap[questionId]
				resMap[questionId] = append(resMap[questionId], TopicTag{
					Name: v.Name,
					Slug: v.Slug,
				})
			}
			resArray = append(resArray, TopicTag{
				Name: v.Name,
				Slug: v.Slug,
			})
		}
	}
	return resMap, resArray
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewLeetcodeServiceClient(conn)

	// login first
	resp, err := client.Login(context.Background(), &pb.LoginRequest{
		Name:     *username,
		Password: *password,
	})
	checkErr(err)

	// then get the status of all problems
	resp, err = client.ListAllProblems(context.Background(), &pb.ListAllProblemsRequest{})

	checkErr(err)
	var questionStatus AllProblemsResponse
	err = json.Unmarshal([]byte(resp.Result), &questionStatus)
	checkErr(err)

	sort.Sort(questionStatus.StatStatusPairs)

	updateDocs(questionStatus.StatStatusPairs, client)
}

func updateDocs(stats StatStatusPairs, c pb.LeetcodeServiceClient) {
	topicMap, allTopics := getTopis()

	var str bytes.Buffer
	var templateProblems []template2.Question
	templateProblemsByTopic := make(map[string][]template2.Question, 0)
	var templateTopics []template2.TopicTag
	for _, v := range allTopics {
		templateTopics = append(templateTopics, template2.TopicTag{
			Name: v.Name,
			Slug: v.Slug,
		})
	}

	for _, stat := range stats {
		url := fmt.Sprintf("https://leetcode.com/problems/%s", stat.Stat.QuestionTitleSlug)
		// only show the accept problems
		if stat.Status == "ac" {
			title := fmt.Sprintf("%03d.%s", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitleSlug)
			githubUrl := fmt.Sprintf("[Detail](https://github.com/ZacharyChang/leetcode/tree/master/%s)", title)
			str.Write([]byte(fmt.Sprintf("|%d|[%s](%s)|%s|%s|\n", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitle, url, githubUrl, stat.Difficulty.LevelName())))

			templateProblems = append(templateProblems, stat.toTemplateQuestion(topicMap[stat.Stat.QuestionId]))
			for _, topic := range topicMap[stat.Stat.QuestionId] {
				templateProblemsByTopic[topic.Slug] = append(templateProblemsByTopic[topic.Slug], stat.toTemplateQuestion(topicMap[stat.Stat.QuestionId]))
			}
			updateProblemDocs(title, stat.Stat.QuestionTitleSlug, c)
		}
	}
	updateAllProblemDocs(template2.Page{
		Title:     "Algorithm",
		Questions: templateProblems,
		Topics:    templateTopics,
	})
	for topic, problems := range templateProblemsByTopic {
		updateTagDocs(topic, template2.Page{
			Title:     topic,
			Questions: problems,
			Topics:    templateTopics,
		})
	}
}

func updateProblemDocs(title string, titleSlug string, c pb.LeetcodeServiceClient) {
	dir := *outputDir + title
	readmeFile := dir + "/README.md"
	if _, err := os.Stat(readmeFile); os.IsNotExist(err) {

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, os.ModePerm)
			checkErr(err)
		}

		resp, err := c.QueryProblem(context.Background(), &pb.QueryProblemRequest{
			Slug: titleSlug,
		})
		checkErr(err)

		f, err := os.OpenFile(readmeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		checkErr(err)

		var queryProblemResponse QueryProblemResponse
		err = json.Unmarshal([]byte(resp.Result), &queryProblemResponse)
		checkErr(err)

		templ, err := template.ParseFiles(*templateDir + "/problem.tmpl")
		checkErr(err)
		err = templ.Execute(f, queryProblemResponse.Data.Question)
		checkErr(err)
	}
}

func updateAllProblemDocs(page template2.Page) {
	f, err := os.OpenFile(*outputDir+"README.md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	checkErr(err)

	templ, err := template.ParseFiles(*templateDir + "/readme.tmpl")
	checkErr(err)
	err = templ.Execute(f, page)
	checkErr(err)
}

func updateTagDocs(slugTitle string, page template2.Page) {
	f, err := os.OpenFile(*outputDir+"tags/"+slugTitle+".md", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	checkErr(err)

	templ, err := template.ParseFiles(*templateDir + "/readme.tmpl")
	checkErr(err)
	err = templ.Execute(f, page)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
