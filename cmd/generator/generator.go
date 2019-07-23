package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	source    = flag.String("source", "./README.md", "The file to parse from")
	outputDir = flag.String("outputDir", "./tags/", "The directory to write outputDir files")
)

func main() {
	flag.Parse()

	b, err := ioutil.ReadFile(*source)
	if err != nil {
		log.Fatalf(err.Error())
	}

	tagMap := make(map[string][]string, 0)

	str := string(b) // convert content to a 'string'
	from := "|-|-----|--------|----------|----|"

	for _, v := range strings.Split(str[strings.Index(str, from)+len(from):], "\n") {
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
	footer := make([]string, 0)
	for k := range tagMap {
		footer = append(footer, fmt.Sprintf("[%s]: https://github.com/ZacharyChang/leetcode/tree/master/tags/%s.md", k, camel2Snake(k)))
	}
	sort.Strings(footer)
	for k, v := range tagMap {
		_ = os.MkdirAll(*outputDir, os.ModePerm)
		file, err := os.Create(*outputDir + camel2Snake(k) + ".md")
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer func() {
			_ = file.Close()
		}()
		_, _ = file.Write([]byte("# " + k + "\n"))
		_, _ = file.Write([]byte("|#|Title|Language|Difficulty|Tags|\n|-|-----|--------|----------|----|\n"))
		_, _ = file.Write([]byte(strings.Join(v, "\n")))
		_, _ = file.Write([]byte("\n\n## Link\n"))
		_, _ = file.Write([]byte(strings.Join(footer, "\n")))

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
