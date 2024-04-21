package blogposts

import "embed"

var (
	//go:embed templates/*.gohtml
	PostTemplate embed.FS
)
