package assets

import "embed"

var (
	//go:embed templates/*.gohtml
	PostTemplate embed.FS

	//go:embed examples/*.md
	Examples embed.FS
)
