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
	username = flag.String("username", os.Getenv("LEETCODE_USERNAME"), "")
	password = flag.String("password", os.Getenv("LEETCODE_PASSWORD"), "")
)

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
	var str bytes.Buffer
	var problems []template2.Question
	for _, stat := range questionStatus.StatStatusPairs {
		url := fmt.Sprintf("https://leetcode.com/problems/%s", stat.Stat.QuestionTitleSlug)
		// only show the accept problems
		if stat.Status == "ac" {
			title := fmt.Sprintf("%03d.%s", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitleSlug)
			githubUrl := fmt.Sprintf("[Detail](https://github.com/ZacharyChang/leetcode/tree/master/%s)", title)
			str.Write([]byte(fmt.Sprintf("|%d|[%s](%s)|%s|%s|\n", stat.Stat.FrontendQuestionId, stat.Stat.QuestionTitle, url, githubUrl, stat.Difficulty.LevelName())))

			problems = append(problems, template2.Question{
				Id:          stat.Stat.FrontendQuestionId,
				Title:       stat.Stat.QuestionTitle,
				LeetcodeUrl: url,
				GithubUrl:   githubUrl,
				Difficulty:  stat.Difficulty.LevelName(),
			})

			dir := "../../" + title
			readmeFile := dir + "/README.md"
			if _, err := os.Stat(readmeFile); os.IsNotExist(err) {

				fmt.Println(readmeFile)

				if _, err := os.Stat(dir); os.IsNotExist(err) {
					err = os.Mkdir(dir, os.ModePerm)
					checkErr(err)
				}

				resp, err := client.QueryProblem(context.Background(), &pb.QueryProblemRequest{
					Slug: stat.Stat.QuestionTitleSlug,
				})
				checkErr(err)

				f, err := os.OpenFile(readmeFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				checkErr(err)

				var queryProblemResponse QueryProblemResponse
				err = json.Unmarshal([]byte(resp.Result), &queryProblemResponse)
				checkErr(err)

				templ, err := template.ParseFiles("template/problem.tmpl")
				checkErr(err)
				err = templ.Execute(f, queryProblemResponse.Data.Question)
				checkErr(err)

			}
		}
	}

	templ, err := template.ParseFiles("template/readme.tmpl")
	checkErr(err)
	err = templ.Execute(os.Stdout, problems)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
