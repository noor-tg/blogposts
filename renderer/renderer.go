package renderer

import (
	"alnoor/blogposts"
	"alnoor/blogposts/reader"
	"html/template"
	"io"
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	// parse template from embedded fs
	templ, err := template.ParseFS(blogposts.PostTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post reader.Post) error {

	// exec template with post data
	// handle error if exist
	// store returned string from rendered template to buffer
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}
