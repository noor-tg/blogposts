package renderer

import (
	"alnoor/blogposts/reader"
	"bytes"
	"io"
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

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}

		if postRenderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRenderer(b *testing.B) {
	var post = reader.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, post)
	}
}
