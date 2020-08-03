package Aiken

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Aiken struct {
	Question string            `json:"question"`
	Options  map[string]string `json:"options"`
	Answer   string            `json:"answer"`
}

func ReadAiken(path string) (result []Aiken, err error) {
	names, err := ioutil.ReadFile(path)

	//question := make(map[string]string)
	if err == nil {
		start := true
		content := string(names)
		contentPerLines := strings.Split(content, "\n")


		var tmpData Aiken
		tmpOptions := make(map[string]string)
		for _, val := range contentPerLines {
			line := strings.Trim(val, " ")
			isAnswer, _ := regexp.MatchString("^ANSWER:{1}", line)
			if line != "" {
				if start {
					tmpData.Question = strings.TrimSuffix(line, "\r")
					start = false
				} else if isAnswer {
					answer := regexp.MustCompile(`^ANSWER:{1}`).ReplaceAllString(line, "")
					tmpData.Answer = strings.TrimSuffix(answer, "\r")
					tmpData.Options = tmpOptions
					result = append(result, tmpData)
					start = true
					tmpData = Aiken{}
					tmpOptions = map[string]string{}
				} else {
					options := strings.SplitN(line, " ", 2)
					tmpOptions[regexp.MustCompile(`\.|\)`).ReplaceAllString(options[0], "")] = strings.Trim(options[1], " \r")
				}
			}
		}
	}
	return
}