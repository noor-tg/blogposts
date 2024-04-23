package renderer

import (
	"alnoor/blogposts/reader"
	"bytes"
	"io"
	"os"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	approvals.UseFolder("testData")
	posts, err := reader.NewPostsFromFS(os.DirFS("../assets/examples"))
	if err != nil {
		t.Fatal(err)
	}

	post := posts[0]
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

	t.Run("it converts multiple posts to html list in index file", func(t *testing.T) {
		buf := bytes.Buffer{}

		if postRenderer.RenderIndex(&buf, posts); err != nil {
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
