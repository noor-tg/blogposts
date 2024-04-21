package renderer

import (
	"alnoor/blogposts/reader"
	"html/template"
	"io"
)

const postTemplate = `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags:<ul>
{{ range .Tags }}<li>{{.}}</li>{{end}}
</ul>`

func Render(w io.Writer, post reader.Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
