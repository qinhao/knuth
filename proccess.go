package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var head = `#cover
  h1.book-title Knuth Interview 2006
  h2.author by Donald E. Knuth

`

var foot = `
template#page-footer
  style(type='text/css').
    .pdf-footer {
      font-family: Helvetica;
      font-size: 10px;
      width: 100%;
      text-align: center;
    }  
    .pdf-footer #[span.pageNumber]

style
  include:scss knuth.scss
`

var temp = `
.chapter
  include:markdown-it(html=true) sections/%d.md
`

func main() {
	sharding()
	generatePug()
}

func generatePug() {
	var (
		name        = "knuth.pug"
		chapterNums = 97
	)

	var content = head
	for i := 0; i <= chapterNums; i++ {
		content += fmt.Sprintf(temp, i)
	}
	content += foot

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		fmt.Println("Error writing to thr file with os.OpenFile and *File.WriteString method.", err)
	}

	f.Sync()
}

func sharding() {
	f, err := os.Open("./knuth-interview-2006/README.md")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f.Close()

	var (
		sectionNum int
		content    string
		oldContent string
		chapter    string
		isEnd      bool
	)
	var (
		br      = bufio.NewReader(f)
		summary = "## Summary"
	)

	for {
		a, _, err := br.ReadLine()
		if err == io.EOF {
			isEnd = true
			content = strings.Replace(oldContent, "Code used to originally pull the above", "End of The Book", -1)
		}

		var name string
		if strings.Contains(string(a), "](http") || isEnd {
			name = fmt.Sprintf("./sections/%d.md", sectionNum)
			if chapter == "" {
				chapter = summary
				content = strings.Replace(content, "Donald Knuth Interview 2006\n===========================", "", -1)
			}
			content = chapter + content

			f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				fmt.Println("Failed to open the file", err.Error())
				os.Exit(2)
			}
			defer f.Close()

			if _, err := f.WriteString(content); err != nil {
				fmt.Println("Error writing to thr file with os.OpenFile and *File.WriteString method.", err)
				os.Exit(2)
			}

			f.Sync()

			sectionNum++

			chapter = "## " + string(a) + "\n"
		}

		if strings.Contains(string(a), "------------------------------") {
			oldContent = content
			content = ""
		} else {
			content += string(a) + "\n"
		}

		if isEnd {
			break
		}

	}
}
