package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	source = "../../README.md"
	target = "../../out/"
)

func main() {
	b, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatalf(err.Error())
	}

	tagMap := make(map[string][]string, 0)

	str := string(b) // convert content to a 'string'
	from := "|-|-----|--------|----------|----|"
	to := "## Link"
	for _, v := range strings.Split(str[strings.Index(str, from)+len(from):strings.Index(str, to)], "\n") {
		if len(strings.Split(v, "|")) <= 5 {
			continue
		}
		tags := strings.Split(v, "|")[5]
		for _, t := range strings.Split(tags, ",") {
			name := strings.Trim(t, " ")
			name = strings.ReplaceAll(name, "[", "")
			name = strings.ReplaceAll(name, "]", "")
			if tagMap[name] == nil {
				tagMap[name] = make([]string, 0)
			}
			tagMap[name] = append(tagMap[name], v)
		}
	}
	for k, v := range tagMap {
		os.MkdirAll(target, os.ModePerm)
		file, err := os.Create(target + camel2Snake(k) + ".md")
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer file.Close()
		file.Write([]byte(strings.Join(v, "\n")))
	}
}

func camel2Snake(str string) string {
	if len(str) == 0 {
		return str
	}
	res := make([]byte, 0)
	for i, v := range str {
		if i > 0 && v >= 'A' && v <= 'Z' {
			res = append(res, '-')
		}
		res = append(res, toLower(byte(v)))
	}
	return string(res)
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}