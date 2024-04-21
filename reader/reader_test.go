package reader_test

import (
	"alnoor/blogposts/reader"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail!!")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: title
Description: desc
Tags: blog, work, test
---
this is the body`
	)
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
	}

	posts, err := reader.NewPostsFromFS(fs)

	assertNoError(t, err)

	assertSameLength(t, posts, fs)

	got := posts[0]

	assertPost(t, got, reader.Post{
		Title:       "title",
		Description: "desc",
		Tags:        []string{"blog", "work", "test"},
		Body:        "this is the body",
	})
}

func assertPost(t *testing.T, got reader.Post, want reader.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertSameLength(t *testing.T, posts []reader.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted  %d posts", len(posts), len(fs))
	}
}
