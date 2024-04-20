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

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// read line to get title info
	title := readLine(titleSeparator)
	description := readLine(DescriptionSeparator)

	// convert file content to string and get the chars after 7 char to end as title
	return Post{
		Title:       title,
		Description: description,
	}, nil
}
