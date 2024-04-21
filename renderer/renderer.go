package renderer

import (
	"alnoor/blogposts"
	"alnoor/blogposts/reader"
	"html/template"
	"io"
)

func Render(w io.Writer, post reader.Post) error {
	// parse template from embedded fs
	templ, err := template.ParseFS(blogposts.PostTemplate, "templates/*.gohtml")

	if err != nil {
		return err
	}

	// exec template with post data
	// handle error if exist
	// store returned string from rendered template to buffer
	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
