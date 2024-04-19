package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/nour_dev/blogposts"
)

type StubFailingFS struct{}

func (fs StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("always fail!!")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: hi")},
		"hello-world2.md": {Data: []byte("Title: hola")},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted  %d posts", len(posts), len(fs))
	}

	got := posts[0]

	want := blogposts.Post{Title: "hi"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
