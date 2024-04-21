package reader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
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

	body := readBody(scanner)

	// convert file content to string and get the chars after 7 char to end as title
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	// over step line
	scanner.Scan()

	// start buffer
	buf := bytes.Buffer{}

	// scan each line for loop until return false (return false with eof)
	for scanner.Scan() {

		// store to buffer scanner string
		fmt.Fprintln(&buf, scanner.Text())
	}

	// add new line to end of buffer string because it is removed by fprintln
	return strings.TrimSuffix(buf.String(), "\n")
}
