package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Topic struct {
	Categories string
	Direct_link string `yaml:"direct_link"`
	Prompt string
	Tags []string
}

type Source struct {
	Title string
	Date string
	Description string
	Topics []Topic `yaml:"topics"`
}

func main() {
	sourcesPath := "./data/sources/"

	files, err := ioutil.ReadDir(sourcesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		sourceFile, err := ioutil.ReadFile(sourcesPath + file.Name())

		if err != nil {
			log.Fatal(err)
		}

		source := Source{}

		err = yaml.Unmarshal([]byte(sourceFile), &source)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		newLines := make(map[int]string)
		re := regexp.MustCompile(`(.*)(www)?youtu`)
		for _, topic := range source.Topics {
			if re.MatchString(topic.Direct_link) {
				parse, err := url.Parse(topic.Direct_link)

				if err != nil {
					panic(err)
				}

				timestampAsStr := parse.Query().Get("t")

				undividedTimestamp, err := strconv.ParseInt(timestampAsStr, 10, 32)

				if err != nil {
					continue
				}

				minutes := undividedTimestamp / 60
				seconds := undividedTimestamp % 60

				newLines[int(undividedTimestamp)] = fmt.Sprintf("%v:%v %v | %v\n", minutes, seconds, topic.Prompt, strings.Join(topic.Tags, ", "))
			}
		}

		if len(newLines) > 0 {
			newFile, err := os.Create("./tmp/" + file.Name()[:len(file.Name())-5] + ".txt")
			defer newFile.Close()

			if err != nil {
				return
			}
			_, err = newFile.WriteString(fmt.Sprintf("%v - %v\n", source.Date, source.Title))

			if err != nil {
				panic(err)
			}

			sortableTimestamps := make([]int, len(newLines))
			for k := range newLines {
				sortableTimestamps = append(sortableTimestamps, k)
			}
			sort.Ints(sortableTimestamps)

			for _, timestamp := range sortableTimestamps {
				_, err := newFile.WriteString(fmt.Sprintf("%v", newLines[timestamp]))

				if err != nil {
					panic(err)
				}
			}
		}
	}
}
