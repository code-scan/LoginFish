package static

import (
	"embed"
	_ "embed"
)

//go:embed dist
var StatFiles embed.FS
