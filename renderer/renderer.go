package renderer

import (
	"alnoor/blogposts/reader"
	"fmt"
	"io"
)

func Render(w io.Writer, post reader.Post) error {
	_, err := fmt.Fprintf(w, `<h1>%s</h1>`, post.Title)
	return err
}
