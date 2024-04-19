package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	DescriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	// read line to get title info
	title := readLine(scanner, titleSeparator)
	description := readLine(scanner, DescriptionSeparator)

	// convert file content to string and get the chars after 7 char to end as title
	post := Post{title, description}

	return post, nil
}

func readLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}
