package renderer

import (
	"alnoor/blogposts"
	"alnoor/blogposts/reader"
	"html/template"
	"io"
)

func Render(w io.Writer, post reader.Post) error {
	templ, err := template.ParseFS(blogposts.PostTemplate, "templates/*.gohtml")

	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
