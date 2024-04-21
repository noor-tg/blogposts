package renderer

import (
	"alnoor/blogposts/reader"
	"fmt"
	"io"
)

func Render(w io.Writer, post reader.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n", post.Title, post.Description)
	if err != nil {
		return err
	}

	fmt.Fprint(w, `Tags:<ul>`)
	for _, tag := range post.Tags {
		_, err := fmt.Fprintf(w, `<li>%s</li>`, tag)
		if err != nil {
			return err
		}
	}

	fmt.Fprint(w, `</ul>`)
	if err != nil {
		return err
	}

	return nil
}
