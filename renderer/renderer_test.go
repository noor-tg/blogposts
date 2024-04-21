package renderer_test

import (
	"bytes"
	"testing"

	"github.com/nour_dev/blogposts/reader"
	"github.com/nour_dev/blogposts/renderer"
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

		want := `<h1>hello world</h1>`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
