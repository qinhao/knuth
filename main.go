package main

import (
	"fmt"
	"os"
)

var head = `#cover
  h1.book-title Knuth Interview 2006
  h2.author by Donald Ervin Knuth

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
	var (
		name        = "knuth.pug"
		chapterNums = 97
	)

	var content = head
	for i := 1; i <= chapterNums; i++ {
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
