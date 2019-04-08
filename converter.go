package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"
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

		newFile, err := os.Create("./tmp/" + file.Name())
		defer newFile.Close()

		if err != nil {
			return
		}

		newLines := make([]string, 1)
		for _, topic := range source.Topics {
			if strings.HasPrefix(topic.Direct_link, "https://youtu") {
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

				newLines = append(newLines, fmt.Sprintf("%v:%v %v | %v\n", minutes, seconds, topic.Prompt, strings.Join(topic.Tags, ", ")))
			}
		}

		if len(newLines) > 0 {
			_, err := newFile.WriteString(fmt.Sprintf("%v - %v\n", source.Date, source.Title))

			if err != nil {
				panic(err)
			}

			for _, newLine := range newLines {
				_, err := newFile.WriteString(fmt.Sprintf("%v", newLine))

				if err != nil {
					panic(err)
				}
			}
		}
	}
}
