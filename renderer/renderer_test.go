package renderer_test

import (
	"alnoor/blogposts/reader"
	"alnoor/blogposts/renderer"
	"bytes"
	"testing"
)

func TestRenderer(t *testing.T) {
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

		got := buf.String()

		want := "<h1>hello world</h1>\n<p>this is description</p>\nTags:<ul><li>go</li><li>tdd</li></ul>"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
