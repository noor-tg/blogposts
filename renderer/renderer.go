package renderer

import (
	"alnoor/blogposts"
	"alnoor/blogposts/reader"
	"bytes"
	"html/template"
	"io"

	"github.com/yuin/goldmark"
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	// make new template parser
	templ := template.New("out")
	// add function to output raw html for body parsed from markdown
	// scaping of html is there
	templ.Funcs(template.FuncMap{"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	}})

	// parse template from embedded fs
	out, err := templ.ParseFS(blogposts.PostTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	// return parser
	return &PostRenderer{templ: out}, nil
}

func (r *PostRenderer) Render(w io.Writer, post reader.Post) error {

	// convert body markdown to html
	post.Body = bodyParser(post)

	// exec template with post data
	// handle error if exist
	// store returned string from rendered template to buffer
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func bodyParser(post reader.Post) string {
	var buf bytes.Buffer

	if err := goldmark.Convert([]byte(post.Body), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}
