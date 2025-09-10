package templates

import (
	"embed"
)

//go:embed prompts/*.md
var PromptFiles embed.FS
