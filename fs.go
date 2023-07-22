package main

import (
	"embed"
)

//go:embed dist/index.html
var IndexHtml []byte

//go:embed dist/assets
var StaticAssets embed.FS
