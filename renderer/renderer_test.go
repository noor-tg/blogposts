package renderer_test

import (
	"alnoor/blogposts/reader"
	"alnoor/blogposts/renderer"
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	approvals.UseFolder("testData")
	post := reader.Post{
		Title:       "hello world",
		Body:        "this is post",
		Description: "this is description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.Render(&buf, post)

		if err != nil {
			t.Fatalf("error %s", err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
