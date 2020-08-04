package Aiken

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type Options struct {
	Desc   string `json:"desc"`
	Answer string `json:"answer"`
}

type Aiken struct {
	Question string    `json:"question"`
	Options  []Options `json:"options"`
	Answer   string    `json:"answer"`
}

func ReadAiken(path string) (result []Aiken, err error) {
	names, err := ioutil.ReadFile(path)

	//question := make(map[string]string)
	if err == nil {
		start := true
		content := string(names)
		contentPerLines := strings.Split(content, "\n")

		var tmpData Aiken
		var tmpChoices []Options

		for _, val := range contentPerLines {
			line := strings.TrimSuffix(strings.Trim(val, " \r\n"), "\r\n")
			if line != "" {
				isAnswer, _ := regexp.MatchString("^ANSWER:{1}", line)
				if start {
					tmpData.Question = strings.TrimSuffix(line, "\r")
					start = false
				} else if isAnswer {
					answer := regexp.MustCompile(`^ANSWER:{1}`).ReplaceAllString(line, "")
					tmpData.Answer = strings.TrimSuffix(answer, "\r")
					tmpData.Options = tmpChoices
					result = append(result, tmpData)
					start = true
					tmpData = Aiken{}
					tmpChoices = []Options{}
				} else {
					options := strings.SplitN(line, " ", 2)
					choice := Options{
						Answer: strings.Trim(options[1], " \r"),
						Desc:   regexp.MustCompile(`\.|\)`).ReplaceAllString(options[0], ""),
					}
					tmpChoices = append(tmpChoices, choice)
				}
			}
		}
	}
	return
}
