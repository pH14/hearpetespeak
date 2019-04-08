package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Topic struct {
	Categories string
	Direct_link string `yaml:"direct_link"`
	Tags []string
}

type Source struct {
	Title string
	Date string
	Description string
	Topics []Topic `yaml:"topics"`
}

//
//title: InForum
//date: 2019-03-28
//description:
//
//topics:
//-
//  categories: "All"
//  direct_link: "https://www.motherjones.com/politics/2019/03/pete-buttigieg-south-bend-mueller-elizabeth-warren/"
//-
//  categories: "Values"
//  direct_link: "https://www.facebook.com/INFORUMsf/videos/308353089832485/?t=233"
//  tags:
//  - political involvement
//  - civic volunteerism
//  - teaching
//  - journalism
//  - impact


func main() {
	sources_path := "./data/sources/"

	files, err := ioutil.ReadDir(sources_path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())

		source_file, err := ioutil.ReadFile(sources_path + file.Name())

		if err != nil {
			log.Fatal(err)
		}

		entry := Source{}

		err = yaml.Unmarshal([]byte(source_file), &entry)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		newFile, err := os.Create("./tmp/" + file.Name())
		defer newFile.Close()

		if err != nil {
			return
		}

		newFile.WriteString(entry.Title + "\n")
		newFile.WriteString(entry.Date)


		for _, topic := range entry.Topics {
			fmt.Printf("%v", topic)
			if strings.HasPrefix("https://you", topic.Direct_link) {

			}
		}


		//fmt.Printf("--- t:\n%v\n\n", entry)
	}
}
