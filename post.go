package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	DescriptionSeparator = "Description: "
	TagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// read line to get title info
	title := readMetaLine(titleSeparator)
	description := readMetaLine(DescriptionSeparator)
	tags := strings.Split(readMetaLine(TagsSeparator), ", ")

	// convert file content to string and get the chars after 7 char to end as title
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}, nil
}
