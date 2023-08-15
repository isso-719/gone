package main

import (
	"embed"
	"github.com/isso-719/gone/src"
)

//go:embed audio/*.mp3
var gone embed.FS

func main() {
	src.Exec(gone)
}
