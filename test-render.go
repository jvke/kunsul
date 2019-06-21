package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/list"
	"github.com/jedib0t/go-pretty/text"
)

type PageData struct {
	Title string
	Ingresses []string
	Services []string
}

func print(title string, content string, prefix string) {
	fmt.Printf("%s:\n", title)
	fmt.Println(strings.Repeat("-", len(title)+1))
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s%s\n", prefix, line)
	}
	fmt.Println()
}

func main() {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)	
	print(pageData.Title
	
}
